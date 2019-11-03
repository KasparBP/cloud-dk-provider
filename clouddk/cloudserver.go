package clouddk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: createCloudServerResource,
		Read: readCloudServerResource,
		Delete: deleteCloudServerResource,
		Update: updateCloudServerResource,
		Importer: &schema.ResourceImporter {
			State: schema.ImportStatePassthrough,
		},
		// todo
	}
}

func createCloudServerResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}

func readCloudServerResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}


func updateCloudServerResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}

func deleteCloudServerResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}