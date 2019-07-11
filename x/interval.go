// Package repeat/x will be replaced soon.
package x

import "time"

type Func func() error

func Interval(freq time.Duration, f Func) error {
	// manually execute first iteration immediately
	if err := f(); err != nil {
		return err
	}

	for range time.Tick(freq) {
		if err := f(); err != nil {
			return err
		}
	}
	// never get here
	return nil
}
