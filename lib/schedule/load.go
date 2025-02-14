package schedule

func LoadDevices() []*Device {
	return []*Device{
		{
			FriendlyName: "Bedroom",
			ID:           "E0D5118D-1554-4394-93A8-EFC6C7276D0A",
			Schedules: []*DeviceSchedule{
				{
					ID:         "3265D1FD-4FE5-4662-8AFE-C966089BCCB9",
					OnTime:     8 * Hour,
					OffTime:    5*Hour + PM,
					Color:      "#33b73c",
					Brightness: 100,
				},
				{
					ID:         "3265D1FD-4FE5-4662-8AFE-C966089BCCB0",
					OnTime:     8*Hour + PM,
					OffTime:    10*Hour + 30*Minute + PM,
					Color:      "#f60080",
					Brightness: 75,
				},
			},
		}, {
			FriendlyName: "Kitchen",
			ID:           "31CD5DBD-E5F9-43FE-A6D3-FB7D5E07E57F",
			Schedules: []*DeviceSchedule{
				{
					ID:         "271FA53F-7CB8-4624-A164-5203BCCBB4FA",
					OnTime:     12 * Hour,
					OffTime:    9*Hour + PM,
					Color:      "#f6b73c",
					Brightness: 60,
				},
			},
		},
	}
}
