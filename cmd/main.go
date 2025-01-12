package main

import (
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt/devices"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt/payloads"
	"fmt"
	"os"
	"time"
)

func main() {
	client := zigbee2mqtt.NewClient(os.Getenv("MQTT_HOST"))

	client.SubscribeDeviceCatalog(func(devices []payloads.MessagePayload) {
		for _, device := range devices {
			fmt.Println(device.FriendlyName)
		}
	})

	client.SetDeviceState(os.Getenv("TEST_DEVICE_NAME"), devices.OnMessage())

	for {
		time.Sleep(20 * time.Second)
	}
}
