package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA SunnyBoy / TriPower / other PV-inverter (PV Meter)",
		Sample: `uri: 192.0.2.2:502
id: 126 # sma default sunspec modbus id`,
	}

	registry.Add(template)
}
