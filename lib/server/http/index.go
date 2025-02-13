package http

import (
	"LukeWinikates/january-twenty-five/lib/schedule"
	"LukeWinikates/january-twenty-five/lib/server/http/index"
	"fmt"
	"net/http"
)

func indexPage() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		deviceList := schedule.LoadDevices()
		viewModel := index.Grid(deviceList)
		err := homepageTemplate.Execute(writer, viewModel)
		if err != nil {
			writer.WriteHeader(500)
			fmt.Println(err.Error())
			return
		}
	}
}
