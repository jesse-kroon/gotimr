package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {
	titlePtr := flag.String("title", "Timer", "title of the timer you want to set")
	durationPtr := flag.Int("duration", 10, "duration of the timer in seconds")
	intervalPtr := flag.Int("interval", 1, "interval of ticks used (in seconds) to print timer progress")
	flag.Parse()

	duration := time.Duration(*durationPtr) * time.Second
	interval := time.Duration(*intervalPtr) * time.Second

	timer := timr.NewTimr(*titlePtr, duration, interval)

	for p := range timer.Progress {
		fmt.Printf("%s is at %0.f%%\n", timer.Title, p)
	}
}
