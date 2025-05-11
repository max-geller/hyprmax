package config

import (
	"testing"
)

func TestParseMonitor(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid monitor config",
			input:   "monitor=eDP-1,1920x1080,0x0,1",
			wantErr: false,
		},
		{
			name:    "invalid monitor config",
			input:   "monitor=eDP-1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &HyprlandConfig{}
			err := parseLine(tt.input, config)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseSection(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(*testing.T, *HyprlandConfig)
	}{
		{
			name: "valid general section",
			input: `general {
				border_size=2
				gaps_in=5
				gaps_out=10
				layout=dwindle
			}`,
			wantErr: false,
			check: func(t *testing.T, c *HyprlandConfig) {
				if c.General.BorderSize != 2 {
					t.Errorf("expected BorderSize=2, got %d", c.General.BorderSize)
				}
				if c.General.GapIn != 5 {
					t.Errorf("expected GapIn=5, got %d", c.General.GapIn)
				}
			},
		},
		{
			name: "valid decoration section",
			input: `decoration {
				rounding=10
				blur=true
				blur_size=3
				active_opacity=0.95
			}`,
			wantErr: false,
			check: func(t *testing.T, c *HyprlandConfig) {
				if c.Decoration.Rounding != 10 {
					t.Errorf("expected Rounding=10, got %d", c.Decoration.Rounding)
				}
				if !c.Decoration.BlurEnabled {
					t.Error("expected BlurEnabled=true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &HyprlandConfig{}
			err := parseLine(tt.input, config)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.check != nil {
				tt.check(t, config)
			}
		})
	}
}
