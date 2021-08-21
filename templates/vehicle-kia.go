package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "kia",
		Name:   "Kia (e-Niro, e-Soul, etc)",
		Sample: `title: e-Niro # display name for UI
capacity: 64 # kWh
user: # user
password: # password`,
	}

	registry.Add(template)
}
