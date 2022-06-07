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
			"virtual_machine_resources": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						// "creation_timestamp": &schema.Schema{
						// 	Type:     schema.Type,
						// 	Computed: true,
						// },
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

func dataSourceVMsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*orka.Client)

	var diags diag.Diagnostics

	id := "1"
	d.SetId(id)

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
			// time_string := validation.IsRFC3339Time()
			// vmri["creation_timestamp"] = vmr.Status[0].CreationTimestamp
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

	if err := d.Set("virtual_machine_resources", vmrs); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
