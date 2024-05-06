package lib

import "strings"

func CenterString(value string, lineWidth int) string {
	sideWidth := (lineWidth - len(value)) / 2
	return strings.Repeat(" ", sideWidth) + value + strings.Repeat(" ", sideWidth)
}
