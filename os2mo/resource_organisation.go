package os2mo

import (
	"text/template"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrganisation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganisationCreate,
		ReadContext:   resourceOrganisationRead,
		UpdateContext: resourceOrganisationUpdate,
		DeleteContext: resourceOrganisationDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"user_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

type UUIDStruct struct {
	Uuid string
}

type OrganisationArgs struct {
	Name    string
	UserKey string
}

func resourceOrganisationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// url := m.(*string)
	url := "http://localhost:8080/organisation/organisation"
	fmt.Println("URL:>", url)

	// var diags diag.Diagnostics

	name := d.Get("name").(string)
	user_key := d.Get("user_key").(string)

	org_args := OrganisationArgs{name, user_key}

	tmpl, err := template.New("org_json").Parse(`
    {
      "attributter": {
        "organisationegenskaber": [
          {
            "brugervendtnoegle": "{{.UserKey}}",
            "organisationsnavn": "{{.Name}}",
            "virkning": {
              "from": "-infinity",
              "to": "infinity"
            }
          }
        ]
      },
      "tilstande": {
        "organisationgyldighed": [
          {
            "gyldighed": "Aktiv",
            "virkning": {
              "from": "-infinity",
              "to": "infinity"
            }
          }
        ]
      }
    }
    `)

	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, org_args)
	if err != nil {
		panic(err)
	}

	result := tpl.String()
	fmt.Println("json str:", result)

	var jsonStr = []byte(result)
	fmt.Println("json bytes:", jsonStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var uuid_struct UUIDStruct
	unmarshal_err := json.Unmarshal([]byte(body), &uuid_struct)
	if unmarshal_err != nil {
		fmt.Println("error:", unmarshal_err)
		return diag.FromErr(unmarshal_err)
	}
	uuid := uuid_struct.Uuid
	d.SetId(uuid)

	return nil
}

func resourceOrganisationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceOrganisationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceOrganisationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
