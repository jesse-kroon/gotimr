package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

type TimeUnit int

func main() {
	// cmd := &cli.Command{
	// 	Commands: []*cli.Command{
	// 		{},
	// 	},
	// }
	//
	// if err := cmd.Run(context.Background(), os.Args); err != nil {
	// 	log.Fatal(err)
	// }
	titlePtr := flag.String("title", "Timer", "timer title")
	durationPtr := flag.Duration("duration", time.Minute, "duration of the timer. Accepted inputs are e.g. \"45s\"  \"1m17s\"")
	intervalPtr := flag.Int("interval", 1, "tick interval in seconds")
	flag.Parse()

	duration := *durationPtr
	interval := time.Duration(*intervalPtr) * time.Second

	timer := timr.NewTimr(*titlePtr, duration, interval)

	fmt.Printf("%s\n", timer.Title)
	for p := range timer.Progress {
		fmt.Printf("\r\033[K%s", p)
	}
}
