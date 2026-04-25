package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {
	durationPtr := flag.Int("duration", 120, "duration of the timer in minutes")
	flag.Parse()
	duration := time.Duration(*durationPtr) * time.Minute

	timer := timr.NewTimr("test", duration, 1*time.Second)

	for p := range timer.Progress {
		fmt.Println(p)
	}
}
