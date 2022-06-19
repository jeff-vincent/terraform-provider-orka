package orka

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-vincent/orka-client-go"
)

func dataSourceNodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMsRead,
		Schema: map[string]*schema.Schema{
			"nodes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_cpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allocatable_cpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_gpu": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allocatable_gpu": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_memory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_cpu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_memory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"orka_tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*orka.Client)

	var diags diag.Diagnostics

	id := "1"
	d.SetId(id)
	// TODO: write GetNodes() method
	nodes, err := c.GetNodes()
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: flatten GetNodes() response

	if err := d.Set("nodes", nil); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
