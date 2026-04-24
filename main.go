package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices  []string
	cursor   int
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
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", "space":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString("Welcome to Timr!\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "->"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "X"
		}

		fmt.Fprintf(&s, "%s [%s] %s\n", cursor, checked, choice)
	}

	s.WriteString("\nPress q to quit.\n")
	return tea.NewView(s.String())
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
