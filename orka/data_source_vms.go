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
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_owner": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"base_image": &schema.Schema{
							Type:     schema.TypeString,
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

	vmrs := make([]interface{}, len(vms.VirtualMachineResources))

	for i, vmr := range vms.VirtualMachineResources {
		vmri := make(map[string]interface{})

		if vmr.VMDeploymentStatus == "Deployed" {
			vmri["vm_owner"] = vmr.Status[0].Owner
			vmri["base_image"] = vmr.Status[0].Image
		}
		if vmr.VMDeploymentStatus == "Not Deployed" {
			vmri["vm_owner"] = vmr.Owner
			vmri["base_image"] = vmr.Image
		}
		vmrs[i] = vmri
	}

	if err := d.Set("virtual_machine_resources", vmrs); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
