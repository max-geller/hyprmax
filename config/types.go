package config

// HyprlandConfig represents the main configuration structure
type HyprlandConfig struct {
	General      GeneralSection
	Decoration   DecorationSection
	Animations   AnimationsSection
	Input        InputSection
	Touchpad     TouchpadSection
	Gestures     GesturesSection
	Misc         MiscSection
	WindowRules  []WindowRule
	LayerRules   []LayerRule
	Binds        []Bind
	Monitors     []Monitor
	Workspaces   []Workspace
	Debug        DebugSection
	XWayland     XWaylandSection
	OpenGL       OpenGLSection
	Cursor       CursorSection
}

type GeneralSection struct {
	BorderSize         int     `hypr:"border_size"`
	GapIn             int     `hypr:"gaps_in"`
	GapOut            int     `hypr:"gaps_out"`
	Cursor            string  `hypr:"cursor_inactive_timeout"`
	Layout            string  `hypr:"layout"`
	NoFocusFollowMouse bool    `hypr:"no_focus_fallback"`
	SensitivityMultiplier float64 `hypr:"sensitivity"`
	ApplyTweaks          bool    `hypr:"apply_sens_to_raw"`
	CursorZoomFactor     float64 `hypr:"cursor_zoom_factor"`
	ResizeOnBorder       bool    `hypr:"resize_on_border"`
	ExtendBorderGrabArea int     `hypr:"extend_border_grab_area"`
	HoverIconOnBorder    bool    `hypr:"hover_icon_on_border"`
	AllowTearing         bool    `hypr:"allow_tearing"`
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
	Beziers []BezierCurve
	Animations []Animation
}

type InputSection struct {
	KBModel              string  `hypr:"kb_model"`
	KBLayout            string  `hypr:"kb_layout"`
	KBVariant           string  `hypr:"kb_variant"`
	KBOptions           string  `hypr:"kb_options"`
	NumLockByDefault    bool    `hypr:"numlock_by_default"`
	ScrollMethod        string  `hypr:"scroll_method"`
	ScrollButton        int     `hypr:"scroll_button"`
	ScrollButtonLock    bool    `hypr:"scroll_button_lock"`
	ScrollFactor        float64 `hypr:"scroll_factor"`
	FollowMouse         int     `hypr:"follow_mouse"`
	MouseRefocus         bool    `hypr:"mouse_refocus"`
	// Add other input settings
}

type TouchpadSection struct {
	DisableWhileTyping bool    `hypr:"disable_while_typing"`
	NaturalScroll      bool    `hypr:"natural_scroll"`
	ScrollFactor       float64 `hypr:"scroll_factor"`
	TapToClick         bool    `hypr:"tap-to-click"`
	DragLock           bool    `hypr:"drag_lock"`
}

type GesturesSection struct {
	// Add gesture settings
}

type MiscSection struct {
	DisableHyprlandLogo    bool    `hypr:"disable_hyprland_logo"`
	DisableAutoreload      bool    `hypr:"disable_autoreload"`
	DisableStartupDrop     bool    `hypr:"disable_startup_drop"`
	VFRAlgorithm          string  `hypr:"vfr"`
	VRFMode               int     `hypr:"vrr"`
	MouseMoveEnableDPMS   bool    `hypr:"mouse_move_enables_dpms"`
	AlwaysFollowOnDND     bool    `hypr:"always_follow_on_dnd"`
	LayersHog             bool    `hypr:"layers_hog_keyboard_focus"`
	AnimateManualResizes  bool    `hypr:"animate_manual_resizes"`
	EnableSwallow         bool    `hypr:"enable_swallow"`
	SwallowRegex         string  `hypr:"swallow_regex"`
	FocusOnActivate      bool    `hypr:"focus_on_activate"`
	// Add misc settings
}

type LayerRule struct {
	// Add layer rule fields
}

type Bind struct {
	Mods        string `hypr:"mods"`
	Key         string `hypr:"key"`
	Dispatcher  string `hypr:"dispatcher"`
	Params      string `hypr:"params"`
	Flags       string `hypr:"flags"`
	Description string `hypr:"description"`
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

type DebugSection struct {
	DisableLogsInDisks bool   `hypr:"disable_logs_in_disk"`
	LogLevel           string `hypr:"log_level"`
	// Add debug settings
}

type XWaylandSection struct {
	UseNearest bool    `hypr:"use_nearest_neighbor"`
	ForceScale float64 `hypr:"force_scale"`
	// Add XWayland settings
}

type OpenGLSection struct {
	NvidiaPatches bool `hypr:"nvidia_patches"`
	// Add OpenGL settings
}

type CursorSection struct {
	HideInactive bool  `hypr:"hide_when_inactive"`
	HideTimeout  int   `hypr:"hide_timeout"`
	// Add cursor settings
}

type BezierCurve struct {
	Name      string
	Points    [4]float64
}

type Animation struct {
	Target    string
	Bezier    string
	Duration  int
	Style     string
} 