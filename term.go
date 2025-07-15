package color

import (
	"os"
	"slices"
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

var colorTerms = []string{
	"xterm", "xterm-256color", "screen", "tmux",
	"color", "linux", "cygwin", "putty", "konsole", "gnome",
}

func init() {
	initFlags()
}

func initFlags() {
	noColor := os.Getenv("NO_COLOR")
	if b, err := strconv.ParseBool(noColor); b && err == nil {
		NoColor = true
	}

	forceColor := os.Getenv("FORCE_COLOR")
	if b, err := strconv.ParseBool(forceColor); b && err == nil {
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
	if term == "dumb" || term == "" {
		return false
	}

	colorTerm := os.Getenv("COLORTERM")
	if colorTerm != "" {
		return true
	}

	return slices.Contains(colorTerms, term)
}
