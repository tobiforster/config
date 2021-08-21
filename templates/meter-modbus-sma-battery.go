package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA Sunny Island / Sunny Boy Storage (Battery Meter)",
		Sample: `uri: 192.0.2.2:502
id: 126 # sma default sunspec modbus id
soc: ChargeState`,
	}

	registry.Add(template)
}
