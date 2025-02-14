package index

import (
	"LukeWinikates/january-twenty-five/lib/schedule"
	"fmt"
	"html/template"
)

type Schedule struct {
	OnTime     schedule.SecondsInDay
	OffTime    schedule.SecondsInDay
	Row        int
	ID         string
	Brightness uint8
	Color      string
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
	return fmt.Sprintf("%s - %s", s.OnTime.HumanReadable(), s.OffTime.HumanReadable())
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

func Grid(list []*schedule.Device) ViewGrid {
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
			ID:         deviceSchedule.ID,
			OnTime:     deviceSchedule.OnTime,
			OffTime:    deviceSchedule.OffTime,
			Brightness: deviceSchedule.Brightness,
			Row:        row,
			Color:      deviceSchedule.Color,
		})
	}
	return result
}
