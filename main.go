package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {
	titlePtr := flag.String("title", "Timer", "timer title")
	intervalPtr := flag.Int("interval", 1, "tick interval in seconds")
	durationPtr := flag.Int("duration", 1, "duration of the timer in minutes")
	flag.Parse()

	duration := time.Duration(*durationPtr) * time.Minute
	interval := time.Duration(*intervalPtr) * time.Second

	timer := timr.NewTimr(*titlePtr, duration, interval)

	fmt.Printf("%s\n", timer.Title)
	for p := range timer.Progress {
		fmt.Printf("\r\033[K%s", p)
	}
}
