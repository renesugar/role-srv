package handler

import (
	"errors"
	"fmt"
	uauthsvc "github.com/chremoas/auth-srv/proto"
	discord "github.com/chremoas/discord-gateway/proto"
	"github.com/chremoas/role-srv/proto"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"regexp"
)

type rolesHandler struct {
	Client client.Client
}

type clientList struct {
	chremoasQuery uauthsvc.EntityQueryClient
	chremoasAdmin uauthsvc.EntityAdminClient
	discord       discord.DiscordGatewayClient
}

var clients clientList

func NewRolesHandler(conf *config.Configuration, service micro.Service) chremoas_roles.RolesHandler {
	c := service.Client()

	clients = clientList{
		chremoasQuery: uauthsvc.NewEntityQueryClient(conf.LookupService("srv", "auth"), c),
		chremoasAdmin: uauthsvc.NewEntityAdminClient(conf.LookupService("srv", "auth"), c),
		discord:       discord.NewDiscordGatewayClient(conf.LookupService("gateway", "discord"), c),
	}
	return &rolesHandler{}
}

func (h *rolesHandler) AddRole(ctx context.Context, request *chremoas_roles.AddRoleRequest, response *chremoas_roles.AddRoleResponse) error {
	roleName := request.Role.Name
	chatServiceGroup := request.Role.RoleNick

	_, err := clients.chremoasAdmin.RoleUpdate(ctx, &uauthsvc.RoleAdminRequest{
		Role:      &uauthsvc.Role{RoleName: roleName, ChatServiceGroup: chatServiceGroup},
		Operation: uauthsvc.EntityOperation_ADD_OR_UPDATE,
	})

	if err != nil {
		return err
	}

	_, err = clients.discord.CreateRole(ctx, &discord.CreateRoleRequest{Name: chatServiceGroup})

	if err != nil {
		return err
	}

	return nil
}

func (h *rolesHandler) RemoveRole(ctx context.Context, request *chremoas_roles.RemoveRoleRequest, response *chremoas_roles.RemoveRoleResponse) error {
	var dRoleName string
	roleName := request.Name

	chremoasRoles, err := clients.chremoasQuery.GetRoles(ctx, &uauthsvc.EntityQueryRequest{})

	for cr := range chremoasRoles.List {
		if chremoasRoles.List[cr].RoleName == roleName {
			dRoleName = chremoasRoles.List[cr].ChatServiceGroup
		}
	}

	_, err = clients.chremoasAdmin.RoleUpdate(ctx, &uauthsvc.RoleAdminRequest{
		Role:      &uauthsvc.Role{RoleName: roleName, ChatServiceGroup: "Doesn't matter"},
		Operation: uauthsvc.EntityOperation_REMOVE,
	})

	if err != nil {
		return err
	}

	_, err = clients.discord.DeleteRole(ctx, &discord.DeleteRoleRequest{Name: dRoleName})

	if err != nil {
		return err
	}

	return nil
}

func (h *rolesHandler) GetRoles(ctx context.Context, request *chremoas_roles.GetRolesRequest, response *chremoas_roles.GetRolesResponse) error {
	output, err := clients.chremoasQuery.GetRoles(ctx, &uauthsvc.EntityQueryRequest{})

	if err != nil {
		return err
	}

	if output.String() == "" {
		return errors.New("There are no roles defined")
	}

	for role := range output.List {
		response.Roles = append(response.Roles, &chremoas_roles.DiscordRole{Name: output.List[role].RoleName, RoleNick: output.List[role].ChatServiceGroup})
	}

	return nil
}

func (h *rolesHandler) SyncRoles(ctx context.Context, request *chremoas_roles.SyncRolesRequest, response *chremoas_roles.SyncRolesResponse) error {
	var matchSpace = regexp.MustCompile(`\s`)
	var matchDBError = regexp.MustCompile(`^Error 1062:`)
	var matchDiscordError = regexp.MustCompile(`^The role '.*' already exists$`)

	//listDRoles(ctx, req)
	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})

	if err != nil {
		return err
	}

	//listRoles(ctx, req)
	chremoasRoles, err := clients.chremoasQuery.GetRoles(ctx, &uauthsvc.EntityQueryRequest{})

	if err != nil {
		return err
	}

	for dr := range discordRoles.Roles {
		_, err := clients.chremoasAdmin.RoleUpdate(ctx, &uauthsvc.RoleAdminRequest{
			Role:      &uauthsvc.Role{ChatServiceGroup: discordRoles.Roles[dr].Name, RoleName: matchSpace.ReplaceAllString(discordRoles.Roles[dr].Name, "_")},
			Operation: uauthsvc.EntityOperation_ADD_OR_UPDATE,
		})

		if err != nil {
			if !matchDBError.MatchString(err.Error()) {
				return err
			}
			fmt.Printf("dr err: %+v\n", err)
		} else {
			response.Roles = append(response.Roles, &chremoas_roles.SyncRoles{Source: "Discord", Destination: "Chremoas", Name: discordRoles.Roles[dr].Name})
		}
	}

	for cr := range chremoasRoles.List {
		_, err := clients.discord.CreateRole(ctx, &discord.CreateRoleRequest{Name: chremoasRoles.List[cr].ChatServiceGroup})

		if err != nil {
			if !matchDiscordError.MatchString(err.Error()) {
				return err
			}
			fmt.Printf("cr err: %+v\n", err)
		} else {
			response.Roles = append(response.Roles, &chremoas_roles.SyncRoles{Source: "Chremoas", Destination: "Discord", Name: chremoasRoles.List[cr].ChatServiceGroup})
		}
	}

	return nil
}
