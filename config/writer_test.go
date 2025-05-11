package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteConfig(t *testing.T) {
	// Create a test config
	cfg := &HyprlandConfig{
		General: GeneralSection{
			BorderSize: 2,
			GapIn:      5,
			GapOut:     10,
			Layout:     "dwindle",
		},
		Decoration: DecorationSection{
			Rounding:    10,
			BlurEnabled: true,
			BlurSize:    3,
			Opacity:     0.95,
		},
		Binds: []Bind{
			{
				Mods:        "SUPER",
				Key:         "Return",
				Dispatcher:  "exec",
				Params:      "kitty",
				Description: "Launch terminal",
			},
		},
	}

	// Create temp file for testing
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.conf")

	// Write config
	err := WriteConfig(cfg, testFile)
	if err != nil {
		t.Fatalf("WriteConfig() error = %v", err)
	}

	// Read and verify content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Verify some expected content
	contentStr := string(content)
	if !strings.Contains(contentStr, "border_size = 2") {
		t.Error("Config missing border_size setting")
	}
	if !strings.Contains(contentStr, "bind = SUPER, Return, exec, kitty") {
		t.Error("Config missing terminal bind")
	}
} 