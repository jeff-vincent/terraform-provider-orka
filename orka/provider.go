package orka

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-vincent/orka-client-go"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ORKA_HOST", nil),
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ORKA_EMAIL", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ORKA_PASSWORD", nil),
			},
			"license_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ORKA_LICENSE_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vm_configs": resourceVMConfigs(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"orka_vms": dataSourceVMs(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	license_key := d.Get("license_key").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (email != "") && (password != "") {
		c, err := orka.NewClient(host, &email, &password, &license_key)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Orka client",
				Detail:   "Unable to authenticate user for authenticated Orka client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := orka.NewClient(host, nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Orka client",
			Detail:   "Unable to create anonymous Orka client",
		})
		return nil, diags
	}

	return c, diags
}
