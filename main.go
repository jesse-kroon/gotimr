package main

import (
	"fmt"
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {
	timer1 := timr.NewTimr("timer1", 17*time.Minute)

	for p := range timer1.Progress {
		fmt.Printf("%s is at %.2f%%\n", timer1.Title, p)
	}
	<-timer1.Done
}
