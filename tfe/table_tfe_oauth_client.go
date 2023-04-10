package tfe

import (
	"context"

	"github.com/hashicorp/go-tfe"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTfeOauthClient(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tfe_oauth_client",
		Description: "OAuth clients in the organization.",
		List: &plugin.ListConfig{
			Hydrate: listOauthClient,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getOauthClient,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the oauth client."},
			// Others columns
			{Name: "key", Type: proto.ColumnType_STRING, Description: "The key of the oauth client."},
			{Name: "api_url", Type: proto.ColumnType_STRING, Description: "The API url of the service provider."},
			{Name: "http_url", Type: proto.ColumnType_STRING, Description: "The HTTP url of the service provider."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the oauth client was created."},
			{Name: "callback_url", Type: proto.ColumnType_STRING, Description: "The callback url of the oauth client."},
			{Name: "connect_path", Type: proto.ColumnType_STRING, Description: "The connection path of the oauth client."},
			{Name: "organization_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Organization.Name"), Description: "Name of the organization containing the oauth client."},
			{Name: "service_provider", Type: proto.ColumnType_STRING, Description: "The VCS provider being connected with. Valid options are ado_server, ado_services, github, github_enterprise, gitlab_hosted, gitlab_community_edition, or gitlab_enterprise_edition."},
			{Name: "service_provider_name", Type: proto.ColumnType_STRING, Description: "The name of VCS provider being connected with."},
			{Name: "rsa_public_key", Type: proto.ColumnType_STRING, Description: "The public key of the oauth client.", Transform: transform.FromField("RSAPublicKey")},
			{Name: "oauth_token", Type: proto.ColumnType_JSON, Description: "The token information you were given by your VCS provider."},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: "The organization information."},
		},
	}
}

func listOauthClient(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_oauth_client.listOauthClient", "connection_error", err)
		return nil, err
	}
	data, err := GetOrganizationName(ctx, d, h)
	if err != nil {
		return nil, err
	}
	organizationName := data.(string)

	limit := d.QueryContext.Limit
	options := tfe.OAuthClientListOptions{
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
		result, err := conn.OAuthClients.List(ctx, organizationName, options)
		if err != nil {
			if isNotFoundError(err) {
				return nil, nil
			}
			plugin.Logger(ctx).Error("tfe_oauth_client.listOauthClient", "query_error", err)
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

func getOauthClient(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tfe_oauth_client.getOauthClient", "connection_error", err)
		return nil, err
	}
	result, err := conn.OAuthClients.Read(ctx, d.EqualsQuals["id"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("tfe_oauth_client.getOauthClient", "query_error", err)
		return nil, err
	}
	return result, nil
}
