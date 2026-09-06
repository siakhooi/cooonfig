/*
Package application run the application
*/
package application

import "github.com/fatih/color"

const (
	// ColorDefault represents default color (no color)
	ColorDefault string = "default"

	// ColorRed represents red color
	ColorRed string = "red"
	// ColorGreen represents green color
	ColorGreen string = "green"
	// ColorBlue represents blue color
	ColorBlue string = "blue"
	// ColorYellow represents yellow color
	ColorYellow string = "yellow"
	// ColorMagenta represents magenta color
	ColorMagenta string = "magenta"
	// ColorCyan represents cyan color
	ColorCyan string = "cyan"
	// ColorWhite represents white color
	ColorWhite string = "white"
	// ColorBlack represents black color
	ColorBlack string = "black"

	// ColorHiRed represents high intensity red color
	ColorHiRed string = "hiRed"
	// ColorHiGreen represents high intensity green color
	ColorHiGreen string = "hiGreen"
	// ColorHiBlue represents high intensity blue color
	ColorHiBlue string = "hiBlue"
	// ColorHiYellow represents high intensity yellow color
	ColorHiYellow string = "hiYellow"
	// ColorHiMagenta represents high intensity magenta color
	ColorHiMagenta string = "hiMagenta"
	// ColorHiCyan represents high intensity cyan color
	ColorHiCyan string = "hiCyan"
	// ColorHiWhite represents high intensity white color
	ColorHiWhite string = "hiWhite"
	// ColorHiBlack represents high intensity black color
	ColorHiBlack string = "hiBlack"

	// ColorBgRed represents red background color
	ColorBgRed string = "bgRed"
	// ColorBgGreen represents green background color
	ColorBgGreen string = "bgGreen"
	// ColorBgBlue represents blue background color
	ColorBgBlue string = "bgBlue"
	// ColorBgYellow represents yellow background color
	ColorBgYellow string = "bgYellow"
	// ColorBgMagenta represents magenta background color
	ColorBgMagenta string = "bgMagenta"
	// ColorBgCyan represents cyan background color
	ColorBgCyan string = "bgCyan"
	// ColorBgWhite represents white background color
	ColorBgWhite string = "bgWhite"
	// ColorBgBlack represents black background color
	ColorBgBlack string = "bgBlack"

	// ColorHiBgRed represents high intensity red background color
	ColorHiBgRed string = "hiBgRed"
	// ColorHiBgGreen represents high intensity green background color
	ColorHiBgGreen string = "hiBgGreen"
	// ColorHiBgBlue represents high intensity blue background color
	ColorHiBgBlue string = "hiBgBlue"
	// ColorHiBgYellow represents high intensity yellow background color
	ColorHiBgYellow string = "hiBgYellow"
	// ColorHiBgMagenta represents high intensity magenta background color
	ColorHiBgMagenta string = "hiBgMagenta"
	// ColorHiBgCyan represents high intensity cyan background color
	ColorHiBgCyan string = "hiBgCyan"
	// ColorHiBgWhite represents high intensity white background color
	ColorHiBgWhite string = "hiBgWhite"
	// ColorHiBgBlack represents high intensity black background color
	ColorHiBgBlack string = "hiBgBlack"

	// ColorBold represents bold text
	ColorBold string = "bold"
	// ColorFaint represents faint text
	ColorFaint string = "faint"
	// ColorItalic represents italic text
	ColorItalic string = "italic"
	// ColorUnderline represents underlined text
	ColorUnderline string = "underline"
	// ColorBlinkSlow represents slow blinking text
	ColorBlinkSlow string = "blinkSlow"
	// ColorBlinkRapid represents rapid blinking text
	ColorBlinkRapid string = "blinkRapid"
	// ColorReverseVideo represents reverse video text
	ColorReverseVideo string = "reverseVideo"
	// ColorConcealed represents concealed text
	ColorConcealed string = "concealed"
	// ColorCrossedOut represents crossed out text
	ColorCrossedOut string = "crossedOut"
)

type colorMeta struct {
	color color.Attribute
}

var supportedColorMeta = map[string]colorMeta{
	ColorDefault:      {color: color.Reset},
	ColorRed:          {color: color.FgRed},
	ColorGreen:        {color: color.FgGreen},
	ColorBlue:         {color: color.FgBlue},
	ColorYellow:       {color: color.FgYellow},
	ColorMagenta:      {color: color.FgMagenta},
	ColorCyan:         {color: color.FgCyan},
	ColorWhite:        {color: color.FgWhite},
	ColorBlack:        {color: color.FgBlack},
	ColorHiRed:        {color: color.FgHiRed},
	ColorHiGreen:      {color: color.FgHiGreen},
	ColorHiBlue:       {color: color.FgHiBlue},
	ColorHiYellow:     {color: color.FgHiYellow},
	ColorHiMagenta:    {color: color.FgHiMagenta},
	ColorHiCyan:       {color: color.FgHiCyan},
	ColorHiWhite:      {color: color.FgHiWhite},
	ColorHiBlack:      {color: color.FgHiBlack},
	ColorBgRed:        {color: color.BgRed},
	ColorBgGreen:      {color: color.BgGreen},
	ColorBgBlue:       {color: color.BgBlue},
	ColorBgYellow:     {color: color.BgYellow},
	ColorBgMagenta:    {color: color.BgMagenta},
	ColorBgCyan:       {color: color.BgCyan},
	ColorBgWhite:      {color: color.BgWhite},
	ColorBgBlack:      {color: color.BgBlack},
	ColorHiBgRed:      {color: color.BgHiRed},
	ColorHiBgGreen:    {color: color.BgHiGreen},
	ColorHiBgBlue:     {color: color.BgHiBlue},
	ColorHiBgYellow:   {color: color.BgHiYellow},
	ColorHiBgMagenta:  {color: color.BgHiMagenta},
	ColorHiBgCyan:     {color: color.BgHiCyan},
	ColorHiBgWhite:    {color: color.BgHiWhite},
	ColorHiBgBlack:    {color: color.BgHiBlack},
	ColorBold:         {color: color.Bold},
	ColorFaint:        {color: color.Faint},
	ColorItalic:       {color: color.Italic},
	ColorUnderline:    {color: color.Underline},
	ColorBlinkSlow:    {color: color.BlinkSlow},
	ColorBlinkRapid:   {color: color.BlinkRapid},
	ColorReverseVideo: {color: color.ReverseVideo},
	ColorConcealed:    {color: color.Concealed},
	ColorCrossedOut:   {color: color.CrossedOut},
}

func isValidTextColor(colorName string) bool {
	_, ok := supportedColorMeta[colorName]
	return ok
}
