package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/max-geller/hyprmax/config"
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7dcfff")).
		PaddingLeft(2)

	itemStyle = lipgloss.NewStyle().
		PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#bb9af7")).
		SetString("► ")
)

type model struct {
	config    *config.HyprlandConfig
	choices   []string
	cursor    int
	selected  map[int]struct{}
	err       error
	page      page
}

type page int

const (
	pageMain page = iota
	pageGeneral
	pageDecoration
	pageAnimations
	pageInput
	pageWindowRules
)

func initialModel() model {
	cfg, err := config.LoadConfig("")
	return model{
		config: cfg,
		err:    err,
		choices: []string{
			"General Settings",
			"Decoration",
			"Animations",
			"Input",
			"Window Rules",
			"Keybindings",
			"Save & Quit",
		},
		selected: make(map[int]struct{}),
		page:     pageMain,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor == len(m.choices)-1 {
				return m, tea.Quit
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "\n"
	s += titleStyle.Render("Hyprland Settings Manager") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = selectedItemStyle.String()
		}
		s += fmt.Sprintf("%s%s\n", cursor, itemStyle.Render(choice))
	}
	s += "\n"
	s += itemStyle.Render("(use arrow keys to navigate, enter to select, q to quit)") + "\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
} 