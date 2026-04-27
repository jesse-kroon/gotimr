package timr

import (
	"context"
	"time"
)

type Tick struct {
	Remaining time.Duration
	Elapsed   time.Duration
	Done      bool
}

type Timr struct {
	Title string
	C     <-chan Tick
}

func New(ctx context.Context, title string, duration, interval time.Duration) *Timr {
	ch := make(chan Tick)

	t := &Timr{
		Title: title,
		C:     ch,
	}

	go func() {
		defer close(ch)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		remaining := duration
		elapsed := time.Duration(0)

		// initial tick (immediate feedback)
		ch <- Tick{
			Remaining: remaining,
			Elapsed:   elapsed,
			Done:      false,
		}

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				elapsed += interval
				remaining -= interval

				if remaining <= 0 {
					ch <- Tick{
						Remaining: 0,
						Elapsed:   elapsed,
						Done:      true,
					}
					return
				}

				ch <- Tick{
					Remaining: remaining,
					Elapsed:   elapsed,
					Done:      false,
				}
			}
		}
	}()

	return t
}
