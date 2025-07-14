package color

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	tests := []struct {
		name           string
		noColorEnv     string
		forceColorEnv  string
		wantNoColor    bool
		wantForceColor bool
		setUp          func()
		reset          func()
	}{
		{
			name:           "NO_COLOR enabled",
			wantNoColor:    true,
			wantForceColor: false,
			setUp: func() {
				_ = os.Setenv("NO_COLOR", "true")
			},
			reset: func() {
				_ = os.Unsetenv("NO_COLOR")
				NoColor = false
			},
		},
		{
			name:           "FORCE_COLOR enabled",
			wantNoColor:    false,
			wantForceColor: true,
			setUp: func() {
				_ = os.Setenv("FORCE_COLOR", "true")
			},
			reset: func() {
				_ = os.Unsetenv("FORCE_COLOR")
				ForceColor = false
			},
		},
		{
			name:           "both disabled",
			wantNoColor:    false,
			wantForceColor: false,
			setUp:          func() {},
			reset:          func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setUp()

			initialize()

			assert.Equal(t, tc.wantNoColor, NoColor)
			assert.Equal(t, tc.wantForceColor, ForceColor)

			tc.reset()
		})
	}
}

func TestSupportsColor(t *testing.T) {
	tests := []struct {
		name         string
		termEnv      string
		colorTermEnv string
		want         bool
		setup        func()
		reset        func()
	}{
		{
			name: "TERM=dumb",
			want: false,
			setup: func() {
				_ = os.Setenv("TERM", "dumb")
			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name:    "TERM is empty",
			termEnv: "",
			want:    false,
			setup: func() {
				_ = os.Setenv("TERM", "")
			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name:    "TERM=xterm-256color",
			termEnv: "xterm-256color",
			want:    true,
			setup: func() {
				_ = os.Setenv("TERM", "xterm-256color")
			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name:    "TERM=screen",
			termEnv: "screen",
			want:    true,
			setup: func() {
				_ = os.Setenv("TERM", "screen")
			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name:    "TERM contains 'xterm'",
			termEnv: "my-xterm-compatible",
			want:    true,
			setup: func() {
				_ = os.Setenv("TERM", "xterm")
			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name:    "TERM=linux",
			termEnv: "linux",
			want:    true,
			setup: func() {
				_ = os.Setenv("TERM", "linux")

			},
			reset: func() {
				_ = os.Unsetenv("TERM")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			result := SupportsColor()

			assert.Equal(t, tc.want, result)
			tc.reset()
		})
	}
}
