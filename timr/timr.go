package timr

import (
	"context"
	"fmt"
	"math"
	"time"
)

type Timr struct {
	Title string
	C     <-chan Tick
}

type Tick struct {
	Remaining time.Duration
	Elapsed   time.Duration
	Done      bool
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

func New(ctx context.Context, title string, duration, interval time.Duration) *Timr {
	ch := make(chan Tick)
	timr := &Timr{
		Title: title,
		C:     ch,
	}

	go func() {
		defer close(ch)
		start := time.Now()
		endTime := start.Add(duration)

		timer := time.NewTimer(duration)
		defer timer.Stop()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		ch <- Tick{
			Remaining: duration,
			Elapsed:   0,
			Done:      false,
		}

		for {
			select {
			case <-ctx.Done():
				return

			case <-timer.C:
				ch <- Tick{
					Remaining: 0,
					Elapsed:   time.Since(start),
					Done:      true,
				}
				return

			case <-ticker.C:
				now := time.Now()
				remaining := max(endTime.Sub(now), 0)

				ch <- Tick{
					Remaining: remaining,
					Elapsed:   now.Sub(start),
					Done:      false,
				}
			}
		}
	}()

	return timr
}
