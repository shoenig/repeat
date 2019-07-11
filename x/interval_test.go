package x

import (
	"fmt"
	"testing"
	"time"
)

func Test_Interval(t *testing.T) {
	counter := 0

	err := Interval(10*time.Millisecond, func() error {
		if counter < 3 {
			counter++
			return nil
		}
		return fmt.Errorf("counter is >= 3")
	})

	if counter != 3 {
		t.Fatal("counter should have been 3")
	}

	if err == nil {
		t.Fatal("expected an error")
	}
}
