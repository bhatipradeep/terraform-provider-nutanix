package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFoundationHypervisorISOSDataSource(t *testing.T) {
	name := "hypervisor_isos"
	resourcePath := "data.nutanix_foundation_hypervisor_isos." + name
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccFoundationPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testHypervisorISOSDSConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourcePath, "kvm.0.filename"),
					resource.TestCheckResourceAttrSet(resourcePath, "esx.0.filename"),
				),
			},
		},
	})
}

func testHypervisorISOSDSConfig(name string) string {
	return fmt.Sprintf(`data "nutanix_foundation_hypervisor_isos" "%s" {}`, name)
}
