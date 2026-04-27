package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jesse-kroon/gotimr/timr"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gotimr",
		Usage: "a cli pomodoro tool written in go",
		Flags: []cli.Flag{
			&cli.DurationFlag{
				Name:    "duration",
				Aliases: []string{"d"},
				Value:   time.Minute,
				Usage:   "duration (e.g. 45s, 1m20s, 2h",
			},
			&cli.DurationFlag{
				Name:    "interval",
				Aliases: []string{"a"},
				Value:   time.Second,
				Usage:   "update interval",
			},
			&cli.StringFlag{
				Name:    "title",
				Aliases: []string{"t"},
				Value:   "Timer",
				Usage:   "timer title",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			duration := c.Duration("duration")
			interval := c.Duration("interval")
			title := c.String("title")

			if duration <= 0 {
				return fmt.Errorf("duration must be > 0")
			}

			if interval <= 0 {
				return fmt.Errorf("interval must be > 0")
			}

			if interval > duration {
				return fmt.Errorf("interval cannot exceed duration")
			}

			ctx, cancel := context.WithCancel()
			defer cancel()

			timr := timr.New(ctx, title, duration, interval)

			fmt.Println(timr.Title)

			for tick := range timr.C {
				fmt.Printf("\r\033[K%s", formatRemaining(tick.Remaining))

				if tick.Done {
					fmt.Print("\nDone\n")
				}
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	// titlePtr := flag.String("title", "Timer", "timer title")
	// intervalPtr := flag.Int("interval", 1, "tick interval in seconds")
	// durationPtr := flag.Int("duration", 1, "duration of the timer in minutes")
	// flag.Parse()
	//
	// duration := time.Duration(*durationPtr) * time.Minute
	// interval := time.Duration(*intervalPtr) * time.Second
	//
	// timer := timr.NewTimr(*titlePtr, duration, interval)
	//
	// fmt.Printf("%s\n", timer.Title)
	// for p := range timer.Progress {
	// 	fmt.Printf("\r\033[K%s", p)
	// }
}
