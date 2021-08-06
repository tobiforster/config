package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "custom",
		Name:   "evNotify (https://evnotify.de/)",
		Sample: `title: My Car # display name for UI
capacity: 39 # kWh
charge:
  type: http
  uri: https://app.evnotify.de/soc?akey=AKEY&token=1234567890abcdef # evNotify Server + AKEY
  method: GET
  jq: .soc_display
cache: 5m # cache duration`,
	}

	registry.Add(template)
}
