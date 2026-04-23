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
				fmt.Println("Timer expired")
				close(t.Done)
				return
			case <-ticker.C:
				fmt.Println(fmt.Sprintf("%s is still ticking", title))
			}
		}

	}()

	return t
}
