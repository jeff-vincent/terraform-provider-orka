package orka

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-vincent/orka-client-go"
)

func dataSourceVMs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMsRead,
		Schema: map[string]*schema.Schema{
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"help": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"errors": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"virtual_machine_resources": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceVMsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*orka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vms, err := c.GetVMs()
	if err != nil {
		return diag.FromErr(err)
	}

	r := make(map[string]interface{})

	data, _ := json.Marshal(vms)
	rm := json.Unmarshal(data, &r)

	if err := d.Set("virtual_machine_resources", rm); err != nil {
		return diag.FromErr(err)
	}

	// if err := d.Set("help", vms.Help); err != nil {
	// 	return diag.FromErr(err)
	// }

	return diags
}
