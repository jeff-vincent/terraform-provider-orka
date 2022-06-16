package orka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-vincent/orka-client-go"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMCreate,
		DeleteContext: resourceVMCreate,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vms": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"orka_vm_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"orka_base_image": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"orka_cpu_core": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"vcpu_count": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"deployment_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"screen_sharing_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reserved_port_1_host": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_1_guest": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_1_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reserved_port_2_host": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_2_guest": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_2_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reserved_port_3_host": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_3_guest": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_port_3_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_machine_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"io_boost": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vcpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ssh_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ram": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"base_image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration_template": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vm_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"use_saved_state": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"creation_timestamp": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_required": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"replicas": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"gpu_passthrough": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*orka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vms := d.Get("vms").([]interface{})
	rb, _ := json.Marshal(vms[0])
	vm_data := orka.VMConfig{}
	err := json.Unmarshal(rb, &vm_data)

	if err != nil {
		tflog.Debug(ctx, string(err.Error()), nil)
	}

	vm_string := strings.NewReader(fmt.Sprintf(`{
		"orka_vm_name": "%s",
		"orka_base_image": "%s",
		"orka_cpu_core": %d,
		"vcpu_count": %d
	}`, vm_data.OrkaVMName, vm_data.OrkaBaseImage, vm_data.OrkaCPUCore, vm_data.VcpuCount))

	_, errs := c.CreateVM(vm_string)
	if errs != nil {
		return diag.FromErr(errs)
	}

	d.SetId("1")

	resourceVMRead(ctx, d, m)

	return diags
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

		vmri["deployment_status"] = vmr.VMDeploymentStatus

		if vmri["deployment_status"] == "Deployed" {
			vmri["reserved_port_1_host"] = vmr.Status[0].ReservedPorts[0].HostPort
			vmri["reserved_port_1_guest"] = vmr.Status[0].ReservedPorts[0].GuestPort
			vmri["reserved_port_1_protocol"] = vmr.Status[0].ReservedPorts[0].Protocol
			vmri["reserved_port_2_host"] = vmr.Status[0].ReservedPorts[1].HostPort
			vmri["reserved_port_2_guest"] = vmr.Status[0].ReservedPorts[1].GuestPort
			vmri["reserved_port_2_protocol"] = vmr.Status[0].ReservedPorts[1].Protocol
			vmri["reserved_port_3_host"] = vmr.Status[0].ReservedPorts[2].HostPort
			vmri["reserved_port_3_guest"] = vmr.Status[0].ReservedPorts[2].GuestPort
			vmri["reserved_port_3_protocol"] = vmr.Status[0].ReservedPorts[2].Protocol
			vmri["owner"] = vmr.Status[0].Owner
			vmri["image"] = vmr.Status[0].Image
			vmri["name"] = vmr.Status[0].VirtualMachineName
			vmri["screen_sharing_port"] = vmr.Status[0].ScreenSharingPort
			vmri["ssh_port"] = vmr.Status[0].SSHPort
			vmri["io_boost"] = vmr.Status[0].IoBoost
			vmri["virtual_machine_id"] = vmr.Status[0].VirtualMachineID
			vmri["cpu"] = vmr.Status[0].CPU
			vmri["vcpu"] = vmr.Status[0].Vcpu
			vmri["ram"] = vmr.Status[0].RAM
			vmri["base_image"] = vmr.Status[0].BaseImage
			vmri["configuration_template"] = vmr.Status[0].ConfigurationTemplate
			vmri["vm_status"] = vmr.Status[0].VMStatus
			vmri["use_saved_state"] = vmr.Status[0].UseSavedState
			vmri["creation_timestamp"] = vmr.Status[0].CreationTimestamp.String()
			vmri["tag"] = vmr.Status[0].Tag
			vmri["tag_required"] = vmr.Status[0].TagRequired
			vmri["replicas"] = vmr.Status[0].Replicas
		}

		if vmri["deployment_status"] == "Not Deployed" {
			vmri["owner"] = vmr.Owner
			vmri["image"] = vmr.Image
			vmri["name"] = vmr.VirtualMachineName
			vmri["cpu"] = vmr.CPU
			vmri["vcpu"] = vmr.Vcpu
			vmri["base_image"] = vmr.BaseImage
			vmri["io_boost"] = vmr.IoBoost
			vmri["use_saved_state"] = vmr.UseSavedState
			vmri["gpu_passthrough"] = vmr.GpuPassthrough
			vmri["configuration_template"] = vmr.ConfigurationTemplate
			vmri["tag"] = vmr.Tag
			vmri["tag_required"] = vmr.TagRequired
		}
		vmrs[i] = vmri
	}

	if err := d.Set("vms", vmrs); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// func flattenVMItems(orka.VMs) map[string]interface{} {
// 	return nil
// }

// func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*hc.Client)

// 	orderID := d.Id()

// 	if d.HasChange("items") {
// 		items := d.Get("items").([]interface{})
// 		ois := []hc.OrderItem{}

// 		for _, item := range items {
// 			i := item.(map[string]interface{})

// 			co := i["coffee"].([]interface{})[0]
// 			coffee := co.(map[string]interface{})

// 			oi := hc.OrderItem{
// 				Coffee: hc.Coffee{
// 					ID: coffee["id"].(int),
// 				},
// 				Quantity: i["quantity"].(int),
// 			}
// 			ois = append(ois, oi)
// 		}

// 		_, err := c.UpdateOrder(orderID, ois)
// 		if err != nil {
// 			return diag.FromErr(err)
// 		}

// 		d.Set("last_updated", time.Now().Format(time.RFC850))
// 	}

// 	return resourceOrderRead(ctx, d, m)
// }

// func resourceVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*hc.Client)

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	orderID := d.Id()

// 	err := c.DeleteOrder(orderID)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	// d.SetId("") is automatically called assuming delete returns no errors, but
// 	// it is added here for explicitness.
// 	d.SetId("")

// 	return diags
// }

// func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
// 	if orderItems != nil {
// 		ois := make([]interface{}, len(*orderItems), len(*orderItems))

// 		for i, orderItem := range *orderItems {
// 			oi := make(map[string]interface{})

// 			oi["coffee"] = flattenCoffee(orderItem.Coffee)
// 			oi["quantity"] = orderItem.Quantity
// 			ois[i] = oi
// 		}

// 		return ois
// 	}

// 	return make([]interface{}, 0)
// }

// func flattenCoffee(coffee hc.Coffee) []interface{} {
// 	c := make(map[string]interface{})
// 	c["id"] = coffee.ID
// 	c["name"] = coffee.Name
// 	c["teaser"] = coffee.Teaser
// 	c["description"] = coffee.Description
// 	c["price"] = coffee.Price
// 	c["image"] = coffee.Image

// 	return []interface{}{c}
// }
