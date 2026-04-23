package timr

import (
	"fmt"
	"time"
)

type Timr struct {
	Title    string
	Done     chan struct{}
	Progress chan float64
}

func NewTimr(title string, duration, interval time.Duration) *Timr {
	timr := &Timr{
		Title:    title,
		Done:     make(chan struct{}),
		Progress: make(chan float64),
	}

	go func() {
		start := time.Now()
		timer := time.NewTimer(duration)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-timer.C:
				fmt.Printf("%s is at 100%%\n", title)
				close(timr.Done)
				close(timr.Progress)
				return
			case <-ticker.C:
				elapsed := time.Since(start)
				percent := float64(elapsed) / float64(duration) * 100
				timr.Progress <- percent
			}
		}

	}()

	return timr
}
