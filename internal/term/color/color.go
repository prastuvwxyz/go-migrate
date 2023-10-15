package color

import "github.com/fatih/color"

var (
	Grey     = color.New(color.FgWhite)
	DarkGray = color.New(color.FgBlack)
	Red      = color.New(color.FgHiRed)
	DullRed  = color.New(color.FgRed)
	Green    = color.New(color.FgHiGreen)
	Yellow   = color.New(color.FgHiYellow)
	Magenta  = color.New(color.FgMagenta)
	Blue     = color.New(color.FgHiBlue)

	DullGreen   = color.New(color.FgGreen)
	DullBlue    = color.New(color.FgBlue)
	DullYellow  = color.New(color.FgYellow)
	DullMagenta = color.New(color.FgMagenta)
	DullCyan    = color.New(color.FgCyan)

	HiBlue       = color.New(color.FgHiBlue)
	Cyan         = color.New(color.FgHiCyan)
	Bold         = color.New(color.Bold)
	Faint        = color.New(color.Faint)
	BoldFgYellow = color.New(color.FgYellow).Add(color.Bold)
)
