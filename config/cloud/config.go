package cloud

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "cloud"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("ovh_cloud_project", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "cloud"
        r.Kind = "Project"
    })
}
