package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "wallbe",
		Name:   "Wallbe (pre 2019)",
		Sample: `uri: 192.168.0.8:502 # TCP ModBus address
legacy: true # enable for older Wallbe devices (old controller firmware)
meter: # only if a charge meter is connected to the controller
  power: true
  energy: true
  currents: true`,
	}

	registry.Add(template)
}
