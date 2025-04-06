package schedule

import "fmt"

type Store interface {
	Find(id string) (*Schedule, error)
	All() []*Schedule
	SaveChanges(id string, s *Schedule) error
}

type inMemoryStore struct {
	schedules []*Schedule
}

func (store *inMemoryStore) Find(id string) (*Schedule, error) {
	var err error
	for _, schedule := range store.schedules {
		if schedule.ID == id {
			return schedule, nil
		}
	}
	err = fmt.Errorf("not found: %s", id)
	return nil, err
}

func (store *inMemoryStore) All() []*Schedule {
	return store.schedules
}

func (store *inMemoryStore) SaveChanges(id string, s *Schedule) error {
	found := false
	for i, schedule := range store.schedules {
		if schedule.ID == id {
			found = true
			store.schedules[i] = s
		}
	}
	if found {
		fmt.Printf("updated schedules\n%#v\n", store.schedules)
		return nil
	}
	return fmt.Errorf("did not find schedule with id: %s", id)
}

func NewStore() Store {
	nightstand := &Device{
		FriendlyName: "Bedroom Nightstand",
		ID:           "3265D1FD-4FE5-4662-8AFE-C966089BCCB9",
	}
	playroomLight := &Device{
		FriendlyName: "Playroom Light",
		ID:           "3265D1FD-4FE5-4662-8AFE-C966089BCCB9",
	}
	deskLamp1 := &Device{
		FriendlyName: "Word Desk Lamp",
		ID:           "3265D1FD-4FE5-4662-8AFE-C966089BCCB9",
	}
	deskLamp2 := &Device{
		FriendlyName: "Family Desk Lamp",
		ID:           "3265D1FD-4FE5-4662-8AFE-C966089BCCB9",
	}
	return &inMemoryStore{
		schedules: []*Schedule{
			{
				FriendlyName: "Bedroom",
				ID:           "E0D5118D-1554-4394-93A8-EFC6C7276D0A",
				OnTime:       8 * Hour,
				OffTime:      5*Hour + PM,
				DeviceSettings: []*DeviceSetting{
					{
						Device:     nightstand,
						Brightness: 100,
						Color:      "33b73c",
					},
				},
			}, {
				FriendlyName: "Kitchen",
				ID:           "31CD5DBD-E5F9-43FE-A6D3-FB7D5E07E57F",
				OnTime:       12 * Hour,
				OffTime:      9*Hour + PM,
				DeviceSettings: []*DeviceSetting{
					{
						Device:     playroomLight,
						Brightness: 100,
						Color:      "33b73c",
					},
				},
			},
			{
				ID:           "3265D1FD-4FE5-4662-8AFE-C966089BCCB0",
				FriendlyName: "Study",
				OnTime:       8*Hour + PM,
				OffTime:      10*Hour + 30*Minute + PM,

				DeviceSettings: []*DeviceSetting{
					{
						Device:     deskLamp1,
						Brightness: 75,
						Color:      "33b73c",
					},
					{
						Device:     deskLamp2,
						Brightness: 75,
						Color:      "f60080",
					},
				},
			},
		},
	}
}
