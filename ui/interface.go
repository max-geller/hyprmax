package ui

import tea "github.com/charmbracelet/bubbletea"

// SettingsModel represents the interface that all settings views must implement
type SettingsModel interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}
