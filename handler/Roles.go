package handler

import (
	"errors"
	"fmt"
	discord "github.com/chremoas/discord-gateway/proto"
	rolesrv "github.com/chremoas/role-srv/proto"
	"github.com/chremoas/services-common/config"
	redis "github.com/chremoas/services-common/redis"
	"github.com/chremoas/services-common/sets"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	//"regexp"
	"github.com/fatih/structs"
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

	// Type and Name are required so let's check for those
	if len(request.Type) == 0 {
		return errors.New("Type is required.")
	}

	if len(request.ShortName) == 0 {
		return errors.New("ShortName is required.")
	}

	if len(request.Name) == 0 {
		return errors.New("Name is required.")
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

	_, err = h.Redis.Client.HMSet(roleName, structs.Map(request)).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}

	return nil
}

func (h *rolesHandler) UpdateRole(ctx context.Context, request *rolesrv.UpdateInfo, response *rolesrv.NilMessage) error {
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
		response.Roles = append(response.Roles, &rolesrv.Role{Name: roles[role]})
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

func (h *rolesHandler) GetRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.Role) error {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.ShortName))

	role, err := h.Redis.Client.HGetAll(roleName).Result()

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
	response.Name = role["Name"]
	response.Color = int32(color)
	response.Hoist = hoist
	response.Position = int32(position)
	response.Permissions = int32(permissions)
	response.Managed = managed
	response.Mentionable = mentionable

	return nil
}

func (h *rolesHandler) SyncMembers(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.SyncRolesResponse) error {
	roles, err := h.getRoles()

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", roles)

	members, err := clients.discord.GetAllMembers(ctx, &discord.GetAllMembersRequest{NumberPerPage: 100})

	if err != nil {
		return err
	}

	for role := range roles {
		roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", roles[role]))
		cRole, err := h.Redis.Client.HGetAll(roleName).Result()

		if err != nil {
			return err
		}

		fmt.Printf("cRole.Name: %s, cRole.ShortName: %s\n", cRole["Name"], cRole["ShortName"])
	}

	if err != nil {
		return err
	}

	for member := range members.Members {
		fmt.Printf("Member: %+v\n", members.Members[member])
		for role := range members.Members[member].Roles {
			fmt.Printf("\tRole.Name: %+v\n", members.Members[member].Roles[role].Name)
		}
	}

	return nil
}

func (h *rolesHandler) SyncRoles(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.SyncRolesResponse) error {
	chremoasRoleSet := sets.NewStringSet()
	discordRoleSet := sets.NewStringSet()

	//var matchSpace = regexp.MustCompile(`\s`)
	//var matchDBError = regexp.MustCompile(`^Error 1062:`)
	//var matchDiscordError = regexp.MustCompile(`^The role '.*' already exists$`)

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
		//cRole = append(cRole, c)
	}

	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})
	if err != nil {
		return err
	}

	for role := range discordRoles.Roles {
		discordRoleSet.Add(discordRoles.Roles[role].Name)
	}

	fmt.Printf("chremoasRoleSet: %+v\n", chremoasRoleSet)
	fmt.Printf("discordRoleSet: %+v\n", discordRoleSet)
	fmt.Printf("Add: %+v\n", )

	return nil
}

//func syncRole(roleName string) error {
//	return nil
//}
