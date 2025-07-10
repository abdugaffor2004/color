# Terminal Color

A lightweight Go package for colorful and styled terminal output with ANSI escape sequences support.

## Features

- Support for all 16 basic ANSI colors (8 normal + 8 bright) for foreground and background
- Text styling: bold, italic, underline, dim
- Automatic terminal capability detection
- Environment variable support (NO_COLOR, FORCE_COLOR)
- Thread-safe operations with optimized caching
- Cross-platform compatibility (Unix/Linux/macOS, Windows 10+)
- Zero dependencies

## Installation

```bash
go get github.com/abdugaffor2004/terminal-color
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/abdugaffor2004/terminal-color"
)

func main() {
    fmt.Println(color.Red("Error: file not found"))
    fmt.Println(color.Green("Success: operation completed"))
    fmt.Println(color.Yellow("Warning: deprecated method"))
    fmt.Println(color.Blue("Info: processed 100 files"))
}
```

### Advanced Usage

```go
// Combining multiple styles
fmt.Println(color.Style("Critical Error", color.AttrFgRed, color.AttrBold, color.AttrUnderline))

// Background colors
fmt.Println(color.Style("Alert", color.AttrFgWhite, color.AttrBgRed, color.AttrBold))

// Bright colors
fmt.Println(color.BrightGreen("Tests Passed"))
fmt.Println(color.BrightRed("Critical Issue"))

// Formatted output
fmt.Println(color.Redf("Error: %s", err.Error()))
fmt.Println(color.Greenf("Processed %d of %d files", processed, total))
```

### Environment Control

```go
// Global color control
color.NoColor = true    // Disable all colors
color.ForceColor = true // Force colors even in non-TTY

// Check terminal capabilities
if color.IsTerminal() && color.SupportsColor() {
    fmt.Println(color.Green("Terminal supports colors"))
} else {
    fmt.Println("Plain text output")
}
```

## API

### Basic Color Functions

- `Red(text string) string` - Red text
- `Green(text string) string` - Green text
- `Yellow(text string) string` - Yellow text
- `Blue(text string) string` - Blue text
- `Magenta(text string) string` - Magenta text
- `Cyan(text string) string` - Cyan text
- `White(text string) string` - White text
- `Black(text string) string` - Black text

### Bright Color Functions

- `BrightRed(text string) string` - Bright red text
- `BrightGreen(text string) string` - Bright green text
- `BrightYellow(text string) string` - Bright yellow text
- `BrightBlue(text string) string` - Bright blue text
- `BrightMagenta(text string) string` - Bright magenta text
- `BrightCyan(text string) string` - Bright cyan text
- `BrightWhite(text string) string` - Bright white text
- `BrightBlack(text string) string` - Bright black text

### Background Color Functions

- `BgRed(text string) string` - Red background
- `BgGreen(text string) string` - Green background
- `BgYellow(text string) string` - Yellow background
- `BgBlue(text string) string` - Blue background
- `BgMagenta(text string) string` - Magenta background
- `BgCyan(text string) string` - Cyan background
- `BgWhite(text string) string` - White background
- `BgBlack(text string) string` - Black background

### Bright Background Color Functions

- `BgBrightRed(text string) string` - Bright red background
- `BgBrightGreen(text string) string` - Bright green background
- `BgBrightYellow(text string) string` - Bright yellow background
- `BgBrightBlue(text string) string` - Bright blue background
- `BgBrightMagenta(text string) string` - Bright magenta background
- `BgBrightCyan(text string) string` - Bright cyan background
- `BgBrightWhite(text string) string` - Bright white background
- `BgBrightBlack(text string) string` - Bright black background

### Text Style Functions

- `Bold(text string) string` - Bold text
- `Italic(text string) string` - Italic text
- `Underline(text string) string` - Underlined text
- `Dim(text string) string` - Dimmed text

### Formatted Functions

- `Redf(format string, a ...interface{}) string` - Red formatted text
- `Greenf(format string, a ...interface{}) string` - Green formatted text
- `Yellowf(format string, a ...interface{}) string` - Yellow formatted text
- `Bluef(format string, a ...interface{}) string` - Blue formatted text
- `Magentaf(format string, a ...interface{}) string` - Magenta formatted text
- `Cyanf(format string, a ...interface{}) string` - Cyan formatted text
- `Whitef(format string, a ...interface{}) string` - White formatted text
- `Blackf(format string, a ...interface{}) string` - Black formatted text

### Advanced Functions

- `Style(text string, attrs ...Attr) string` - Apply multiple attributes
- `IsTerminal() bool` - Check if output is connected to terminal
- `SupportsColor() bool` - Check if terminal supports colors

### Attributes

```go
type Attr int

const (
    // Foreground colors
    AttrFgBlack, AttrFgRed, AttrFgGreen, AttrFgYellow,
    AttrFgBlue, AttrFgMagenta, AttrFgCyan, AttrFgWhite,
    
    // Bright foreground colors
    AttrFgBrightBlack, AttrFgBrightRed, AttrFgBrightGreen, AttrFgBrightYellow,
    AttrFgBrightBlue, AttrFgBrightMagenta, AttrFgBrightCyan, AttrFgBrightWhite,
    
    // Background colors
    AttrBgBlack, AttrBgRed, AttrBgGreen, AttrBgYellow,
    AttrBgBlue, AttrBgMagenta, AttrBgCyan, AttrBgWhite,
    
    // Bright background colors
    AttrBgBrightBlack, AttrBgBrightRed, AttrBgBrightGreen, AttrBgBrightYellow,
    AttrBgBrightBlue, AttrBgBrightMagenta, AttrBgBrightCyan, AttrBgBrightWhite,
    
    // Text styles
    AttrBold, AttrItalic, AttrUnderline, AttrDim
)
```

### Global Variables

- `NoColor bool` - Disable all colors (also controlled by NO_COLOR env var)
- `ForceColor bool` - Force colors even in non-TTY (also controlled by FORCE_COLOR env var)


## Color Reference

| Input | Output |
|-------|--------|
| `color.Red("text")` | <span style="color: red;">text</span> |
| `color.Green("text")` | <span style="color: green;">text</span> |
| `color.Yellow("text")` | <span style="color: yellow;">text</span> |
| `color.Blue("text")` | <span style="color: blue;">text</span> |
| `color.BrightRed("text")` | <span style="color: #ff6b6b;">text</span> |
| `color.BgRed("text")` | <span style="background-color: red; color: white;">text</span> |
| `color.Bold("text")` | <span style="font-weight: bold;">text</span> |
| `color.Style("text", color.AttrFgRed, color.AttrBold)` | <span style="color: red; font-weight: bold;">text</span> |

## Environment Variables

- `NO_COLOR=1` - Disable all colors (follows [no-color.org](https://no-color.org/) standard)
- `FORCE_COLOR=1` - Force colors even when not in terminal
- `TERM=dumb` - Automatically disables colors
- `COLORTERM=truecolor` - Enables color support

## Platform Support

- **Unix/Linux/macOS**: Full support out of the box
- **Windows 10+**: Full support with modern terminals
- **Windows Terminal**: Full support
- **Git Bash**: Full support
- **WSL**: Full support as Unix environment
- **Older Windows**: Colors automatically disabled, use `FORCE_COLOR=1` if needed

## Performance

- ANSI sequences are pre-computed and cached
- Thread-safe operations with optimized locking
- Minimal memory allocations for typical usage
- Efficient string building with `strings.Builder`


