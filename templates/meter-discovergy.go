package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "discovergy",
		Name:   "Discovergy",
		Sample: `user: demo@discovergy.com 
password: demo # password 
meter: 659a3da00324400da66cef81e1cbe3c5`,
	}

	registry.Add(template)
}
