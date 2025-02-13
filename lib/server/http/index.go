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

type Schedule struct {
	OnTime  schedule.SecondsInDay
	OffTime schedule.SecondsInDay
	Row     int
}

type GridDevice struct {
	RowNumber      int
	DisplayClasses string
	Schedules      []Schedule
	FriendlyName   string
}

func (gd GridDevice) InlineStyles() template.HTMLAttr {
	return template.HTMLAttr(fmt.Sprintf("style=\"grid-row-start: %v; grid-column-start:1 ; grid-column-end: 1\"", gd.RowNumber+1))
}

func (s Schedule) Title() string {
	return fmt.Sprintf("%s - %s", s.OnTime.HumanReadble(), s.OffTime.HumanReadble())
}

func (s Schedule) InlineStyles() template.HTMLAttr {
	onTime := s.OnTime   // time in seconds
	offTime := s.OffTime // time in seconds
	// to column means -> 86400 second, divided by grid size 48
	columnSize := 86400 / 48
	// number of seconds as a half-hour
	startColumn := 1 + (int(onTime) / columnSize)
	endColumn := 1 + (int(offTime) / columnSize)
	return template.HTMLAttr(fmt.Sprintf("style=\"grid-row-start: %v; grid-column-start:tick %v ; grid-column-end: tick %v\"", s.Row+1, startColumn, endColumn))
}

type Legend struct {
	DisplayClasses string
	Style          template.HTMLAttr
	Title          string
}

type ViewGrid struct {
	Devices     []GridDevice
	Legends     []Legend
	GridClasses string
}

func grid(list []*schedule.Device) ViewGrid {
	gridDevices := make([]GridDevice, len(list))
	for i, device := range list {
		schedules := displaySchedules(device.Schedules, i+1)
		gridDevices[i] = GridDevice{
			RowNumber:    i + 1,
			Schedules:    schedules,
			FriendlyName: device.FriendlyName,
		}
	}

	var legends = make([]Legend, 48)

	for i := 0; i < 48; i++ {
		title := ""
		if i%2 == 0 {
			hour := i / 2 % 12
			if hour == 0 {
				hour = 12
			}
			title = fmt.Sprintf("%d", hour)
		}
		legends[i] = Legend{
			DisplayClasses: "legend",
			Style:          template.HTMLAttr(fmt.Sprintf("style=\"grid-column-start:tick %d\"", i+1)),
			Title:          title,
		}
	}
	return ViewGrid{
		Devices:     gridDevices,
		Legends:     legends,
		GridClasses: "",
	}
}

func displaySchedules(schedules []*schedule.DeviceSchedule, row int) []Schedule {
	var result []Schedule
	for _, deviceSchedule := range schedules {
		result = append(result, Schedule{
			OnTime:  deviceSchedule.OnTime,
			OffTime: deviceSchedule.OffTime,
			Row:     row,
		})
	}
	return result
}
