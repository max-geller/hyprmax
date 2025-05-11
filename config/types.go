package config

// HyprlandConfig represents the main configuration structure
type HyprlandConfig struct {
	General      GeneralSection
	Decoration   DecorationSection
	Animations   AnimationsSection
	Input        InputSection
	Gestures     GesturesSection
	Misc         MiscSection
	WindowRules  []WindowRule
	LayerRules   []LayerRule
	Binds        []Keybind
	Monitors     []Monitor
	Workspaces   []Workspace
}

type GeneralSection struct {
	BorderSize         int    `hypr:"border_size"`
	GapIn             int    `hypr:"gaps_in"`
	GapOut            int    `hypr:"gaps_out"`
	Cursor            string `hypr:"cursor_inactive_timeout"`
	Layout            string `hypr:"layout"`
	NoFocusFollowMouse bool   `hypr:"no_focus_fallback"`
	// Add other general settings
}

type DecorationSection struct {
	Rounding           int     `hypr:"rounding"`
	BlurEnabled        bool    `hypr:"blur"`
	BlurSize           int     `hypr:"blur_size"`
	BlurPasses         int     `hypr:"blur_passes"`
	Opacity            float64 `hypr:"active_opacity"`
	InactiveOpacity    float64 `hypr:"inactive_opacity"`
	DropShadow         bool    `hypr:"drop_shadow"`
	ShadowRange        int     `hypr:"shadow_range"`
	ShadowColor        string  `hypr:"shadow_color"`
}

type WindowRule struct {
	Rule    string
	Value   string
	Target  string
}

// Add these missing types that are referenced in HyprlandConfig
type AnimationsSection struct {
	Enabled bool `hypr:"enabled"`
	// Add animation settings
}

type InputSection struct {
	// Add input settings
}

type GesturesSection struct {
	// Add gesture settings
}

type MiscSection struct {
	// Add misc settings
}

type LayerRule struct {
	// Add layer rule fields
}

type Keybind struct {
	Key     string
	Command string
}

type Monitor struct {
	Name     string
	Resolution string
	Position   string
}

type Workspace struct {
	Name     string
	Monitor  string
}

// Add other necessary types... 