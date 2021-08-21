package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "SENEC.Home (PV)",
		Sample: `power:
  source: script
  cmd: /bin/bash -c "curl -d '{\"ENERGY\":{\"GUI_INVERTER_POWER\":\"\"}}' -H \"Content-Type: application/json\" -X POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_INVERTER_POWER | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[3:]))[0])'"
  timeout: 5s`,
	}

	registry.Add(template)
}
