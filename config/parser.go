package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
		parseLine(line, config)
	}

	return config, scanner.Err()
}

func parseLine(line string, config *HyprlandConfig) {
	// TODO: Implement parsing logic
	// This will need to handle all the different config sections and formats
} 