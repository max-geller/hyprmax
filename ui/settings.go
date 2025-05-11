package ui

import (
	"fmt"
	"strconv"
	"strings"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/max-geller/hyprmax/config"
)

var (
	settingStyle = lipgloss.NewStyle().
		PaddingLeft(4)
	
	valueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7dcfff"))
	
	editStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#bb9af7")).
		Bold(true)
	
	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6666")).
		PaddingLeft(4)
	
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7dcfff")).
		PaddingLeft(4)
	
	errorListStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6666")).
		PaddingLeft(4)
)

type settingsModel struct {
	config    *config.HyprlandConfig
	section   string
	cursor    int
	editing   bool
	editValue string
	settings  []setting
	saveCmd   chan<- config.HyprlandConfig
	errorMsg  string
	errors    []string
	showHelp  bool
}

type setting struct {
	name     string
	value    interface{}
	editable bool
}

func (m settingsModel) Init() tea.Cmd {
	return nil
}

func (m settingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "?":
			m.showHelp = !m.showHelp
			return m, nil
		case "n":
			if m.section == "Window Rules" {
				return NewRuleEditor(m), nil
			} else if m.section == "Keybindings" {
				return NewBindEditor(m), nil
			}
		case "esc":
			if m.editing {
				m.editing = false
				return m, nil
			}
			return m, tea.Quit
		case "enter":
			if m.editing {
				// Validate and save the edited value
				if err := m.validateAndSave(); err != nil {
					// TODO: Show error message
					return m, nil
				}
				m.editing = false
				// Trigger config save
				if m.saveCmd != nil {
					m.saveCmd <- *m.config
				}
				return m, nil
			}
			if m.settings[m.cursor].editable {
				m.editing = true
				m.editValue = fmt.Sprintf("%v", m.settings[m.cursor].value)
			}
		case "up", "k":
			if !m.editing && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if !m.editing && m.cursor < len(m.settings)-1 {
				m.cursor++
			}
		default:
			if m.editing {
				// Handle editing
				switch msg.String() {
				case "backspace":
					if len(m.editValue) > 0 {
						m.editValue = m.editValue[:len(m.editValue)-1]
					}
				default:
					m.editValue += msg.String()
				}
			}
		}
	}
	return m, nil
}

func (m settingsModel) View() string {
	var s string
	s += titleStyle.Render(m.section + " Settings") + "\n\n"

	// Show errors if present
	if len(m.errors) > 0 {
		s += errorStyle.Render("Errors:") + "\n"
		for _, err := range m.errors {
			s += errorListStyle.Render("• " + err) + "\n"
		}
		s += "\n"
	}

	// Show help if enabled
	if m.showHelp {
		s += helpStyle.Render("Help for " + m.section) + "\n"
		switch m.section {
		case "Window Rules":
			s += helpStyle.Render("Format: rule,value,target") + "\n"
			s += helpStyle.Render("Example: workspace 1,float,title:*Firefox*") + "\n"
		case "Keybindings":
			s += helpStyle.Render("Format: MODS + KEY → ACTION PARAMS") + "\n"
			s += helpStyle.Render("Example: SUPER + Return → exec kitty") + "\n"
		}
		s += "\n"
	}

	// Show error message if present
	if m.errorMsg != "" {
		s += errorStyle.Render("Error: " + m.errorMsg) + "\n\n"
	}

	for i, setting := range m.settings {
		cursor := " "
		if m.cursor == i {
			cursor = "► "
		}

		value := fmt.Sprintf("%v", setting.value)
		if m.editing && m.cursor == i {
			value = editStyle.Render(m.editValue + "█")
		} else {
			value = valueStyle.Render(value)
		}

		s += fmt.Sprintf("%s%s: %s\n",
			cursor,
			settingStyle.Render(setting.name),
			value)
	}

	s += "\n" + itemStyle.Render("(↑/↓) navigate • (enter) edit • (esc) back")
	return s
}

func (m *settingsModel) validateAndSave() error {
	setting := m.settings[m.cursor]
	
	var err error
	switch v := setting.value.(type) {
	case int:
		err = config.ValidateInt(setting.name, m.editValue, 0, 1000)
		if err == nil {
			val, _ := strconv.Atoi(m.editValue)
			m.updateConfigValue(setting.name, val)
		}
	case float64:
		err = config.ValidateFloat(setting.name, m.editValue, 0, 10)
		if err == nil {
			val, _ := strconv.ParseFloat(m.editValue, 64)
			m.updateConfigValue(setting.name, val)
		}
	case bool:
		err = config.ValidateBool(setting.name, m.editValue)
		if err == nil {
			val := strings.ToLower(m.editValue) == "true"
			m.updateConfigValue(setting.name, val)
		}
	case string:
		// Add specific validation based on the field
		m.updateConfigValue(setting.name, m.editValue)
	}
	
	if err != nil {
		m.errorMsg = err.Error()
		return err
	}
	
	m.errorMsg = "" // Clear error on success
	return nil
}

func (m *settingsModel) updateConfigValue(name string, value interface{}) {
	// Update the config based on the section and field name
	switch m.section {
	case "General":
		switch name {
		case "Border Size":
			m.config.General.BorderSize = value.(int)
		case "Gaps In":
			m.config.General.GapIn = value.(int)
		// Add other cases...
		}
	case "Decoration":
		// Handle decoration settings
	// Add other sections...
	}
	
	// Update the displayed value
	for i, s := range m.settings {
		if s.name == name {
			m.settings[i].value = value
			break
		}
	}
} 