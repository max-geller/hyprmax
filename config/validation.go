package config

import (
	"fmt"
	"strconv"
	"strings"
)

// Add at the top of the file, after imports
var validKeys = map[string]bool{
	"Return": true, "Space": true, "Tab": true, "Backspace": true,
}

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (got: %s)", e.Field, e.Message, e.Value)
}

// Validator functions for different types of values
func ValidateInt(field, value string, min, max int) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return &ValidationError{field, value, "must be a number"}
	}
	if v < min || v > max {
		return &ValidationError{field, value, fmt.Sprintf("must be between %d and %d", min, max)}
	}
	return nil
}

func ValidateFloat(field, value string, min, max float64) error {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return &ValidationError{field, value, "must be a number"}
	}
	if v < min || v > max {
		return &ValidationError{field, value, fmt.Sprintf("must be between %.2f and %.2f", min, max)}
	}
	return nil
}

func ValidateBool(field, value string) error {
	value = strings.ToLower(value)
	if value != "true" && value != "false" {
		return &ValidationError{field, value, "must be true or false"}
	}
	return nil
}

func ValidateKey(field, value string) error {
	if !validKeys[value] {
		return &ValidationError{field, value, "invalid key name"}
	}
	return nil
}

// Add more validation functions

func ValidateWindowRule(field, value string) error {
	parts := strings.Split(value, ",")
	if len(parts) != 3 {
		return &ValidationError{field, value, "must be in format: rule,value,target"}
	}
	// Add specific rule validation
	return nil
}

func ValidateKeybind(field, value string) error {
	parts := strings.Split(value, ",")
	if len(parts) < 3 {
		return &ValidationError{field, value, "must be in format: mods,key,dispatcher[,params]"}
	}

	// Validate mods
	validMods := map[string]bool{"SUPER": true, "ALT": true, "CTRL": true, "SHIFT": true}
	mods := strings.Split(parts[0], "+")
	for _, mod := range mods {
		if !validMods[strings.TrimSpace(strings.ToUpper(mod))] {
			return &ValidationError{field, mod, "invalid modifier key"}
		}
	}

	// Validate key
	if err := ValidateKey(field, parts[1]); err != nil {
		return err
	}

	return nil
}

// Add more specific validators as needed

func ValidateRuleType(field, value string) error {
	validRules := map[string]bool{
		"workspace": true,
		"float":     true,
		"tile":      true,
		"pseudo":    true,
		"size":      true,
		"minsize":   true,
		"maxsize":   true,
		"opacity":   true,
		// Add other valid rule types
	}

	if !validRules[strings.ToLower(value)] {
		return &ValidationError{field, value, "invalid rule type"}
	}
	return nil
}

func ValidateDispatcher(field, value string) error {
	validDispatchers := map[string]bool{
		"exec":           true,
		"killactive":     true,
		"workspace":      true,
		"togglefloating": true,
		"fullscreen":     true,
		"pseudo":         true,
		"movefocus":      true,
		"movewindow":     true,
		// Add other valid dispatchers
	}

	if !validDispatchers[strings.ToLower(value)] {
		return &ValidationError{field, value, "invalid dispatcher"}
	}
	return nil
}

// Enhance ValidateKey with more keys
func init() {
	// Add to validKeys map
	for i := 0; i < 10; i++ {
		validKeys[fmt.Sprintf("%d", i)] = true
	}
	for c := 'A'; c <= 'Z'; c++ {
		validKeys[string(c)] = true
	}
	// Add function keys
	for i := 1; i <= 12; i++ {
		validKeys[fmt.Sprintf("F%d", i)] = true
	}
	// Add more special keys
	specialKeys := []string{
		"Escape", "Print", "Insert", "Delete",
		"Home", "End", "PageUp", "PageDown",
		"Left", "Right", "Up", "Down",
	}
	for _, key := range specialKeys {
		validKeys[key] = true
	}
}
