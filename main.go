package main

import (
	"github.com/KasparBP/cloud-dk-provider/clouddk"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts {
		ProviderFunc: clouddk.NewProvider,
	})
}
