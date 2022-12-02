package os2mo

import (
	"context"
	"errors"

	"github.com/Khan/genqlient/graphql"
	guuid "github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcesItsystem() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcesReadItsystem,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

var TooManyResults = errors.New("query returned too many results")
var TooFewResults = errors.New("query returned no results")

func one[K any](objects []K) (*K, error) {
	if len(objects) > 1 {
		return nil, TooManyResults
	}
	if len(objects) < 1 {
		return nil, TooFewResults
	}
	return &objects[0], nil
}

func itsystemByUUID(uuid string, client graphql.Client) (*ITSystemFields, error) {
	parsed_uuid, err := guuid.Parse(uuid)
	if err != nil {
		return nil, err
	}

	response, err := getITSystemByUUID(
		client,
		parsed_uuid,
	)
	if err != nil {
		return nil, err
	}
	itsystem, err := one(response.GetItsystems())
	if err != nil {
		return nil, err
	}
	return &itsystem.ITSystemFields, nil
}

func itsystemByUserKey(user_key string, client graphql.Client) (*ITSystemFields, error) {
	response, err := getITSystemByUserKey(
		client,
		user_key,
	)
	if err != nil {
		return nil, err
	}
	itsystems := response.GetItsystems()
	if len(itsystems) > 1 {
		return nil, TooManyResults
	}
	if len(itsystems) < 1 {
		return nil, TooFewResults
	}
	return &itsystems[0].ITSystemFields, nil
}

func dataSourcesReadItsystem(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	uuid := d.Get("uuid").(string)
	user_key := d.Get("user_key").(string)

	if uuid != "" && user_key != "" {
		return diag.Errorf("uuid and user_key cannot both be set")
	}
	if uuid == "" && user_key == "" {
		return diag.Errorf("uuid or user_key must be set")
	}

	var itsystem *ITSystemFields
	var err error
	if uuid != "" {
		itsystem, err = itsystemByUUID(uuid, meta.(graphql.Client))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if user_key != "" {
		itsystem, err = itsystemByUserKey(user_key, meta.(graphql.Client))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.Set("uuid", itsystem.Uuid.String())
	d.Set("system_type", itsystem.System_type)
	d.Set("type", itsystem.Type)
	d.Set("user_key", itsystem.User_key)
	d.Set("name", itsystem.Name)
	d.SetId(itsystem.Uuid.String())

	return diags
}
