package main

import (
	"LukeWinikates/january-twenty-five/lib/server"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := createServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("received signal: %s\n", sig.String())
		fmt.Println(s.Stop())
	}()

	fmt.Println("starting server")
	fmt.Println(s.Start())

}

func createServer() (server.Server, error) {
	//client := zigbee2mqtt.NewClient(os.Getenv("MQTT_HOST"))
	client := zigbee2mqtt.NoOpClient()
	dataPath := os.Getenv("DATA_PATH")
	s, err := server.New(client, dataPath)
	return s, err
}
