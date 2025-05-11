package config

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

func BackupConfig(path string) error {
	if path == "" {
		path = DefaultConfigPath
	}

	// Expand home directory
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(home, path[2:])
	}

	// Read original file
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Create backup with timestamp
	timestamp := time.Now().Format("20060102_150405")
	backupPath := path + "." + timestamp + ".backup"
	return os.WriteFile(backupPath, content, 0644)
}
