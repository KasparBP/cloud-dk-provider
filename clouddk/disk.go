package clouddk

import (
	"context"
	"github.com/KasparBP/cloud-dk-provider/clouddk/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	dSize    = "size"
	dLabel   = "label"
	dPrimary = "primary"
	dServer =  "server_id"
)

func ResourceDisk() *schema.Resource {
	return &schema.Resource{
		Create: diskCreateResource,
		Read:   diskReadResource,
		Delete: diskDeleteResource,
		Update: diskUpdateResource,
		Importer: &schema.ResourceImporter {
			State: schema.ImportStatePassthrough,
		},
		Schema: diskSchema(),
	}
}


func diskSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		dSize:    &schema.Schema{Type: schema.TypeInt, Required: true},
		dLabel:   &schema.Schema{Type: schema.TypeString, Required: true},
		dPrimary: &schema.Schema{Type: schema.TypeBool, Optional: true, Computed: true},
		dServer:  &schema.Schema{Type: schema.TypeString, Required: true},
	}
}

func diskCreateResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}

func diskReadResource(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.Client)
	ctx := context.Background()
	serverId := d.Get(dServer).(string)
	disk, err := client.ClouddkService.GetDisk(ctx, serverId, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set(dSize, disk.Size)
	_ = d.Set(dLabel, disk.Label)
	_ = d.Set(dPrimary, disk.Primary)
	return err
}


func diskUpdateResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}

func diskDeleteResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}