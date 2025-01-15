package server

import (
	"LukeWinikates/january-twenty-five/lib/server/http"
	"LukeWinikates/january-twenty-five/lib/zigbee2mqtt"
)

type Server interface {
	Start()
	Stop() error
}

type realServer struct {
	client     zigbee2mqtt.Client
	httpServer http.Server
	dataDir    string
}

func (r *realServer) Start() {
	// initialize database
	// start device listener
}

func (r *realServer) Stop() error {
	return nil
}

func New(client zigbee2mqtt.Client, dataDir string) (Server, error) {
	return &realServer{
		client:  client,
		dataDir: dataDir,
	}, nil
}
