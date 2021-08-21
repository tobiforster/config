package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "SENEC.Home (Grid)",
		Sample: `power:
  cmd: /bin/bash -c "curl -d '{\"ENERGY\":{\"GUI_GRID_POW\":\"\"}}' -H \"Content-Type: application/json\" -X POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_GRID_POW | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[3:]))[0])'"
  timeout: 5s
  scale: -1`,
	}

	registry.Add(template)
}
