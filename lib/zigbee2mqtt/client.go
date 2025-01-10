package zigbee2mqtt

import (
	"LukeWinikates/january-twenty-five/lib/payloads"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func init() {
	//mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	//mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	//mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
}

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

type Client interface {
	SubscribeDeviceCatalog(func(devices []payloads.MessagePayload))
}

type RealClient struct {
	mqttClient mqtt.Client
}

func (c *RealClient) SubscribeDeviceCatalog(f func(devices []payloads.MessagePayload)) {
	topic := "zigbee2mqtt/bridge/devices"
	token := c.mqttClient.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message from topic: %s\n", message.Topic())
		deviceList, err := payloads.Parse(message.Payload())
		if err != nil {
			log.Default().Printf("err: %s", err.Error())
			return
		}
		f(deviceList)
	})
	token.Wait()
}

func NewClient(mqttHost string) Client {
	options := mqtt.NewClientOptions()
	options.AddBroker(mqttHost)
	options.SetClientID("hogepiyo")
	client := mqtt.NewClient(options)

	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &RealClient{
		mqttClient: client,
	}
}
