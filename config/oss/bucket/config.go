package oss

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oss_bucket", func(r *config.Resource) {
		r.ShortGroup = "oss"
	})
}
