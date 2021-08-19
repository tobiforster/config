package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Shelly 3EM (HTTP)",
		Sample: `power: # power reading
  source: http
  uri: http://192.0.2.2/emeter/0
  jq: .power`,
	}

	registry.Add(template)
}
