package handler

import (
	"errors"
	"fmt"
	discord "github.com/chremoas/discord-gateway/proto"
	rolesrv "github.com/chremoas/role-srv/proto"
	"github.com/chremoas/services-common/config"
	redis "github.com/chremoas/services-common/redis"
	"github.com/chremoas/services-common/sets"
	"github.com/fatih/structs"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"regexp"
	"strconv"
	"strings"
)

type rolesHandler struct {
	Client client.Client
	Redis  *redis.Client
}

type clientList struct {
	discord discord.DiscordGatewayClient
}

var clients clientList
var roleKeys = []string{"Name", "Color", "Hoist", "Position", "Permissions", "Managed", "Mentionable"}
var roleTypes = []string{"internal", "discord"}

func NewRolesHandler(config *config.Configuration, service micro.Service) rolesrv.RolesHandler {
	c := service.Client()

	clients = clientList{
		discord: discord.NewDiscordGatewayClient(config.LookupService("gateway", "discord"), c),
	}

	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
	redisClient := redis.Init(addr, config.Redis.Password, config.Redis.Database, config.LookupService("srv", "perms"))

	_, err := redisClient.Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &rolesHandler{Redis: redisClient}
}

func (h *rolesHandler) GetRoleKeys(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.StringList) error {
	response.Value = roleKeys
	return nil
}

func (h *rolesHandler) GetRoleTypes(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.StringList) error {
	response.Value = roleTypes
	return nil
}

func (h *rolesHandler) AddRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.NilMessage) error {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.ShortName))
	filterA := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.FilterA))
	filterB := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.FilterB))

	// Type, Name and the filters are required so let's check for those
	if len(request.Type) == 0 {
		return errors.New("Type is required.")
	}

	if len(request.ShortName) == 0 {
		return errors.New("ShortName is required.")
	}

	if len(request.Name) == 0 {
		return errors.New("Name is required.")
	}

	if len(request.FilterA) == 0 {
		return errors.New("FilterA is required.")
	}

	if len(request.FilterB) == 0 {
		return errors.New("FilterB is required.")
	}

	if !validListItem(request.Type, roleTypes) {
		return fmt.Errorf("`%s` isn't a valid Role Type.", request.Type)
	}

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("Role `%s` already exists.", request.Name)
	}

	// Check if filter A exists
	exists, err = h.Redis.Client.Exists(filterA).Result()

	if err != nil {
		return err
	}

	if exists == 0 && request.FilterA != "wildcard" {
		return fmt.Errorf("FilterA `%s` doesn't exists.", request.FilterA)
	}

	// Check if filter B exists
	exists, err = h.Redis.Client.Exists(filterB).Result()

	if err != nil {
		return err
	}

	if exists == 0 && request.FilterB != "wildcard" {
		return fmt.Errorf("FilterB `%s` doesn't exists.", request.FilterB)
	}

	_, err = h.Redis.Client.HMSet(roleName, structs.Map(request)).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}

	return nil
}

func (h *rolesHandler) UpdateRole(ctx context.Context, request *rolesrv.UpdateInfo, response *rolesrv.NilMessage) error {
	// Does this actually work? -brian
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.Name))

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Role `%s` doesn't exists.", request.Name)
	}

	if !validListItem(request.Key, roleKeys) {
		return fmt.Errorf("`%s` isn't a valid Role Key.", request.Key)
	}

	h.Redis.Client.HSet(roleName, request.Key, request.Value)

	return nil
}

func validListItem(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (h *rolesHandler) RemoveRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.NilMessage) error {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.ShortName))

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Role `%s` doesn't exists.", request.ShortName)
	}

	_, err = h.Redis.Client.Del(roleName).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) GetRoles(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.GetRolesResponse) error {
	roles, err := h.getRoles()

	if err != nil {
		return err
	}

	for role := range roles {
		roleName, err := h.Redis.Client.HGet(h.Redis.KeyName(fmt.Sprintf("role:%s", roles[role])), "Name").Result()
		if err != nil {
			return err
		}

		response.Roles = append(response.Roles, &rolesrv.Role{
			ShortName: roles[role],
			Name:      roleName,
		})
	}

	return nil
}

func (h *rolesHandler) getRoles() ([]string, error) {
	var roleList []string
	roles, err := h.Redis.Client.Keys(h.Redis.KeyName("role:*")).Result()

	if err != nil {
		return nil, err
	}

	for role := range roles {
		roleName := strings.Split(roles[role], ":")
		roleList = append(roleList, roleName[len(roleName)-1])
	}

	return roleList, nil
}

