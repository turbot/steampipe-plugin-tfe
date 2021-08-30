package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableTfeOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_organization",
		Description: "Organizations for the user.",
		List: &plugin.ListConfig{
			Hydrate: listOrganization,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getOrganization,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			// Other columns
			{Name: "collaborator_auth_policy", Type: proto.ColumnType_STRING, Description: "Authentication policy: password, two_factor_mandatory."},
			{Name: "cost_estimation_enabled", Type: proto.ColumnType_BOOL, Description: "Whether or not the cost estimation feature is enabled for all workspaces in the organization."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the token was created."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address for notifications."},
			{Name: "external_id", Type: proto.ColumnType_STRING, Description: "External ID for the organization."},
			{Name: "owners_team_saml_role_id", Type: proto.ColumnType_STRING, Description: "SAML role mapped to the owners team."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "Permissions for the organization."},
			{Name: "saml_enabled", Type: proto.ColumnType_BOOL, Description: "True if SAML is enabled for the organization."},
			{Name: "session_remember", Type: proto.ColumnType_INT, Transform: transform.FromField("SessionRemember"), Description: "Session expiration (minutes)."},
			{Name: "session_timeout", Type: proto.ColumnType_INT, Transform: transform.FromField("SessionTimeout"), Description: "Session timeout after inactivity (minutes)."},
			{Name: "trial_expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the trial, if any, expires."},
			{Name: "two_factor_conformant", Type: proto.ColumnType_BOOL, Description: "If true, members are required to use two factor authentication."},
		},
	}
}

func listOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.listOrganization", "connection_error", err)
		return nil, err
	}
	limit := d.QueryContext.Limit
	var defaultPageSize, pagesToIterate, lastPageSize int64
	defaultPageSize = 20
	options := tfe.OrganizationListOptions{}
	if limit != nil {
		// default size is 20
		if *limit < defaultPageSize {
			options.PageSize = int(*limit)
		}
		pagesToIterate = *limit / defaultPageSize
		lastPageSize = *limit % defaultPageSize
	}
	pagesLeft := true
	for pagesLeft {
		result, err := conn.Organizations.List(ctx, options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_organization.listOrganization", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
		}
		// Pagination with limit
		if limit != nil && *limit > defaultPageSize {
			if result.Pagination.CurrentPage < int(pagesToIterate) {
				options.PageNumber = result.Pagination.NextPage
			} else if result.Pagination.CurrentPage == int(pagesToIterate) {
				options.PageSize = int(lastPageSize)
			} else {
				pagesLeft = false
			}
		} else {
			// normal pagination
			if result.Pagination.CurrentPage < result.Pagination.TotalPages {
				options.PageNumber = result.Pagination.NextPage
			} else {
				pagesLeft = false
			}
		}
	}
	return nil, nil
}

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getOrganization", "connection_error", err)
		return nil, err
	}
	result, err := conn.Organizations.Read(ctx, d.KeyColumnQuals["name"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization.getOrganization", "query_error", err)
		return nil, err
	}
	return result, nil
}
