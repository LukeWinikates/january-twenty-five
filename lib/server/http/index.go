package http

import (
	"LukeWinikates/january-twenty-five/lib/schedule"
	"fmt"
	"html/template"
	"net/http"
)

func indexPage() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		deviceList := []*schedule.Device{
			{
				FriendlyName: "One",
				Schedules: []*schedule.DeviceSchedule{
					{
						OnTime:  8 * schedule.Hour,
						OffTime: 5*schedule.Hour + schedule.PM,
					},
				},
			}, {
				FriendlyName: "Two",
				Schedules: []*schedule.DeviceSchedule{
					{
						OnTime:  12 * schedule.Hour,
						OffTime: 9*schedule.Hour + schedule.PM,
					},
				},
			},
		}
		viewModel := grid(deviceList)
		err := homepageTemplate.Execute(writer, viewModel)
		if err != nil {
			writer.WriteHeader(500)
			fmt.Println(err.Error())
			return
		}
	}
}

type GridDevice struct {
	RowNumber      int
	DisplayClasses string
	Device         schedule.Device
}

func (g GridDevice) InlineStyles() template.HTMLAttr {
	onTime := g.Device.Schedules[0].OnTime   // time in seconds
	offTime := g.Device.Schedules[0].OffTime // time in seconds
	// to column means -> 86400 second, divided by grid size 48
	columnSize := 86400 / 48
	// number of seconds as a half-hour
	startColumn := int(onTime) / columnSize
	endColumn := int(offTime) / columnSize
	return template.HTMLAttr(fmt.Sprintf("style=\"grid-row-start: %v; grid-column-start:%v ; grid-column-end: %v\"", g.RowNumber+1, startColumn, endColumn))
}

type ViewGrid struct {
	Devices     []GridDevice
	GridClasses string
}

func grid(list []*schedule.Device) ViewGrid {
	gridDevices := make([]GridDevice, len(list))
	for i, device := range list {
		gridDevices[i] = GridDevice{
			RowNumber:      i + 1,
			DisplayClasses: "blue",
			Device:         *device,
		}
	}

	return ViewGrid{
		Devices:     gridDevices,
		GridClasses: "",
	}
}
