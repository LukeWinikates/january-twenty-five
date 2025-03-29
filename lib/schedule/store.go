package schedule

import "fmt"

type Store interface {
	Find(id string) (*DeviceSchedule, error)
	All() []*Device
	SaveChanges(path string, s *DeviceSchedule) error
}

type inMemoryStore struct {
	devices []*Device
}

func (store *inMemoryStore) Find(id string) (*DeviceSchedule, error) {
	var err error
	for _, device := range store.devices {
		for _, deviceSchedule := range device.Schedules {
			if deviceSchedule.ID == id {
				return deviceSchedule, err
			}
		}
	}
	err = fmt.Errorf("not found: %s", id)
	return nil, err
}

func (store *inMemoryStore) All() []*Device {
	return store.devices
}

func (store *inMemoryStore) SaveChanges(id string, s *DeviceSchedule) error {
	found := false
	for _, device := range store.devices {
		for j, deviceSchedule := range device.Schedules {
			if deviceSchedule.ID == id {
				found = true
				device.Schedules[j] = s
				// check to see if this would merge the schedule, and offer that suggestion?
				fmt.Printf("%#v", s)
			}
		}
	}
	if found {
		fmt.Printf("updated devices\n%#v %#v\n", store.devices[0], store.devices[1])
		return nil
	}
	return fmt.Errorf("did not find schedule with id: %s", id)
}

func NewStore() Store {
	return &inMemoryStore{
		devices: []*Device{
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
		},
	}
}
