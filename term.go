package color

import (
	"os"
	"strconv"

	"golang.org/x/term"
)

var (
	// NoColor disables color output if set to true.
	// This can be controlled via the NO_COLOR environment variable.
	NoColor bool

	// ForceColor forces color output, its ignore terminal support check.
	// This can be controlled via the FORCE_COLOR environment variable.
	ForceColor bool
)

func init() {
	if b, err := strconv.ParseBool(os.Getenv("NO_COLOR")); b && err == nil {
		NoColor = true
	}

	if b, err := strconv.ParseBool(os.Getenv("FORCE_COLOR")); b && err == nil {
		ForceColor = true
	}
}

func allowColor() bool {
	if NoColor {
		return false
	}

	if ForceColor {
		return true
	}

	return IsTerminal() && SupportsColor()
}

// IsTerminal reports whether the standard output is connected to a terminal (TTY).
func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// SupportsColor reports whether the current environment supports color output,
// based on the TERM and COLORTERM environment variables.
func SupportsColor() bool {
	term := os.Getenv("TERM")
	if term == "dumb" {
		return false
	}

	if os.Getenv("COLORTERM") != "" {
		return true
	}

	return false
}
