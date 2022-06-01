package orka

import (
	"context"

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

	if err := d.Set("virtual_machine_resources", vms.VirtualMachineResources); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// Need to flatten...
// func flattenVMsItemsData(orderItems *[]orka.GetVMs) []interface{} {
// 	if orderItems != nil {
// 		ois := make([]interface{}, len(*orderItems), len(*orderItems))

// 		for i, orderItem := range *orderItems {
// 			oi := make(map[string]interface{})

// 			oi["coffee_id"] = orderItem.Coffee.ID
// 			oi["coffee_name"] = orderItem.Coffee.Name
// 			oi["coffee_teaser"] = orderItem.Coffee.Teaser
// 			oi["coffee_description"] = orderItem.Coffee.Description
// 			oi["coffee_price"] = orderItem.Coffee.Price
// 			oi["coffee_image"] = orderItem.Coffee.Image
// 			oi["quantity"] = orderItem.Quantity

// 			ois[i] = oi
// 		}

// 		return ois
// 	}

// 	return make([]interface{}, 0)
// }

// map of strings with the described limitations.

// Description: "Metadata associated with the client, in the form of an object with string values (" +
// 	"max 255 chars). Maximum of 10 metadata properties allowed. Field names (" +
// 	"max 255 chars) are alphanumeric and may only include the following special characters: :," +
// 	"-+=_*?\"/\\()<>@ [Tab] [Space]",
