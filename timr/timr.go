package timr

import (
	"fmt"
	"time"
)

type Timr struct {
	title string
	Done  chan struct{}
}

func NewTimr(title string, duration time.Duration) *Timr {
	t := &Timr{
		title: title,
		Done:  make(chan struct{}),
	}

	go func() {
		timer := time.NewTimer(duration)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-timer.C:
				fmt.Printf("%s expired\n", title)
				close(t.Done)
				return
			case <-ticker.C:
				fmt.Printf("%s is still ticking\n", title)
			}
		}

	}()

	return t
}
