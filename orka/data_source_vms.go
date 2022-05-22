package orka

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-vincent/orka-client-go"
)

func dataSourceVMs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": {
							Type:     schema.TypeList,
							Computed: true,
						},
						"coffee_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_teaser": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_image": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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
	fmt.Println(vms)
	return diags
}
