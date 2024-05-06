package lib

import (
	"os"
	"strings"
)

func CenterString(value string, lineWidth int) string {
	sideWidth := (lineWidth - len(value)) / 2
	return strings.Repeat(" ", sideWidth) + value + strings.Repeat(" ", sideWidth)
}

/* Converts a string to a byte buffer using extended ASCII table.*/
func String2ExtASCII(str string) []byte {

	var buffer = make([]byte, len(str))

	for i, v := range str {
		// fmt.Printf("%X ", int(v))
		buffer[i] = byte(int(v))
	}

	return buffer
}

func GetEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
