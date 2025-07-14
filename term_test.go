package color

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFlags(t *testing.T) {
	tests := []struct {
		name           string
		noColorEnv     string
		forceColorEnv  string
		wantNoColor    bool
		wantForceColor bool
		setup          func()
		cleanup        func()
	}{
		{
			name:           "NO_COLOR enabled",
			wantNoColor:    true,
			wantForceColor: false,
			setup: func() {
				_ = os.Setenv("NO_COLOR", "true")
			},
			cleanup: func() {
				_ = os.Unsetenv("NO_COLOR")
				NoColor = false
			},
		},
		{
			name:           "FORCE_COLOR enabled",
			wantNoColor:    false,
			wantForceColor: true,
			setup: func() {
				_ = os.Setenv("FORCE_COLOR", "true")
			},
			cleanup: func() {
				_ = os.Unsetenv("FORCE_COLOR")
				ForceColor = false
			},
		},
		{
			name:           "both disabled",
			wantNoColor:    false,
			wantForceColor: false,
			setup:          func() {},
			cleanup:        func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()

			initFlags()

			assert.Equal(t, tc.wantNoColor, NoColor)
			assert.Equal(t, tc.wantForceColor, ForceColor)

			tc.cleanup()
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
		cleanup      func()
	}{
		{
			name: "COLORTERM",
			want: true,
			setup: func() {
				_ = os.Setenv("COLORTERM", "xterm256")
			},
			cleanup: func() {
				_ = os.Unsetenv("COLORTERM")
			},
		},
		{
			name: "TERM=dumb",
			want: false,
			setup: func() {
				_ = os.Setenv("TERM", "dumb")
			},
			cleanup: func() {
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
			cleanup: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name: "TERM=xterm-256color",
			want: true,
			setup: func() {
				_ = os.Setenv("TERM", "xterm-256color")
			},
			cleanup: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name: "TERM=screen",
			want: true,
			setup: func() {
				_ = os.Setenv("TERM", "screen")
			},
			cleanup: func() {
				_ = os.Unsetenv("TERM")
			},
		},
		{
			name: "TERM=linux",
			want: true,
			setup: func() {
				_ = os.Setenv("TERM", "linux")

			},
			cleanup: func() {
				_ = os.Unsetenv("TERM")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			result := SupportsColor()

			assert.Equal(t, tc.want, result)
			tc.cleanup()
		})
	}
}