func (h *rolesHandler) getRole(name string) (role map[string]string, err error) {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", name))
	r, err := h.Redis.Client.HGetAll(roleName).Result()

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h *rolesHandler) GetRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.Role) error {
	role, err := h.getRole(request.ShortName)

	if err != nil {
		return err
	}

	color, err := strconv.ParseInt(role["Color"], 10, 32)
	position, err := strconv.ParseInt(role["Position"], 10, 32)
	permissions, err := strconv.ParseInt(role["Permissions"], 10, 32)
	hoist, err := strconv.ParseBool(role["Hoist"])
	managed, err := strconv.ParseBool(role["Managed"])
	mentionable, err := strconv.ParseBool(role["Mentionable"])

	response.ShortName = request.ShortName
	response.Type = role["Type"]
	response.FilterA = role["FilterA"]
	response.FilterB = role["FilterB"]
	response.Name = role["Name"]
	response.Color = int32(color)
	response.Hoist = hoist
	response.Position = int32(position)
	response.Permissions = int32(permissions)
	response.Managed = managed
	response.Mentionable = mentionable

	return nil
}

func (h *rolesHandler) SyncMembers(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.NilMessage) error {
	return h.syncMembers(ctx)
}

func (h *rolesHandler) syncMembers(ctx context.Context) error {
	var roleNameMap = make(map[string]string)
	var membershipSets = make(map[string]*sets.StringSet)

	// Discord limit is 1000, should probably make this a config option. -brian
	var numberPerPage int32 = 1000
	var memberCount = 1
	var memberId = ""

	// Need to pre-populate the membership sets with all the users so we can pick up users with no roles.
	for memberCount > 0 {
		members, err := clients.discord.GetAllMembers(ctx, &discord.GetAllMembersRequest{NumberPerPage: numberPerPage, After: memberId})

		if err != nil {
			return err
		}

		for m := range members.Members {
			userId := members.Members[m].User.Id
			if _, ok := membershipSets[userId]; !ok {
				membershipSets[userId] = sets.NewStringSet()
			}

			if members.Members[m].User.Id > memberId {
				memberId = members.Members[m].User.Id
			}
		}

		memberCount = len(members.Members)
	}

	// Get all the Roles from discord and create a map of their name to theid Id
	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})
	if err != nil {
		return err
	}

	for d := range discordRoles.Roles {
		roleNameMap[discordRoles.Roles[d].Name] = discordRoles.Roles[d].Id
	}

	// Get all the Chremoas roles and build membership Sets
	chremoasRoles, err := h.getRoles()
	if err != nil {
		return err
	}

	for r := range chremoasRoles {
		membership, err := h.getRoleMembership(chremoasRoles[r])
		roleName, err := h.getRole(chremoasRoles[r])
		roleId := roleNameMap[roleName["Name"]]

		for m := range membership.Set {
			membershipSets[m].Add(roleId)
		}

		if err != nil {
			return err
		}
	}

	// Apply the membership sets to discord overwriting anything that's there.
	for m := range membershipSets {
		clients.discord.UpdateMember(ctx, &discord.UpdateMemberRequest{
			Operation: discord.MemberUpdateOperation_ADD_OR_UPDATE_ROLES,
			UserId:    m,
			RoleIds:   membershipSets[m].ToSlice(),
		})
	}
	return nil
}

func (h *rolesHandler) getRoleMembership(role string) (members *sets.StringSet, err error) {
	var filterASet = sets.NewStringSet()
	var filterBSet = sets.NewStringSet()

	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", role))

	r, err := h.Redis.Client.HGetAll(roleName).Result()

	if err != nil {
		return filterASet, err
	}

	filterAName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", r["FilterA"]))
	filterBName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", r["FilterB"]))

	exists, err := h.Redis.Client.Exists(filterAName).Result()

	if err != nil {
		return filterASet, err
	}

	if exists == 0 && r["FilterA"] != "wildcard" {
		return filterASet, fmt.Errorf("Filter `%s` doesn't exists.", r["FilterA"])
	}

	exists, err = h.Redis.Client.Exists(filterBName).Result()

	if err != nil {
		return filterASet, err
	}

	if exists == 0 && r["FilterB"] != "wildcard" {
		return filterASet, fmt.Errorf("Filter `%s` doesn't exists.", r["FilterB"])
	}

	filterA, err := h.Redis.Client.SMembers(filterAName).Result()

	if err != nil {
		return filterASet, err
	}

	filterB, err := h.Redis.Client.SMembers(filterBName).Result()

	if err != nil {
		return filterASet, err
	}

	filterASet.FromSlice(filterA)
	filterBSet.FromSlice(filterB)

	if r["FilterA"] == "wildcard" {
		return filterBSet, nil
	}

	if r["FilterB"] == "wildcard" {
		return filterASet, nil
	}

	return filterASet.Intersection(filterBSet), nil
}

