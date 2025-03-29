package schedule

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	time := 10*Hour + 30*Minute + PM

	if time.HumanReadable() != "10:30 pm" {
		fmt.Println(time.HumanReadable())
		t.Fail()
	}

	time = 12*Hour + PM

	if time.HumanReadable() != "12:00 pm" {
		fmt.Println(time.HumanReadable())
		t.Fail()
	}
}
