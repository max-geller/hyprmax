package ui

import (
	"fmt"
	"github.com/max-geller/hyprmax/config"
)

func NewGeneralSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	return settingsModel{
		config:  cfg,
		section: "General",
		settings: []setting{
			{"Border Size", cfg.General.BorderSize, true},
			{"Gaps In", cfg.General.GapIn, true},
			{"Gaps Out", cfg.General.GapOut, true},
			{"Sensitivity", cfg.General.SensitivityMultiplier, true},
			{"Cursor Zoom Factor", cfg.General.CursorZoomFactor, true},
			{"Layout", cfg.General.Layout, true},
			{"Allow Tearing", cfg.General.AllowTearing, true},
			// Add other general settings...
		},
	}
}

func NewDecorationSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	return settingsModel{
		config:  cfg,
		section: "Decoration",
		settings: []setting{
			{"Rounding", cfg.Decoration.Rounding, true},
			{"Blur Enabled", cfg.Decoration.BlurEnabled, true},
			{"Blur Size", cfg.Decoration.BlurSize, true},
			{"Blur Passes", cfg.Decoration.BlurPasses, true},
			{"Active Opacity", cfg.Decoration.Opacity, true},
			{"Inactive Opacity", cfg.Decoration.InactiveOpacity, true},
			{"Drop Shadow", cfg.Decoration.DropShadow, true},
			{"Shadow Range", cfg.Decoration.ShadowRange, true},
			{"Shadow Color", cfg.Decoration.ShadowColor, true},
		},
	}
}

func NewAnimationsSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	return settingsModel{
		config:  cfg,
		section: "Animations",
		settings: []setting{
			{"Enabled", cfg.Animations.Enabled, true},
			// Add bezier curves as sub-settings
			{"Add Bezier Curve", "New", true},
			// Add animations as sub-settings
			{"Add Animation", "New", true},
		},
	}
}

func NewInputSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	return settingsModel{
		config:  cfg,
		section: "Input",
		settings: []setting{
			{"Keyboard Model", cfg.Input.KBModel, true},
			{"Keyboard Layout", cfg.Input.KBLayout, true},
			{"Keyboard Variant", cfg.Input.KBVariant, true},
			{"Keyboard Options", cfg.Input.KBOptions, true},
			{"NumLock by Default", cfg.Input.NumLockByDefault, true},
			{"Scroll Method", cfg.Input.ScrollMethod, true},
			{"Scroll Button", cfg.Input.ScrollButton, true},
			{"Scroll Factor", cfg.Input.ScrollFactor, true},
			{"Follow Mouse", cfg.Input.FollowMouse, true},
			{"Mouse Refocus", cfg.Input.MouseRefocus, true},
		},
	}
}

func NewWindowRulesSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	// Convert existing rules to settings
	settings := []setting{
		{"Add New Rule", "New", true},
	}
	
	for i, rule := range cfg.WindowRules {
		name := fmt.Sprintf("Rule %d", i+1)
		value := fmt.Sprintf("%s,%s,%s", rule.Rule, rule.Value, rule.Target)
		settings = append(settings, setting{name, value, true})
	}
	
	return settingsModel{
		config:  cfg,
		section: "Window Rules",
		settings: settings,
	}
}

func NewKeybindingsSettingsModel(cfg *config.HyprlandConfig) SettingsModel {
	// Convert existing binds to settings
	settings := []setting{
		{"Add New Binding", "New", true},
	}
	
	for i, bind := range cfg.Binds {
		name := fmt.Sprintf("Bind %d", i+1)
		if bind.Description != "" {
			name = bind.Description
		}
		value := fmt.Sprintf("%s + %s → %s %s", bind.Mods, bind.Key, bind.Dispatcher, bind.Params)
		settings = append(settings, setting{name, value, true})
	}
	
	return settingsModel{
		config:  cfg,
		section: "Keybindings",
		settings: settings,
	}
}

// Add other section models... 