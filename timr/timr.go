package timr

import (
	"fmt"
	"math"
	"time"
)

type Timr struct {
	Title    string
	Progress chan string
}

func formatRemaining(d time.Duration) string {
	totalSeconds := int(math.Ceil(d.Seconds()))

	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
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

		t.Progress <- formatRemaining(duration)
		for {
			select {
			case <-timer.C:
				t.Progress <- "Done"
				close(t.Progress)
				return

			case <-ticker.C:
				remaining := max(time.Until(endTime), 0)

				formatted := formatRemaining(remaining)

				select {
				case t.Progress <- formatted:
				default:
				}
			}
		}
	}()

	return t
}
