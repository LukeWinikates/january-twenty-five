package http

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var homepageTemplate *template.Template

func init() {
	homepageTemplate = template.Must(template.ParseFiles("lib/server/http/index.gohtml"))
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
	fs := os.DirFS("./public")
	fmt.Println(os.ReadDir("./public"))

	//mux.Handle("/", http.FileServer(http.FS(fs)))
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(fs))))
	mux.HandleFunc("/", indexPage())
	server := &http.Server{Addr: addr, Handler: mux}
	s.server = server
	return server.ListenAndServe()
}

// TODO: api endpoints
