package main

import (
	"LukeWinikates/january-twenty-five/lib/payloads"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func sub(client mqtt.Client) {
	topic := "zigbee2mqtt/bridge/#"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
	})
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func main() {
	//mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	//mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	//mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	options := mqtt.NewClientOptions()
	options.AddBroker(os.Getenv("MQTT_HOST"))
	options.SetClientID("hogepiyo")
	client := mqtt.NewClient(options)
	//options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	subscribeDevices(client)
	//subscribeLogging(client)
	//sub(client)

	for {
		time.Sleep(20 * time.Second)
	}
	//client.Disconnect(2500)
}

func subscribeDevices(client mqtt.Client) {
	topic := "zigbee2mqtt/bridge/devices"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message from topic: %s\n", message.Topic())
		//err := os.WriteFile("devices.json", message.Payload(), os.ModePerm)
		//if err != nil {
		//	log.Default().Printf("err: %s", err.Error())
		//}
		deviceList, err := payloads.Parse(message.Payload())
		if err != nil {
			log.Default().Printf("err: %s", err.Error())
			return
		}
		for _, device := range deviceList {
			fmt.Println(device.FriendlyName)
		}
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
