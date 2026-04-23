package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {

	durationPtr := flag.Int("duration", 10, "duration of the timer in seconds")
	tickInterval := flag.Int("interval", 1, "interval of ticks used (in seconds) to print timer progress")
	flag.Parse()

	duration := time.Duration(*durationPtr) * time.Second
	interval := time.Duration(*tickInterval) * time.Second

	timer1 := timr.NewTimr("timer1", duration, interval)

	for p := range timer1.Progress {
		fmt.Printf("%s is at %0.f%%\n", timer1.Title, p)
	}
	<-timer1.Done
}
