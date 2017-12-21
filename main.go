package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/kaarolch/terraform-provider-grafana/grafana"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: grafana.Provider})
}
