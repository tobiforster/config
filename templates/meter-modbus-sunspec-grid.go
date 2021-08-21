package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Generic SunSpec 3-phase meter via inverter (Grid Meter)",
		Sample: `uri: 192.0.2.2:502
id: 1
power: 203:W # sunspec model 203 meter`,
	}

	registry.Add(template)
}
