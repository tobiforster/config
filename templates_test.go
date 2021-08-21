package main

import (
	"testing"

	"github.com/evcc-io/config/registry"
	_ "github.com/evcc-io/config/templates"

	"github.com/andig/evcc/charger"
	"github.com/andig/evcc/meter"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/vehicle"
	"gopkg.in/yaml.v2"
)

func TestTemplates(t *testing.T) {
	log := util.NewLogger("foo")
	// log.SetLogThreshold(jww.Th)

	for _, template := range registry.Registry {
		var conf map[string]interface{}
		if err := yaml.Unmarshal([]byte(template.Sample), &conf); err != nil {
			panic(err)
		}

		switch template.Class {
		case "meter":
			meter.NewConfigurableFromConfig(log, conf)
		case "charger":
			charger.NewConfigurableFromConfig(log, conf)
		case "vehicle":
			vehicle.NewConfigurableFromConfig(log, conf)
		}
	}
}
