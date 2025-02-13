package schedule

import (
	"fmt"
	"time"
)

type DeviceSchedule struct {
	OnTime     SecondsInDay
	OffTime    SecondsInDay
	Brightness uint8
	ID         string
}

type Device struct {
	FriendlyName string
	Schedules    []*DeviceSchedule
	ID           string
}

const Second = 1
const Minute = 60 * Second
const Hour = 60 * Minute
const PM = 12 * Hour

type SecondsInDay int

func (s SecondsInDay) HumanReadble() string {
	ampm := "am"
	if s > PM {
		s -= PM
		ampm = "pm"
	}
	hour := s / Hour
	minute := s % Hour
	return fmt.Sprintf("%02d:%02d %s", hour, minute, ampm)
}

func ToFriendlyTime(seconds SecondsInDay) string {
	duration, err := time.ParseDuration(fmt.Sprintf("%vs", seconds))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", duration)
}
