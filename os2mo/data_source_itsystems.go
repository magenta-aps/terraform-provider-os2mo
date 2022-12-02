package os2mo

import (
	"context"
	"strings"

	"github.com/Khan/genqlient/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcesItsystems() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcesReadItsystems,
		Schema: map[string]*schema.Schema{
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func transform[K any, V any](data []K, f func(K) V) []V {

	mapped := make([]V, len(data))

	for i, e := range data {
		mapped[i] = f(e)
	}

	return mapped
}

/*
func reduce[K any, V any](data []K, f func(K, V) V, init V) V {
	result := init

	for _, e := range data {
		result = f(e, result)
	}

	return result
}
*/

func gql2tf(value getITSystemsItsystemsITSystem) map[string]interface{} {
	itsystem := value.ITSystemFields

	e := make(map[string]interface{})
	e["uuid"] = itsystem.Uuid.String()
	e["system_type"] = itsystem.System_type
	e["type"] = itsystem.Type
	e["user_key"] = itsystem.User_key
	e["name"] = itsystem.Name

	return e
}

func gql2id(value getITSystemsItsystemsITSystem) string {
	itsystem := value.ITSystemFields
	return itsystem.Uuid.String()
}

func dataSourcesReadItsystems(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	response, err := getITSystems(meta.(graphql.Client))
	if err != nil {
		return diag.FromErr(err)
	}

	results := transform(response.GetItsystems(), gql2tf)
	d.Set("results", results)
	d.SetId(strings.Join(transform(response.GetItsystems(), gql2id), "|"))

	var diags diag.Diagnostics
	return diags
}
