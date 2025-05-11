package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// EditorMode represents different editing modes
type EditorMode int

const (
	ModeNormal EditorMode = iota
	ModeNewRule
	ModeNewBind
	ModeEditRule
	ModeEditBind
)

type editorModel struct {
	mode       EditorMode
	fields     []editorField
	cursor     int
	parentView SettingsModel
}

type editorField struct {
	name     string
	value    string
	required bool
}

// NewRuleEditor creates a new window rule editor
func NewRuleEditor(parent SettingsModel) editorModel {
	return editorModel{
		mode:       ModeNewRule,
		parentView: parent,
		fields: []editorField{
			{"Rule Type", "", true},
			{"Value", "", true},
			{"Target Window", "", true},
		},
	}
}

// NewBindEditor creates a new keybinding editor
func NewBindEditor(parent SettingsModel) editorModel {
	return editorModel{
		mode:       ModeNewBind,
		parentView: parent,
		fields: []editorField{
			{"Modifiers", "", true},
			{"Key", "", true},
			{"Action", "", true},
			{"Parameters", "", false},
			{"Description", "", false},
		},
	}
}

func (e editorModel) View() string {
	var s string

	switch e.mode {
	case ModeNewRule:
		s += titleStyle.Render("New Window Rule") + "\n\n"
	case ModeNewBind:
		s += titleStyle.Render("New Keybinding") + "\n\n"
	}

	for i, field := range e.fields {
		cursor := " "
		if i == e.cursor {
			cursor = "► "
		}

		required := ""
		if field.required {
			required = "*"
		}

		s += fmt.Sprintf("%s%s%s: %s\n",
			cursor,
			settingStyle.Render(field.name),
			required,
			valueStyle.Render(field.value))
	}

	s += "\n" + itemStyle.Render("(↑/↓) navigate • (enter) edit • (esc) cancel • (ctrl+s) save")
	return s
}

// Add these methods to implement tea.Model
func (e editorModel) Init() tea.Cmd {
	return nil
}

func (e editorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return e.parentView, nil
		case "up", "k":
			if e.cursor > 0 {
				e.cursor--
			}
		case "down", "j":
			if e.cursor < len(e.fields)-1 {
				e.cursor++
			}
		}
	}
	return e, nil
}
