package main

import (
	"time"

	"github.com/jesse-kroon/gotimr/timr"
)

func main() {
	timer1 := timr.NewTimr("timer1", 5*time.Second)
	timer2 := timr.NewTimr("timer2", 10*time.Second)

	<-timer1.Done
	<-timer2.Done
}
