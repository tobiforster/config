package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Generic SunSpec battery inverter (Battery Meter)",
		Sample: `uri: 192.0.2.2:502
id: 1
soc: ChargeState`,
	}

	registry.Add(template)
}
