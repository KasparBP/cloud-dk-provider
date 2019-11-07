package clouddk

import (
	"context"
	"github.com/KasparBP/cloud-dk-provider/clouddk/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	csHostname = "hostname"
	csLabel = "label"
	csInitialRootPassword = "initial_root_password"
	csTemplate = "template"
	csPackage = "package"
	csLocation = "location"
)

func ResourceCloudServer() *schema.Resource {
	return &schema.Resource{
		Create: cloudServerCreateResource,
		Read:   cloudServerReadResource,
		Delete: cloudServerDeleteResource,
		Update: cloudServerUpdateResource,
		Importer: &schema.ResourceImporter {
			State: schema.ImportStatePassthrough,
		},
		Schema: cloudServerSchema(),
		// todo
	}
}

func cloudServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		csHostname: &schema.Schema{Type: schema.TypeString, Required: true},
		csLabel: &schema.Schema{Type: schema.TypeString, Required: true},
		//"cpus": &schema.Schema{Type: schema.TypeInt, Optional: true},
		//"memory": &schema.Schema{Type: schema.TypeInt, Optional: true},
		csInitialRootPassword: &schema.Schema{Type: schema.TypeString, Optional: true},
		csTemplate: &schema.Schema{Type: schema.TypeString, Optional: true},
		csPackage: &schema.Schema{Type: schema.TypeString, Optional: true},
		csLocation: &schema.Schema{Type: schema.TypeString, Optional: true},
	}
}

func cloudServerCreateResource(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.Client)
	rootpw := d.Get(csInitialRootPassword).(string)
	server, err := client.ClouddkService.CreateCloudServer(context.Background(), &api.CloudServer{
		HostName:            d.Get(csHostname).(string),
		Label:               d.Get(csLabel).(string),
		InitialRootPassword: &rootpw,
		Template:            api.Template{Identifier: d.Get(csTemplate).(string)},
		Location:            api.Location{Identifier: d.Get(csLocation).(string)},
		Package:             api.Package{Identifier: d.Get(csPackage).(string)},
	})
	if err != nil {
		return err
	} else {
		d.SetId(server.Identifier)
	}
	return nil
}

func cloudServerReadResource(d *schema.ResourceData, m interface{}) error {
	client := m.(*api.Client)
	ctx := context.Background()
	cloudServer, err := client.ClouddkService.GetCloudServer(ctx, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set(csHostname, cloudServer.HostName)
	_ = d.Set(csLabel, cloudServer.Label)
	_ = d.Set(csInitialRootPassword, cloudServer.InitialRootPassword)
	_ = d.Set(csTemplate, cloudServer.Template.Identifier)
	_ = d.Set(csLocation, cloudServer.Location.Identifier)
	_ = d.Set(csPackage, cloudServer.Package.Identifier)
	return nil
}


func cloudServerUpdateResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}

func cloudServerDeleteResource(d *schema.ResourceData, m interface{}) error {
	// TODO
	//client := m.(*api.Client)
	//ctx := context.Background()
	return nil
}