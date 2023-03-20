package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTfeOrganizationMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_organization_member",
		Description: "List users who are members of the organization.",
		List: &plugin.ListConfig{
			Hydrate: listOrganizationMember,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Hydrate: GetOrganizationName, Transform: transform.FromValue(), Description: "Name of the organization containing the organization member."},
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "Username of the member."},
			// Other columns
			{Name: "email", Type: proto.ColumnType_STRING, Description: "User email."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the organization membership."},
			{Name: "member", Type: proto.ColumnType_JSON, Transform: transform.FromField("User"), Description: "Full user information for the member."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the membership, e.g. active."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "User ID of the member."},
		},
	}
}

func listOrganizationMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_organization_member.listOrganizationMember", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)
	limit := d.QueryContext.Limit
	options := tfe.OrganizationMembershipListOptions{
		Include: "user,teams",
		ListOptions: tfe.ListOptions{
			// https://www.terraform.io/docs/cloud/api/index.html#pagination
			PageSize: 100,
		},
	}
	if limit != nil {
		if *limit < int64(100) {
			options.PageSize = int(*limit)
		}
	}

	pagesLeft := true
	for pagesLeft {
		result, err := conn.OrganizationMemberships.List(ctx, organizationName, options)
		if err != nil {
			plugin.Logger(ctx).Error("tfe_organization_member.listOrganizationMember", "query_error", err)
			return nil, err
		}
		for _, i := range result.Items {
			d.StreamListItem(ctx, i)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// Pagination
		if result.Pagination.CurrentPage < result.Pagination.TotalPages {
			options.PageNumber = result.Pagination.NextPage
		} else {
			pagesLeft = false
		}
	}
	return nil, nil
}
