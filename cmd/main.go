package main

import (
	"fmt"
	mqtt2 "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

func main() {
	options := mqtt2.NewClientOptions()
	options.AddBroker(os.Getenv("MQTT_HOST"))
	client := mqtt2.NewClient(options)
	topic := os.Getenv("MQTT_TOPIC")
	t := client.Subscribe(topic, 0, func(client mqtt2.Client, message mqtt2.Message) {
		log.Println(message)
	})
	<-t.Done() // Can also use '<-t.Done()' in releases > 1.2.0
	if t.Error() != nil {
		fmt.Println(t.Error().Error())
		os.Exit(1)
	}
}
