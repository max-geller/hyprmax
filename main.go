package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/max-geller/hyprmax/config"
	"github.com/max-geller/hyprmax/ui"
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
	config   *config.HyprlandConfig
	choices  []string
	cursor   int
	selected map[int]struct{}
	err      error
	page     page
	settings ui.SettingsModel
	saveChan chan<- config.HyprlandConfig
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

func initialModel(saveChan chan<- config.HyprlandConfig) model {
	// Use test mode during development
	cfg, err := config.LoadConfig("", true)
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
		saveChan: saveChan,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.page != pageMain {
			// Handle settings pages
			var newSettings tea.Model
			newSettings, cmd := m.settings.Update(msg)
			if settingsModel, ok := newSettings.(ui.SettingsModel); ok {
				m.settings = settingsModel
			}
			return m, cmd
		}

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
			switch m.cursor {
			case 0: // General Settings
				m.page = pageGeneral
				m.settings = ui.NewGeneralSettingsModel(m.config)
			case 1: // Decoration
				m.page = pageDecoration
				m.settings = ui.NewDecorationSettingsModel(m.config)
			case 2: // Animations
				m.page = pageAnimations
				m.settings = ui.NewAnimationsSettingsModel(m.config)
			case 3: // Input
				m.page = pageInput
				m.settings = ui.NewInputSettingsModel(m.config)
			case 4: // Window Rules
				m.page = pageWindowRules
				m.settings = ui.NewWindowRulesSettingsModel(m.config)
			case len(m.choices) - 1:
				return m, tea.Quit
			default:
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
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
	// Create a channel for config saves
	saveChan := make(chan config.HyprlandConfig)

	// Start a goroutine to handle config saves
	go func() {
		for cfg := range saveChan {
			if err := config.WriteConfig(&cfg, ""); err != nil {
				// TODO: Handle error
				fmt.Printf("Error saving config: %v\n", err)
			}
		}
	}()

	p := tea.NewProgram(initialModel(saveChan))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
