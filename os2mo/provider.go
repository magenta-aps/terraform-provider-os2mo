// SPDX-FileCopyrightText: Magenta ApS
// SPDX-License-Identifier: MPL-2.0
package os2mo

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MO_URL", nil),
				Description: "URL to GraphQL endpoint on OS2mo.",
			},
			// TODO: Auth server URL
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
				Description: "OIDC Client ID for authentification.",
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", nil),
				Description: "OIDC Client Secret for authentification.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"os2mo_organisation": resourceOrganisation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"os2mo_itsystem":  dataSourcesItsystem(),
			"os2mo_itsystems": dataSourcesItsystems(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	// client_id := d.Get("client_id").(string)
	// client_secret := d.Get("client_secret").(string)

	// TODO: OAUTH2 client
	client := graphql.NewClient(url, http.DefaultClient)
	return client, nil
}
