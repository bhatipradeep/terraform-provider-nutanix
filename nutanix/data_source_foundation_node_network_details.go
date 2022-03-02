package nutanix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-nutanix/client/foundation"
)

func dataSourceNodeNetworkDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeNetworkDetailsRead,
		Schema: map[string]*schema.Schema{
			"ipv6_addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cvm_gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipmi_netmask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cvm_vlan_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hypervisor_hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hypervisor_netmask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cvm_netmask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipmi_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hypervisor_gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cvm_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hypervisor_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipmi_gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeNetworkDetailsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*Client).FoundationClientAPI
	v, ok := d.GetOk("ipv6_addresses")
	if !ok && len(v.([]interface{})) == 0 {
		return diag.Errorf("please provide the ipv6_addresses")
	}
	input := new(foundation.NodeNetworkDetailsInput)
	ipv6Addresses := expandStringList(v.([]interface{}))
	for _, val := range ipv6Addresses {
		input.Nodes = append(input.Nodes, foundation.NodeIpv6Input{Ipv6Address: *val})
	}

	resp, err := conn.Networking.NodeNetworkDetails(ctx, input)
	if err != nil {
		return diag.FromErr(err)
	}
	nodes := make([]map[string]string, len(resp.Nodes))
	for k, v := range resp.Nodes {
		node := make(map[string]string)
		node["cvm_gateway"] = v.CvmGateway
		node["ipmi_netmask"] = v.IpmiNetmask
		node["ipv6_address"] = v.Ipv6Address
		node["cvm_vlan_id"] = v.CvmVlanID
		node["hypervisor_hostname"] = v.HypervisorHostname
		node["hypervisor_netmask"] = v.HypervisorNetmask
		node["cvm_netmask"] = v.IpmiIP
		node["ipmi_ip"] = v.IpmiIP
		node["hypervisor_gateway"] = v.HypervisorGateway
		node["error"] = v.Error
		node["cvm_ip"] = v.CvmIP
		node["hypervisor_ip"] = v.HypervisorIP
		node["ipmi_gateway"] = v.IpmiGateway
		nodes[k] = node
	}
	setErr := d.Set("nodes", nodes)
	if setErr != nil {
		return diag.FromErr(err)
	}
	d.SetId(resource.UniqueId())
	return nil
}
