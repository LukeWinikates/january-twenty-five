package schedule

type DeviceSchedule struct {
	OnTime     int
	OffTime    int
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
