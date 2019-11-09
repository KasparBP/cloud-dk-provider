package clouddk

import (
	"context"
	"fmt"
	"github.com/KasparBP/cloud-dk-provider/clouddk/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

const (
	csHostname            = "hostname"
	csLabel               = "label"
	csInitialRootPassword = "initial_root_password"
	csTemplate            = "template"
	csPackage             = "package"
	csLocation            = "location"
	csCpus                = "cpus"
	csMemory              = "memory"
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
		csLabel:    &schema.Schema{Type: schema.TypeString, Required: true},
		csCpus:     &schema.Schema{Type: schema.TypeInt, Optional: true, Computed: true},
		csMemory:   &schema.Schema{Type: schema.TypeInt, Optional: true, Computed: true},
		csInitialRootPassword: &schema.Schema{
			Type: schema.TypeString,
			Required: true,
			Sensitive: true,
			ForceNew: true,
			ValidateFunc: func(i interface{}, s string) (warns []string, errs []error) {
				if len(s) < 10 {
					errs = append(errs, fmt.Errorf("%s must be atleast 10 characters", csInitialRootPassword))
				}
				f := func(r rune) bool {
					return (r < 'A' || r > 'Z') && (r > 'z' || r < 'a')
				}
				if strings.IndexFunc(s, f) == -1 {
					errs = append(errs, fmt.Errorf("%s must contains numbers or special characters", csInitialRootPassword))
				}
				return
			},
		},
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
	_ = d.Set(csCpus, cloudServer.Cpus)
	_ = d.Set(csMemory, cloudServer.Memory)
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
	client := m.(*api.Client)
	ctx := context.Background()
	return client.ClouddkService.DeleteCloudServer(ctx, d.Id())
}