# Configuration examples for EVCC

[![Build Status](https://github.com/evcc-io/config/workflows/Build/badge.svg)](https://github.com/evcc-io/config/actions?query=workflow%3ABuild)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

[EVCC](https://github.com/andig/evcc) supports a growing list of [chargers](#chargers), [meters](#meters) and [vehicles](#vehicles). See below for detailed configuration.
Additional devices can be configured using the `generic` device type and related [plugins](#https://github.com/andig/evcc#plugins).

## Contributing

If you want to contribute configurations to this repository please open a Pull Request ("PR"). PRs should contain:

- updated or new `yaml` configuration (see https://github.com/evcc-io/config/tree/master/yaml)
- run `make` to generate compilable Go code (configurations are used for testing) and update the README

## Chargers

- [Easee Home (Cloud API)](#charger-easee-home-cloud-api)
- [EVSE DIN](#charger-evse-din)
- [EVSE-Wifi](#charger-evse-wifi)
- [FritzDECT](#charger-fritzdect)
- [Generic](#charger-generic)
- [go-eCharger](#charger-go-echarger)
- [go-eCharger (Cloud)](#charger-go-echarger-cloud)
- [Heidelberg Energy Control (Modbus RTU)](#charger-heidelberg-energy-control-modbus-rtu)
- [i-CHARGE CION (Modbus RTU-over-TCP)](#charger-i-charge-cion-modbus-rtu-over-tcp)
- [KEBA Connect](#charger-keba-connect)
- [Mobile Charger Connect (Audi, Bentley, Porsche)](#charger-mobile-charger-connect-audi-bentley-porsche)
- [NRGKick BT (Bluetooth)](#charger-nrgkick-bt-bluetooth)
- [NRGKick Connect](#charger-nrgkick-connect)
- [openWB (MQTT)](#charger-openwb-mqtt)
- [Phoenix EM-CP-PP-ETH Controller (Modbus TCP)](#charger-phoenix-em-cp-pp-eth-controller-modbus-tcp)
- [Phoenix EV-ETH Controller (Modbus TCP)](#charger-phoenix-ev-eth-controller-modbus-tcp)
- [Phoenix EV-SER Controller (Modbus RTU)](#charger-phoenix-ev-ser-controller-modbus-rtu)
- [Shelly](#charger-shelly)
- [Tasmota](#charger-tasmota)
- [TinkerForge WARP Charger](#charger-tinkerforge-warp-charger)
- [TP-LINK Smart Plug](#charger-tp-link-smart-plug)
- [Wallbe (Eco, Pro)](#charger-wallbe-eco-pro)

## Meters

- [Discovergy Metering Service (Cloud)](#meter-discovergy-metering-service-cloud)
- [E3DC (Battery Meter)](#meter-e3dc-battery-meter)
- [E3DC (Grid Meter)](#meter-e3dc-grid-meter)
- [E3DC (PV Meter)](#meter-e3dc-pv-meter)
- [Eastron SDM Modbus Meter](#meter-eastron-sdm-modbus-meter)
- [Fronius Solar API V1 (Battery Meter)](#meter-fronius-solar-api-v1-battery-meter)
- [Fronius Solar API V1 (Grid Meter)](#meter-fronius-solar-api-v1-grid-meter)
- [Fronius Solar API V1 (PV Meter)](#meter-fronius-solar-api-v1-pv-meter)
- [Fronius Symo GEN24 Plus (Battery Meter)](#meter-fronius-symo-gen24-plus-battery-meter)
- [Fronius Symo GEN24 Plus (Grid Meter)](#meter-fronius-symo-gen24-plus-grid-meter)
- [Fronius Symo GEN24 Plus (PV Meter)](#meter-fronius-symo-gen24-plus-pv-meter)
- [Generic](#meter-generic)
- [Kostal Energy Meter via inverter (Grid Meter)](#meter-kostal-energy-meter-via-inverter-grid-meter)
- [Kostal Hybrid Inverter (Battery Meter)](#meter-kostal-hybrid-inverter-battery-meter)
- [Kostal Inverter (PV Meter)](#meter-kostal-inverter-pv-meter)
- [Kostal Smart Energy Meter (Grid Meter)](#meter-kostal-smart-energy-meter-grid-meter)
- [Multiple DC MPP strings combined (PV Meter)](#meter-multiple-dc-mpp-strings-combined-pv-meter)
- [Multiple PV inverters combined (PV Meter)](#meter-multiple-pv-inverters-combined-pv-meter)
- [PowerDog (Grid Meter)](#meter-powerdog-grid-meter)
- [PowerDog (PV Meter)](#meter-powerdog-pv-meter)
- [Powerfox Poweropti (Cloud)](#meter-powerfox-poweropti-cloud)
- [RCT Power Storage (Battery)](#meter-rct-power-storage-battery)
- [RCT Power Storage (Grid)](#meter-rct-power-storage-grid)
- [RCT Power Storage (PV)](#meter-rct-power-storage-pv)
- [SENEC.Home (Battery)](#meter-senec-home-battery)
- [SENEC.Home (Grid)](#meter-senec-home-grid)
- [SENEC.Home (PV)](#meter-senec-home-pv)
- [Shelly 3EM (HTTP)](#meter-shelly-3em-http)
- [SMA Sunny Home Manager / Energy Meter (Speedwire)](#meter-sma-sunny-home-manager--energy-meter-speedwire)
- [SMA Sunny Island / Sunny Boy Storage (Battery Meter)](#meter-sma-sunny-island--sunny-boy-storage-battery-meter)
- [SMA SunnyBoy / TriPower / other PV-inverter (PV Meter)](#meter-sma-sunnyboy--tripower--other-pv-inverter-pv-meter)
- [SolarEdge Energy Meter via inverter (Grid Meter)](#meter-solaredge-energy-meter-via-inverter-grid-meter)
- [SolarEdge Hybrid Inverter (PV Meter)](#meter-solaredge-hybrid-inverter-pv-meter)
- [SolarEdge StorEdge (Battery Meter)](#meter-solaredge-storedge-battery-meter)
- [Solarlog (Grid Meter)](#meter-solarlog-grid-meter)
- [Solarlog (PV Meter)](#meter-solarlog-pv-meter)
- [Solarwatt MyReserve (Battery Meter)](#meter-solarwatt-myreserve-battery-meter)
- [Solarwatt MyReserve (Grid Meter)](#meter-solarwatt-myreserve-grid-meter)
- [Solarwatt MyReserve (PV Meter)](#meter-solarwatt-myreserve-pv-meter)
- [Sonnenbatterie Eco/10 (Battery Meter)](#meter-sonnenbatterie-eco-10-battery-meter)
- [Sonnenbatterie Eco/10 (Grid Meter)](#meter-sonnenbatterie-eco-10-grid-meter)
- [Sonnenbatterie Eco/10 (PV Meter)](#meter-sonnenbatterie-eco-10-pv-meter)
- [SunSpec compliant 3-phase meter via inverter (Grid Meter)](#meter-sunspec-compliant-3-phase-meter-via-inverter-grid-meter)
- [SunSpec compliant battery inverter (Battery Meter)](#meter-sunspec-compliant-battery-inverter-battery-meter)
- [SunSpec compliant PV inverter (PV Meter)](#meter-sunspec-compliant-pv-inverter-pv-meter)
- [Tasmota (HTTP)](#meter-tasmota-http)
- [Tesla Powerwall (Battery Meter)](#meter-tesla-powerwall-battery-meter)
- [Tesla Powerwall (Grid Meter)](#meter-tesla-powerwall-grid-meter)
- [Tesla Powerwall (PV Meter)](#meter-tesla-powerwall-pv-meter)
- [vzlogger (HTTP)](#meter-vzlogger-http)
- [vzlogger (Push Server)](#meter-vzlogger-push-server)
- [vzlogger (split import/export channels)](#meter-vzlogger-split-import-export-channels)

## Vehicles

- [Audi (eTron etc)](#vehicle-audi-etron-etc)
- [BMW (i3)](#vehicle-bmw-i3)
- [Citroen](#vehicle-citroen)
- [evNotify (HTTP)](#vehicle-evnotify-http)
- [Ford (Kuga, Mustang, etc)](#vehicle-ford-kuga-mustang-etc)
- [Generic](#vehicle-generic)
- [Hyundai (Kona, Ioniq)](#vehicle-hyundai-kona-ioniq)
- [Kia (e-Niro, e-Soul, etc)](#vehicle-kia-e-niro-e-soul-etc)
- [Nissan (Leaf)](#vehicle-nissan-leaf)
- [NIU E-Scooter](#vehicle-niu-e-scooter)
- [Opel](#vehicle-opel)
- [OVMS](#vehicle-ovms)
- [Peugeot](#vehicle-peugeot)
- [Porsche](#vehicle-porsche)
- [Renault (Zoe)](#vehicle-renault-zoe)
- [Tesla](#vehicle-tesla)
- [VW (e-Up, e-Golf, etc)](#vehicle-vw-e-up-e-golf-etc)
- [VW ID (ID.3, ID.4, but also e-Golf, e-Up)](#vehicle-vw-id-id-3-id-4-but-also-e-golf-e-up)

## Details

### Meters


<a id="meter-discovergy-metering-service-cloud"></a>
#### Discovergy Metering Service (Cloud)

```yaml
- type: discovergy
  user: demo@discovergy.com 
  password: demo # password 
  meter: 1ESY1161229886
```

<a id="meter-e3dc-battery-meter"></a>
#### E3DC (Battery Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration for E3/DC "Simple-Mode"
      address: 40069 # Batterie-Leistung in Watt
      type: holding
      decode: int32s
    scale: -1 # reverse direction
  soc:
    source: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration for E3/DC "Simple-Mode"
      address: 40082 # Batterie-SOC in Prozent
      type: holding
      decode: uint16
```

<a id="meter-e3dc-grid-meter"></a>
#### E3DC (Grid Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration for E3/DC "Simple-Mode"
      address: 40073 # Hausverbrauchs-Leistung in Watt
      type: holding
      decode: int32s
```

<a id="meter-e3dc-pv-meter"></a>
#### E3DC (PV Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration for E3/DC "Simple-Mode"
      address: 40067 # Photovoltaikleistung in Watt
      type: holding
      decode: int32s
```

<a id="meter-eastron-sdm-modbus-meter"></a>
#### Eastron SDM Modbus Meter

```yaml
- type: modbus
  model: sdm # specific non-sunspec meter
  id: 1
  energy: Sum # only required for charge meter usage
  # chose either locally attached:
  device: /dev/ttyUSB0 # serial port
  baudrate: 9600
  comset: 8N1
  # or via TCP:
  uri: 192.0.2.2:502
  rtu: true # serial modbus rtu (rs485) device connected using simple ethernet adapter
```

<a id="meter-fronius-solar-api-v1-battery-meter"></a>
#### Fronius Solar API V1 (Battery Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_Akku == null then 0 else .Body.Data.Site.P_Akku end
  soc:
    source: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: .Body.Data.Inverters."1".SOC
```

<a id="meter-fronius-solar-api-v1-grid-meter"></a>
#### Fronius Solar API V1 (Grid Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_Grid == null then 0 else .Body.Data.Site.P_Grid end
```

<a id="meter-fronius-solar-api-v1-pv-meter"></a>
#### Fronius Solar API V1 (PV Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
    jq: if .Body.Data.Site.P_PV == null then 0 else .Body.Data.Site.P_PV end
```

<a id="meter-fronius-symo-gen24-plus-battery-meter"></a>
#### Fronius Symo GEN24 Plus (Battery Meter)

```yaml
- type: custom
  power:
    source: calc
    add:
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:3:DCW # mppt 3 charge
      scale: -1
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:4:DCW # mppt 4 discharge
  soc:
    source: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: ChargeState
```

<a id="meter-fronius-symo-gen24-plus-grid-meter"></a>
#### Fronius Symo GEN24 Plus (Grid Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 200
  power: 213:W # sunspec model 203 meter
```

<a id="meter-fronius-symo-gen24-plus-pv-meter"></a>
#### Fronius Symo GEN24 Plus (PV Meter)

```yaml
- type: custom
  power:
    source: calc
    add:
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:1:DCW # mpp 1 pv
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
      value: 160:2:DCW # mpp 2 pv
```

<a id="meter-generic"></a>
#### Generic

```yaml
- type: custom
  power: # power (W)
    source: # plugin type
    # ...
  energy: # optional energy (kWh)
    source: # plugin type
    # ...
  soc: # optional battery soc (%)
    source: # plugin type
    # ...
  currents: # optional currents (A)
    - source: # L1 plugin type
      # ...
    - source: # L2 plugin type
      # ...
    - source: # L3 plugin type
      # ...
```

<a id="meter-kostal-energy-meter-via-inverter-grid-meter"></a>
#### Kostal Energy Meter via inverter (Grid Meter)

```yaml
- type: custom
  power:
    source: modbus # use ModBus plugin
    uri: 192.0.2.2:1502 # inverter port
    id: 71
    register: # manual non-sunspec register configuration
      address: 252 # (see ba_kostal_interface_modbus-tcp_sunspec.pdf)
      type: holding
      decode: float32s # may be float32 on specific firmware/devices
```

<a id="meter-kostal-hybrid-inverter-battery-meter"></a>
#### Kostal Hybrid Inverter (Battery Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:1502
  id: 71 # kostal default sunspec modbus id
  power: 802:W # sunspec model 802 battery
  soc: 802:SoC # sunspec model 802 battery
```

<a id="meter-kostal-inverter-pv-meter"></a>
#### Kostal Inverter (PV Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:1502
  id: 71 # kostal default sunspec modbus id
```

<a id="meter-kostal-smart-energy-meter-grid-meter"></a>
#### Kostal Smart Energy Meter (Grid Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 71 # kostal default sunspec modbus id
```

<a id="meter-multiple-dc-mpp-strings-combined-pv-meter"></a>
#### Multiple DC MPP strings combined (PV Meter)

```yaml
- type: custom
  power:
    source: calc
    add:
    - source: modbus
      model: sunspec
      value: 160:1:DCW # SunSpec Model 160 MPP string 1 DCW
      uri: 192.0.2.2:502
      id: 1
    - source: modbus
      model: sunspec
      value: 160:2:DCW # SunSpec Model 160 MPP string 2 DCW
      uri: 192.0.2.2:502
      id: 1
```

<a id="meter-multiple-pv-inverters-combined-pv-meter"></a>
#### Multiple PV inverters combined (PV Meter)

```yaml
- type: custom
  power:
    source: calc
    add:
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502
      id: 1
    - source: modbus
      model: sunspec
      uri: 192.0.2.3:502
      id: 1
```

<a id="meter-powerdog-grid-meter"></a>
#### PowerDog (Grid Meter)

```yaml
- type: custom
  power:
    source: calc #calculate current overall consumption + (current pv effort * (-1) )
    add:
      - source: modbus
        uri: 192.168.1.2:502 #ip-adress and port (default-port: 502)
        id: 1
        register:
          address: 40026 #register for overall consumption
          type: holding
          decode: int32
  
      - source: modbus
        uri: 192.168.1.2:502 #ip-adress and port (default-port: 502)
        id: 1
        register:
          address: 40002 #register for pv effort
          type: holding
          decode: int32
        scale: -1 #scale with -1 to get a substraction
```

<a id="meter-powerdog-pv-meter"></a>
#### PowerDog (PV Meter)

```yaml
- type: custom
  power:
    type: modbus
    uri: 192.168.1.2:502 #ip-adress and port (default-port: 502)
    id: 1
    register:
      address: 40002 #register for pv effort
      type: holding
      decode: int32
```

<a id="meter-powerfox-poweropti-cloud"></a>
#### Powerfox Poweropti (Cloud)

```yaml
- type: custom
  power:
    source: http
    uri: https://backend.powerfox.energy/api/2.0/my/main/current
    auth:
      type: basic
      user: xxxxxxxxx
      password: *****
    jq: .Watt
```

<a id="meter-rct-power-storage-battery"></a>
#### RCT Power Storage (Battery)

```yaml
- type: custom
  power:
    source: script
    cmd: /bin/bash -c "rctclient read-value --host 192.0.2.2 --name g_sync.p_acc_lp"
    timeout: 5s
  soc:
    source: script
    cmd: /bin/bash -c "echo $(rctclient read-value --host 192.0.2.2 --name battery.soc) \* 100. | bc -l"
    timeout: 5s
```

<a id="meter-rct-power-storage-grid"></a>
#### RCT Power Storage (Grid)

```yaml
- type: custom
  power:
    source: script
    cmd: /bin/bash -c "rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_grid_sum_lp"
    timeout: 5s
```

<a id="meter-rct-power-storage-pv"></a>
#### RCT Power Storage (PV)

```yaml
- type: custom
  power:
    source: script
    cmd: /bin/bash -c "echo $(rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_load_sum_lp) \- $(rctclient read-value --host 192.0.2.2 --name g_sync.p_acc_lp) \- $(rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_grid_sum_lp) | bc -l"
    timeout: 5s
```

<a id="meter-senec-home-battery"></a>
#### SENEC.Home (Battery)

```yaml
- type: custom
  power:
    source: script
    cmd: >
      /bin/bash -c "set +H; curl --data '{\"ENERGY\":{\"GUI_BAT_DATA_POWER\":\"\"}}' --header \"Content-Type: application/json\" --request POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_BAT_DATA_POWER | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[4:12]))[0])'"
    timeout: 5s
    scale: -1
  soc:
    source: script
    cmd: >
      /bin/bash -c "set +H; curl --data '{\"ENERGY\":{\"GUI_BAT_DATA_FUEL_CHARGE\":\"\"}}' --header \"Content-Type: application/json\" --request POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_BAT_DATA_FUEL_CHARGE | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[4:12]))[0])'"
    timeout: 5s
```

<a id="meter-senec-home-grid"></a>
#### SENEC.Home (Grid)

```yaml
- type: custom
  power:
    cmd: >
      /bin/bash -c "set +H; curl --data '{\"ENERGY\":{\"GUI_GRID_POW\":\"\"}}' --header \"Content-Type: application/json\" --request POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_GRID_POW | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[4:12]))[0])'"
    timeout: 5s
    scale: -1
```

<a id="meter-senec-home-pv"></a>
#### SENEC.Home (PV)

```yaml
- type: custom
  power:
    source: script
    cmd: >
      /bin/bash -c "set +H; curl --data '{\"ENERGY\":{\"GUI_INVERTER_POWER\":\"\"}}' --header \"Content-Type: application/json\" --request POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_INVERTER_POWER | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[4:12]))[0])'"
    timeout: 5s
```

<a id="meter-shelly-3em-http"></a>
#### Shelly 3EM (HTTP)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/status
    jq: .emeters | map(.power) | add
  energy:
    source: http
    uri: http://192.0.2.2/status
    jq: .emeters | map(.total) | add
    scale: 0.001
  currents:
  - source: http
    uri: http://192.0.2.2/emeter/0
    jq: .current
  - source: http
    uri: http://192.0.2.2/emeter/1
    jq: .current
  - source: http
    uri: http://192.0.2.2/emeter/2
    jq: .current
```

<a id="meter-sma-sunny-home-manager--energy-meter-speedwire"></a>
#### SMA Sunny Home Manager / Energy Meter (Speedwire)

```yaml
- type: sma
  uri: 192.0.2.2
```

<a id="meter-sma-sunny-island--sunny-boy-storage-battery-meter"></a>
#### SMA Sunny Island / Sunny Boy Storage (Battery Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 126 # sma default sunspec modbus id
  soc: ChargeState
```

<a id="meter-sma-sunnyboy--tripower--other-pv-inverter-pv-meter"></a>
#### SMA SunnyBoy / TriPower / other PV-inverter (PV Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 126 # sma default sunspec modbus id
```

<a id="meter-solaredge-energy-meter-via-inverter-grid-meter"></a>
#### SolarEdge Energy Meter via inverter (Grid Meter)

```yaml
- type: custom
  power:
    source: modbus
    model: sunspec
    uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
    id: 1
    subdevice: 1 # Metering device
    value: 203:W
    scale: -1
```

<a id="meter-solaredge-hybrid-inverter-pv-meter"></a>
#### SolarEdge Hybrid Inverter (PV Meter)

```yaml
- type: custom
  power:
    source: calc
    add:
    - source: modbus
      model: sunspec
      uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
      id: 1
      value: 103:DCW
    - source: modbus
      uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
      id: 1
      register:
        address: 62836 # Battery 1 Instantaneous Power
        type: holding
        decode: float32s
```

<a id="meter-solaredge-storedge-battery-meter"></a>
#### SolarEdge StorEdge (Battery Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
    id: 1
    register:
      address: 62836 # Battery 1 Instantaneous Power
      type: holding
      decode: float32s
    scale: -1
  soc:
    source: modbus
    uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
    id: 1
    register:
      address: 62852 # Battery 1 State of Energy (SOE)
      type: holding
      decode: float32s
```

<a id="meter-solarlog-grid-meter"></a>
#### Solarlog (Grid Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: 192.0.2.2:502 # IP address of the SolarLog device and ModBus port address
    id: 1
    register:
      address: 3518
      type: input
      decode: uint32s
```

<a id="meter-solarlog-pv-meter"></a>
#### Solarlog (PV Meter)

```yaml
- type: custom
  power:
    source: modbus
    uri: 192.0.2.2:502 # IP address of the SolarLog  device and ModBus port address
    id: 1
    register:
      address: 3502
      type: input
      decode: uint32s
```

<a id="meter-solarwatt-myreserve-battery-meter"></a>
#### Solarwatt MyReserve (Battery Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/rest/kiwigrid/wizard/devices # EnergyManager
    jq: .result.items[] | select(.deviceModel[].deviceClass == "com.kiwigrid.devices.location.Location" ) | .tagValues.PowerConsumedFromStorage.value - .tagValues.PowerBuffered.value
  soc:
    source: http
    uri: http://192.0.2.2/rest/kiwigrid/wizard/devices # EnergyManager
    jq: .result.items[] | select(.deviceModel[].deviceClass == "com.kiwigrid.devices.solarwatt.MyReserve") | .tagValues.StateOfCharge.value
```

<a id="meter-solarwatt-myreserve-grid-meter"></a>
#### Solarwatt MyReserve (Grid Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/rest/kiwigrid/wizard/devices # EnergyManager
    jq: .result.items[] | select(.deviceModel[].deviceClass == "com.kiwigrid.devices.location.Location" ) | .tagValues.PowerConsumedFromGrid.value - .tagValues.PowerOut.value
```

<a id="meter-solarwatt-myreserve-pv-meter"></a>
#### Solarwatt MyReserve (PV Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2/rest/kiwigrid/wizard/devices # EnergyManager
    jq: .result.items[] | select(.deviceModel[].deviceClass == "com.kiwigrid.devices.location.Location" ) | .tagValues.PowerProduced.value
```

<a id="meter-sonnenbatterie-eco-10-battery-meter"></a>
#### Sonnenbatterie Eco/10 (Battery Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Pac_total_W
  soc:
    source: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .USOC
```

<a id="meter-sonnenbatterie-eco-10-grid-meter"></a>
#### Sonnenbatterie Eco/10 (Grid Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .GridFeedIn_W
    scale: -1 # reverse direction
```

<a id="meter-sonnenbatterie-eco-10-pv-meter"></a>
#### Sonnenbatterie Eco/10 (PV Meter)

```yaml
- type: custom
  power:
    source: http
    uri: http://192.0.2.2:8080/api/v1/status
    jq: .Production_W
```

<a id="meter-sunspec-compliant-3-phase-meter-via-inverter-grid-meter"></a>
#### SunSpec compliant 3-phase meter via inverter (Grid Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 1
  power: 203:W # sunspec model 203 meter
```

<a id="meter-sunspec-compliant-battery-inverter-battery-meter"></a>
#### SunSpec compliant battery inverter (Battery Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 1
  soc: ChargeState
```

<a id="meter-sunspec-compliant-pv-inverter-pv-meter"></a>
#### SunSpec compliant PV inverter (PV Meter)

```yaml
- type: modbus
  model: sunspec
  uri: 192.0.2.2:502
  id: 1
```

<a id="meter-tasmota-http"></a>
#### Tasmota (HTTP)

```yaml
- type: custom
  power: # power reading (W)
    source: http
    uri: http://192.0.2.2/cm?cmnd=Status%208
    jq: .StatusSNS.ENERGY.Power
  energy: # energy reading (Wh), for chargemeter usage only
    source: http
    uri: http://192.0.2.2/cm?cmnd=Status%208
    jq: .StatusSNS.ENERGY.Total * 1000
```

<a id="meter-tesla-powerwall-battery-meter"></a>
#### Tesla Powerwall (Battery Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: battery
```

<a id="meter-tesla-powerwall-grid-meter"></a>
#### Tesla Powerwall (Grid Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: grid
```

<a id="meter-tesla-powerwall-pv-meter"></a>
#### Tesla Powerwall (PV Meter)

```yaml
- type: tesla
  uri: http://192.0.2.2/
  usage: pv
```

<a id="meter-vzlogger-http"></a>
#### vzlogger (HTTP)

```yaml
- type: custom
  power: # power reading
    source: http # use http plugin
    uri: http://demo.volkszaehler.org/api/data/<uuid>.json?from=now
    jq: .data.tuples[0][1] # parse response json
```

<a id="meter-vzlogger-push-server"></a>
#### vzlogger (Push Server)

```yaml
- type: custom
  power:
    source: ws # use websocket plugin
    uri: ws://192.0.2.2:8082/socket
    jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse response json
    timeout: 30s
    scale: 1
```

<a id="meter-vzlogger-split-import-export-channels"></a>
#### vzlogger (split import/export channels)

```yaml
- type: custom
  power:
    source: calc # use calc plugin
    add:
    - source: http # import channel
      uri: http://demo.volkszaehler.org/api/data/<import-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
    - source: http # export channel
      uri: http://demo.volkszaehler.org/api/data/<export-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
      scale: -1 # export must result in negative values
```


### Chargers


<a id="charger-easee-home-cloud-api"></a>
#### Easee Home (Cloud API)

```yaml
- type: easee
  user: foo@example.org
  password: *****
  charger: EH______
  cache: 10s
```

<a id="charger-evse-din"></a>
#### EVSE DIN

```yaml
- type: simpleevse
  # http://evracing.cz/simple-evse-wallbox
  # either locally attached:
  device: /dev/ttyUSB0 # serial RS485 interface
  # or via TCP:
  uri: 192.0.2.2:502 # Modbus TCP converter address
```

<a id="charger-evse-wifi"></a>
#### EVSE-Wifi

```yaml
- type: evsewifi
  uri: http://192.0.2.2
```

<a id="charger-fritzdect"></a>
#### FritzDECT

```yaml
- type: fritzdect
  uri: https://fritz.box # FRITZ!Box ip address (local)
  user: xxxxxxxxxx # FRITZ!Box username (Has to have Smart Home privileges!)
  password: yyyyyyyyyy # FRITZ!Box password
  ain: '007788992233' # switch actor identification number without blanks (see AIN number on switch sticker)
  standbypower: 15 # treat as charging above this power
```

<a id="charger-generic"></a>
#### Generic

```yaml
- type: custom
  status: # charger status A..F
    source: ...
    # ...
  enabled: # charger enabled state (true/false or 0/1)
    source: ...
    # ...
  enable: # set charger enabled state (true/false or 0/1)
    source: ...
    # ...
  maxcurrent: # set charger max current (A)
    source: ...
    # ...
```

<a id="charger-go-echarger"></a>
#### go-eCharger

```yaml
- type: go-e
  uri: http://192.0.2.2 # go-e ip address (local)
```

<a id="charger-go-echarger-cloud"></a>
#### go-eCharger (Cloud)

```yaml
- type: go-e
  token: 4711c # go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

<a id="charger-heidelberg-energy-control-modbus-rtu"></a>
#### Heidelberg Energy Control (Modbus RTU)

```yaml
- type: heidelberg
  device: /dev/ttyUSB0
  baudrate: 19200
  comset: 8E1
  id: 1 # configurable (S2/DIP 1)
```

<a id="charger-i-charge-cion-modbus-rtu-over-tcp"></a>
#### i-CHARGE CION (Modbus RTU-over-TCP)

```yaml
- type: custom
  status:
    source: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
        address: 139 # CP-Status
        type: holding
        decode: uint16
  enabled:
    source: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1 
    register: # manual register configuration
      address: 100 # Zustand
      type: holding
      decode: uint16
  enable:
    source: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
      address: 100 # ein / aus
      type: writesingle
      decode: uint16
  maxcurrent:
    source: modbus
    uri: 192.0.2.2:502
    rtu: true # Modbus over TCP
    id: 1
    register: # manual register configuration
      address: 127 # Strom max
      type: writesingle
      decode: uint16
```

<a id="charger-keba-connect"></a>
#### KEBA Connect

```yaml
- type: keba
  uri: 192.0.2.2
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

<a id="charger-mobile-charger-connect-audi-bentley-porsche"></a>
#### Mobile Charger Connect (Audi, Bentley, Porsche)

```yaml
- type: mcc
  uri: https://192.0.2.2
  password: # home user password
```

<a id="charger-nrgkick-bt-bluetooth"></a>
#### NRGKick BT (Bluetooth)

```yaml
- type: nrgkick-bluetooth
  mac: 00:1E:C0:XX:XX:XX # BT device MAC address (sudo hcitool lescan)
  pin: 1234 # App PIN number (write protection, ignore leading zeros)
```

<a id="charger-nrgkick-connect"></a>
#### NRGKick Connect

```yaml
- type: nrgkick-connect
  uri: http://192.0.2.2
  mac: 00:1E:C0:XX:XX:XX # BT device MAC address (sudo hcitool lescan)
  password: # password
```

<a id="charger-openwb-mqtt"></a>
#### openWB (MQTT)

```yaml
- type: openwb
  broker: 192.0.2.2 # openWB IP
  id: 1 # loadpoint id
```

<a id="charger-phoenix-em-cp-pp-eth-controller-modbus-tcp"></a>
#### Phoenix EM-CP-PP-ETH Controller (Modbus TCP)

```yaml
- type: phoenix-em-eth
  uri: 192.168.0.8:502
  meter: # only if a charge meter is connected to the controller
    power: true
    energy: true
    currents: true
```

<a id="charger-phoenix-ev-eth-controller-modbus-tcp"></a>
#### Phoenix EV-ETH Controller (Modbus TCP)

```yaml
- type: phoenix-ev-eth
  uri: 192.168.0.8:502
  meter: # only if a charge meter is connected to the controller
    power: true
    energy: true
    currents: true
```

<a id="charger-phoenix-ev-ser-controller-modbus-rtu"></a>
#### Phoenix EV-SER Controller (Modbus RTU)

```yaml
- type: phoenix-ev-ser
  device: /dev/ttyUSB0
  baudrate: 9600 # configurable (S2/DIP 1)
  comset: 8N1
  id: 1 # configurable (S2/DIP 2–6)
```

<a id="charger-shelly"></a>
#### Shelly

```yaml
- type: shelly
  uri: http://192.168.xxx.xxx  # shelly device ip address (local)
  channel: 0  # shelly device relay channel 
  standbypower: 15  # treat as charging above this power
```

<a id="charger-tasmota"></a>
#### Tasmota

```yaml
- type: tasmota
  uri: http://192.168.xxx.xxx # tasmota device ip address (local)
  # user: xxxx # user, (optional) in case user + password are defined
  # password: xxxxx #  (optional) in case user + password are defined
  standbypower: 15 # treat as charging above this power
```

<a id="charger-tinkerforge-warp-charger"></a>
#### TinkerForge WARP Charger

```yaml
- type: warp
  broker: 192.0.2.2:1883
  topic: warp
  useMeter: true # WARP Charger Pro
  timeout: 30s
```

<a id="charger-tp-link-smart-plug"></a>
#### TP-LINK Smart Plug

```yaml
- type: tplink
  uri: 192.0.2.2 # TP-LINK Smart Plug ip address (local)
  standbypower: 15 # treat as charging above this power
```

<a id="charger-wallbe-eco-pro"></a>
#### Wallbe (Eco, Pro)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # TCP ModBus address
  legacy: true # set only for older Wallbe devices (pre ~2019, old controller firmware)  
  meter: # only if a charge meter is connected to the controller
    power: true
    energy: true
    currents: true
    encoding: sdm # add only when SDM meter is connected, see https://github.com/andig/evcc/discussions/1398
```


### Vehicles


<a id="vehicle-audi-etron-etc"></a>
#### Audi (eTron etc)

```yaml
- type: audi
  title: eTron # display name for UI
  capacity: 14 # kWh
  user: # user
  password: # password
  vin: WAUZZZ... # optional
```

<a id="vehicle-bmw-i3"></a>
#### BMW (i3)

```yaml
- type: bmw
  title: i3 # display name for UI
  capacity: 65 # kWh
  user: # user
  password: # password
  vin: WBMW... # optional
```

<a id="vehicle-citroen"></a>
#### Citroen

```yaml
- type: citroen
  title: e-C4 # display name for UI
  capacity: 50 # kWh
  user: user@example.com
  password: xxx
  vin: # optional
```

<a id="vehicle-evnotify-http"></a>
#### evNotify (HTTP)

```yaml
- type: custom
  title: My Car # display name for UI
  capacity: 64 # kWh
  charge:
    type: http
    uri: https://app.evnotify.de/soc?akey=AKEY&token=1234567890abcdef # evNotify Server + AKEY
    method: GET
    jq: .soc_display
  cache: 5m # cache duration
```

<a id="vehicle-ford-kuga-mustang-etc"></a>
#### Ford (Kuga, Mustang, etc)

```yaml
- type: ford
  title: Kuga # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WF0FXX... # optional
```

<a id="vehicle-generic"></a>
#### Generic

```yaml
- type: custom
  title: Mein Auto # display name for UI
  capacity: 50 # byttery capacity (kWh)
  charge: # battery soc (%)
    source: # plugin type
    # ...
  status: # optional charge status (A..F)
    source: # plugin type
    # ...
  range: # optional electric range (km)
    source: # plugin type
    # ...
  cache: 5m # optional cache duration
```

<a id="vehicle-hyundai-kona-ioniq"></a>
#### Hyundai (Kona, Ioniq)

```yaml
- type: hyundai
  title: Kona # display name for UI
  capacity: 64 # kWh
  user: # user
  password: # password
```

<a id="vehicle-kia-e-niro-e-soul-etc"></a>
#### Kia (e-Niro, e-Soul, etc)

```yaml
- type: kia
  title: e-Niro # display name for UI
  capacity: 64 # kWh
  user: # user
  password: # password
```

<a id="vehicle-nissan-leaf"></a>
#### Nissan (Leaf)

```yaml
- type: nissan
  title: Leaf # display name for UI
  capacity: 60 # kWh
  user: # user
  password: # password
```

<a id="vehicle-niu-e-scooter"></a>
#### NIU E-Scooter

```yaml
- type: niu
  title: NIU E-Scooter # display name for UI
  capacity: 4 # kWh
  user: xxxxxxx # NIU app user
  password: xxxxxx # NIU app password
  serial: NXXXXXXXXXXXXXXX # NIU E-Scooter serial number like shown in app
```

<a id="vehicle-opel"></a>
#### Opel

```yaml
- type: opel
  title: Corsa-e # display name for UI
  capacity: 50 # kWh
  user: user@example.com
  password: xxx
  vin: # optional
```

<a id="vehicle-ovms"></a>
#### OVMS

```yaml
- type: ovms
  title: Open Vehicle Monitoring System # display name for UI
  capacity: 12 # kWh
  user: # user server
  password: # password server
  vehicleid: # vehicle id
  server: dexters-web.de # used ovms server [dexters-web.de or api.openvehicles.com]
```

<a id="vehicle-peugeot"></a>
#### Peugeot

```yaml
- type: peugeot
  title: e-208 # display name for UI
  capacity: 50 # kWh
  user: user@example.com
  password: xxx
  vin: # optional
```

<a id="vehicle-porsche"></a>
#### Porsche

```yaml
- type: porsche
  title: Taycan # display name for UI
  capacity: 83 # kWh
  user: # user
  password: # password
  vin: WP...
```

<a id="vehicle-renault-zoe"></a>
#### Renault (Zoe)

```yaml
- type: renault
  title: Zoe # display name for UI
  capacity: 60 # kWh
  user: # user
  password: # password
  vin: WREN... # optional
```

<a id="vehicle-tesla"></a>
#### Tesla

```yaml
- type: tesla
  title: Model S # display name for UI
  capacity: 90 # kWh
  user: # email
  password: # password
  vin: WTSLA...
```

<a id="vehicle-vw-e-up-e-golf-etc"></a>
#### VW (e-Up, e-Golf, etc)

```yaml
- type: vw
  title: Golf # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WVWZZZ... # optional
```

<a id="vehicle-vw-id-id-3-id-4-but-also-e-golf-e-up"></a>
#### VW ID (ID.3, ID.4, but also e-Golf, e-Up)

```yaml
- type: id
  title: ID.3 # display name for UI
  capacity: 50 # kWh
  user: # user
  password: # password
  vin: WVWZZZ... # optional
```