func (h *rolesHandler) SyncRoles(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.RoleSyncResponse) error {
	var matchDiscordError = regexp.MustCompile(`^The role '.*' already exists$`)
	chremoasRoleSet := sets.NewStringSet()
	discordRoleSet := sets.NewStringSet()

	chremoasRoles, err := h.getRoles()
	if err != nil {
		return err
	}

	for role := range chremoasRoles {
		roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", chremoasRoles[role]))
		c, err := h.Redis.Client.HGetAll(roleName).Result()

		if err != nil {
			return err
		}

		chremoasRoleSet.Add(c["Name"])
	}

	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})
	if err != nil {
		return err
	}

	ignoreSet := sets.NewStringSet()
	ignoreSet.Add("Chremoas")
	ignoreSet.Add("@everyone")

	for role := range discordRoles.Roles {
		if !ignoreSet.Contains(discordRoles.Roles[role].Name) {
			discordRoleSet.Add(discordRoles.Roles[role].Name)
		}
	}

	toAdd := chremoasRoleSet.Difference(discordRoleSet)
	toDelete := discordRoleSet.Difference(chremoasRoleSet)

	for r := range toAdd.Set {
		_, err := clients.discord.CreateRole(ctx, &discord.CreateRoleRequest{Name: r})

		if err != nil {
			if matchDiscordError.MatchString(err.Error()) {
				// The role list was cached most likely so we'll pretend we didn't try
				// to create it just now. -brian
				continue
			} else {
				return err
			}
		}

		response.Added = append(response.Added, r)
	}

	for r := range toDelete.Set {
		response.Removed = append(response.Removed, r)
		_, err := clients.discord.DeleteRole(ctx, &discord.DeleteRoleRequest{Name: r})

		if err != nil {
			return err
		}
	}

	return nil
}

//
// Filter related stuff
//

func (h *rolesHandler) GetFilters(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.FilterList) error {
	filters, err := h.Redis.Client.Keys(h.Redis.KeyName("filter_description:*")).Result()

	if err != nil {
		return err
	}

	for filter := range filters {
		filterDescription, err := h.Redis.Client.Get(filters[filter]).Result()

		if err != nil {
			return err
		}

		filterName := strings.Split(filters[filter], ":")

		response.FilterList = append(response.FilterList,
			&rolesrv.Filter{Name: filterName[len(filterName)-1], Description: filterDescription})
	}

	return nil
}

func (h *rolesHandler) AddFilter(ctx context.Context, request *rolesrv.Filter, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Name))

	// Type and Name are required so let's check for those
	if len(request.Name) == 0 {
		return errors.New("Name is required.")
	}

	if len(request.Description) == 0 {
		return errors.New("Description is required.")
	}

	exists, err := h.Redis.Client.Exists(filterName).Result()

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("Filter `%s` already exists.", request.Name)
	}

	_, err = h.Redis.Client.Set(filterName, request.Description, 0).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}

	return nil
}

func (h *rolesHandler) RemoveFilter(ctx context.Context, request *rolesrv.Filter, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Name))
	filterMembers := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Name))

	exists, err := h.Redis.Client.Exists(filterName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Name)
	}

	members, err := h.Redis.Client.SMembers(filterMembers).Result()

	if len(members) > 0 {
		return fmt.Errorf("Filter `%s` not empty.", request.Name)
	}

	_, err = h.Redis.Client.Del(filterName).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) GetMembers(ctx context.Context, request *rolesrv.Filter, response *rolesrv.MemberList) error {
	var memberlist []string
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Name))

	filters, err := h.Redis.Client.SMembers(filterName).Result()

	if err != nil {
		return err
	}

	for filter := range filters {
		memberlist = append(memberlist, filters[filter])
	}

	response.Members = memberlist
	return nil
}

func (h *rolesHandler) AddMembers(ctx context.Context, request *rolesrv.Members, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Filter))
	filterDesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Filter))

	exists, err := h.Redis.Client.Exists(filterDesc).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Filter)
	}

	for member := range request.Name {
		_, err = h.Redis.Client.SAdd(filterName, request.Name[member]).Result()
	}

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) RemoveMembers(ctx context.Context, request *rolesrv.Members, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Filter))
	filterDesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Filter))

	exists, err := h.Redis.Client.Exists(filterDesc).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Filter)
	}

	//isMember, err := h.Redis.Client.SIsMember(filterName, request.Name).Result()
	//if !isMember {
	//	return fmt.Errorf("`%s` not a member of filter '%s'", request.Name, request.Filter)
	//}

	for member := range request.Name {
		_, err = h.Redis.Client.SRem(filterName, request.Name[member]).Result()
	}

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}
