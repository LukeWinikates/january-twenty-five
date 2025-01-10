package main

import (
	"LukeWinikates/january-twenty-five/lib/devices"
	"LukeWinikates/january-twenty-five/lib/payloads"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	//subscribeLightSignal(client)
	//sub(client)

	//turnOnLight(client, os.Getenv("TEST_DEVICE_NAME"))

	for {
		time.Sleep(20 * time.Second)
	}
	//client.Disconnect(2500)
}

func turnOnLight(client mqtt.Client, deviceName string) {
	message, err := json.Marshal(devices.OnMessage())
	if err != nil {
		fmt.Println(err)
	}

	topic := fmt.Sprintf("zigbee2mqtt/%s/set", deviceName)
	publish := client.Publish(topic, 0, false, message)
	if publish.Error() != nil {
		fmt.Println(publish.Error().Error())
	}
}

func subscribeLightSignal(client mqtt.Client) {
	topic := "zigbee2mqtt/Office Work Desk/set"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
		fmt.Printf("%v", message)
	})
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

//func subscribeLogging(client mqtt.Client) {
//	topic := "zigbee2mqtt/bridge/logging"
//	token := client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
//		//fmt.Printf("Received message from topic: %s\n", message.Topic())
//		////err := os.WriteFile("devices.json", message.Payload(), os.ModePerm)
//		//if err != nil {
//		//	log.Default().Printf("err: %s", err.Error())
//		//}
//	})
//	token.Wait()
//	fmt.Printf("Subscribed to topic %s", topic)
//}
