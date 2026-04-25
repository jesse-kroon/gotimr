package timr

import (
	"fmt"
	"time"
)

type Timr struct {
	Title    string
	Progress chan string
}

func NewTimr(title string, duration, interval time.Duration) *Timr {
	t := &Timr{
		Title:    title,
		Progress: make(chan string, 1),
	}

	go func() {
		start := time.Now()
		endTime := start.Add(duration)

		timer := time.NewTimer(duration)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-timer.C:
				t.Progress <- "Done"
				close(t.Progress)
				return

			case <-ticker.C:
				remaining := max(time.Until(endTime), 0)

				hours := int(remaining.Hours())
				minutes := int(remaining.Minutes()) % 60
				seconds := int(remaining.Seconds()) % 60

				var formatted string
				if hours > 0 {
					formatted = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
				} else {
					formatted = fmt.Sprintf("%02d:%02d", minutes, seconds)
				}

				select {
				case t.Progress <- formatted:
				default:
				}
			}
		}
	}()

	return t
}
