package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// DefaultConfigPath is the default location of hyprland.conf
const DefaultConfigPath = "~/.config/hypr/hyprland.conf"

// TestConfigPath is the path to the test configuration file
const TestConfigPath = "config/testdata/hyprland.conf"

// LoadConfig reads and parses the Hyprland configuration file
func LoadConfig(path string, testMode bool) (*HyprlandConfig, error) {
	if testMode {
		path = TestConfigPath
	} else if path == "" {
		path = DefaultConfigPath
		// Create backup before loading real config
		if err := BackupConfig(path); err != nil {
			return nil, fmt.Errorf("failed to create backup: %w", err)
		}
	}
	
	// Expand home directory
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(home, path[2:])
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &HyprlandConfig{}
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse the line and populate config structure
		if err := parseLine(line, config); err != nil {
			return nil, err
		}
	}

	return config, scanner.Err()
}

func parseLine(line string, config *HyprlandConfig) error {
	// Handle monitor declarations
	if strings.HasPrefix(line, "monitor=") {
		return parseMonitor(line[8:], config)
	}

	// Handle section blocks
	if strings.Contains(line, "{") {
		return parseSection(line, config)
	}

	// Handle direct assignments
	if strings.Contains(line, "=") {
		return parseAssignment(line, config)
	}

	return nil
}

func parseMonitor(value string, config *HyprlandConfig) error {
	parts := strings.Split(value, ",")
	if len(parts) < 4 {
		return fmt.Errorf("invalid monitor configuration: %s", value)
	}

	monitor := Monitor{
		Name:       parts[0],
		Resolution: parts[1],
		Position:   parts[2],
	}

	config.Monitors = append(config.Monitors, monitor)
	return nil
}

func parseSection(line string, config *HyprlandConfig) error {
	sectionName := strings.TrimSpace(strings.Split(line, "{")[0])
	
	switch sectionName {
	case "general":
		return parseGeneralSection(line, config)
	case "decoration":
		return parseDecorationSection(line, config)
	case "animations":
		return parseAnimationsSection(line, config)
	// Add other sections as needed
	}

	return fmt.Errorf("unknown section: %s", sectionName)
}

// parseAssignment handles direct key=value assignments
func parseAssignment(line string, config *HyprlandConfig) error {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid assignment: %s", line)
	}
	
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	
	// TODO: Handle global assignments
	return fmt.Errorf("unhandled assignment: %s", key)
}

// Helper function to parse section content
type sectionParser struct {
	content string
	pos     int
	inBlock bool
}

func newSectionParser(line string) *sectionParser {
	return &sectionParser{
		content: line,
		pos:     0,
		inBlock: false,
	}
}

func (sp *sectionParser) parseBlock() (map[string]string, error) {
	values := make(map[string]string)
	var currentLine string
	
	for sp.pos < len(sp.content) {
		char := sp.content[sp.pos]
		
		switch char {
		case '{':
			sp.inBlock = true
		case '}':
			sp.inBlock = false
			return values, nil
		case '\n':
			if sp.inBlock && len(strings.TrimSpace(currentLine)) > 0 {
				parts := strings.SplitN(currentLine, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					values[key] = value
				}
			}
			currentLine = ""
		default:
			if sp.inBlock {
				currentLine += string(char)
			}
		}
		sp.pos++
	}
	
	return values, nil
}

func parseGeneralSection(line string, config *HyprlandConfig) error {
	parser := newSectionParser(line)
	values, err := parser.parseBlock()
	if err != nil {
		return err
	}
	
	// Parse values into GeneralSection struct
	for key, value := range values {
		switch key {
		case "border_size":
			if size, err := strconv.Atoi(value); err == nil {
				config.General.BorderSize = size
			}
		case "gaps_in":
			if gaps, err := strconv.Atoi(value); err == nil {
				config.General.GapIn = gaps
			}
		case "gaps_out":
			if gaps, err := strconv.Atoi(value); err == nil {
				config.General.GapOut = gaps
			}
		case "cursor_inactive_timeout":
			config.General.Cursor = value
		case "layout":
			config.General.Layout = value
		case "no_focus_fallback":
			config.General.NoFocusFollowMouse = value == "true"
		}
	}
	
	return nil
}

func parseDecorationSection(line string, config *HyprlandConfig) error {
	parser := newSectionParser(line)
	values, err := parser.parseBlock()
	if err != nil {
		return err
	}
	
	for key, value := range values {
		switch key {
		case "rounding":
			if r, err := strconv.Atoi(value); err == nil {
				config.Decoration.Rounding = r
			}
		case "blur":
			config.Decoration.BlurEnabled = value == "true"
		case "blur_size":
			if size, err := strconv.Atoi(value); err == nil {
				config.Decoration.BlurSize = size
			}
		case "blur_passes":
			if passes, err := strconv.Atoi(value); err == nil {
				config.Decoration.BlurPasses = passes
			}
		case "active_opacity":
			if opacity, err := strconv.ParseFloat(value, 64); err == nil {
				config.Decoration.Opacity = opacity
			}
		case "inactive_opacity":
			if opacity, err := strconv.ParseFloat(value, 64); err == nil {
				config.Decoration.InactiveOpacity = opacity
			}
		case "drop_shadow":
			config.Decoration.DropShadow = value == "true"
		case "shadow_range":
			if range_, err := strconv.Atoi(value); err == nil {
				config.Decoration.ShadowRange = range_
			}
		case "shadow_color":
			config.Decoration.ShadowColor = value
		}
	}
	
	return nil
}

func parseAnimationsSection(line string, config *HyprlandConfig) error {
	parser := newSectionParser(line)
	values, err := parser.parseBlock()
	if err != nil {
		return err
	}
	
	for key, value := range values {
		switch key {
		case "enabled":
			config.Animations.Enabled = value == "true"
		// Add other animation settings as needed
		}
	}
	
	return nil
}

// Add other parsing functions... 