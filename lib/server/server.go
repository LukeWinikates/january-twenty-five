package server

import (
	"LukeWinikates/january-twenty-five/lib/server/http"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt/payloads"
	"fmt"
)

type Server interface {
	Start() error
	Stop() error
}

type realServer struct {
	ztmClient  zigbee2mqtt.Client
	httpServer http.Server
	dataDir    string
}

func (r *realServer) Start() error {
	r.ztmClient.SubscribeDeviceCatalog(func(devices []payloads.MessagePayload) {
		for _, device := range devices {
			fmt.Println(device.FriendlyName)
		}
	})
	return r.httpServer.Serve("localhost:8998")
	// start device listener
}

func (r *realServer) Stop() error {
	return nil
}

func New(client zigbee2mqtt.Client, dataDir string) (Server, error) {
	return &realServer{
		ztmClient:  client,
		dataDir:    dataDir,
		httpServer: http.NewServer(),
	}, nil
}
