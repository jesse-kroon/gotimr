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

type TimeUnit int

func main() {
	var duration time.Duration
	var interval time.Duration
	var title string

	cmd := &cli.Command{
		Name:  "gotimr",
		Usage: "a cli pomodoro tool written in go",
		Flags: []cli.Flag{
			&cli.DurationFlag{
				Name:        "duration",
				Value:       time.Minute,
				Usage:       "duration of the timer (e.g. '45s' '1m20s' '2h30m')",
				Aliases:     []string{"d"},
				Destination: &duration,
			},
			&cli.DurationFlag{
				Name:        "interval",
				Value:       time.Second,
				Usage:       "tick interval (how timer is updated)",
				Aliases:     []string{"i"},
				Destination: &interval,
			},
			&cli.StringFlag{
				Name:        "Title",
				Value:       "Timer",
				Usage:       "title of your timer",
				Aliases:     []string{"t"},
				Destination: &title,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			timer := timr.NewTimr(title, duration, interval)
			fmt.Printf("%s\n", timer.Title)
			for p := range timer.Progress {
				fmt.Printf("\r\033[K%s", p)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
