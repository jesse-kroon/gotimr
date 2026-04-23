package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/jesse-kroon/gotimr/timr"
)

type model struct {
	choices  []string
	pointer  int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Add a Timr"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView("")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh jeez, oh god we've screwed something up Rick.")
		os.Exit(1)
	}
	// durationPtr := flag.Int("duration", 10, "duration of the timer in seconds")
	// tickInterval := flag.Int("interval", 1, "interval of ticks used (in seconds) to print timer progress")
	// flag.Parse()
	//
	// duration := time.Duration(*durationPtr) * time.Second
	// interval := time.Duration(*tickInterval) * time.Second
	//
	// timer1 := timr.NewTimr("timer1", duration, interval)
	//
	// for p := range timer1.Progress {
	// 	fmt.Printf("%s is at %0.f%%\n", timer1.Title, p)
	// }
	// <-timer1.Done
}
