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
