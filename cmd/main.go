package main

import (
	"LukeWinikates/january-twenty-five/lib/server"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
	"log"
	"os"
	"time"
)

func main() {
	// start the server
	// on each client message, update the device list
	// table for schedule entries

	client := zigbee2mqtt.NewClient(os.Getenv("MQTT_HOST"))
	dataPath := os.Getenv("DATA_PATH")
	s, err := server.New(client, dataPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.Start()

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
