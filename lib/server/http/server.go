package http

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

var homepageTemplate *template.Template

func init() {
	homepageTemplate = template.Must(template.ParseFiles("lib/server/http/index.gohtml"))
}

type Device struct {
	FriendlyName string
}

type Server interface {
	Serve(addr string) error
	Stop() error
}

type realServer struct {
	server *http.Server
}

func NewServer() Server {
	return &realServer{}
}

func (s *realServer) Stop() error {
	return s.server.Shutdown(context.TODO())
}

func (s *realServer) Serve(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		deviceList := []*Device{&Device{FriendlyName: "One"}, &Device{"Two"}, &Device{"Three"}}
		err := homepageTemplate.Execute(writer, deviceList)
		if err != nil {
			writer.WriteHeader(500)
			fmt.Println(err.Error())
			return
		}
	})
	server := &http.Server{Addr: addr, Handler: mux}
	s.server = server
	return server.ListenAndServe()
}

// TODO: api endpoints
