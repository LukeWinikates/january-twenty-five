package schedule

import (
	"fmt"
	"time"
)

type DeviceSchedule struct {
	OnTime     SecondsInDay
	OffTime    SecondsInDay
	Brightness uint8
}

type Device struct {
	FriendlyName string
	Schedules    []*DeviceSchedule
}

const Second = 1
const Minute = 60 * Second
const Hour = 60 * Minute
const PM = 12 * Hour

type SecondsInDay int

func ToFriendlyTime(seconds SecondsInDay) string {
	duration, err := time.ParseDuration(fmt.Sprintf("%vs", seconds))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", duration)
}
