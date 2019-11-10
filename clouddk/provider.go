package clouddk

import (
	"github.com/KasparBP/cloud-dk-provider/clouddk/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	tokenKey = "access_token"
	endpointKey = "endpoint"
)

func NewProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: providerSchema(),
		ResourcesMap: providerResourcesMap(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema {
		tokenKey: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc("CLOUDDK_TOKEN", ""),
			Description: "API Access token",
		},
		endpointKey: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("CLOUDDK_ENDPOINT", ""),
		},
	}
}

func providerResourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource {
		"clouddk_server": ResourceCloudServer(),
		"clouddk_disk": ResourceDisk(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return api.NewClient(
		d.Get(tokenKey).(string),
		d.Get(endpointKey).(string))
}
