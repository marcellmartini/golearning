package logs

import (
	"strings"
)

var applicationLogLineIdentifier = map[string]rune{
	"recommendation": '❗',
	"search":         '🔍',
	"weather":        '☀',
}

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, v := range log {
		switch v {
		case applicationLogLineIdentifier["recommendation"]:
			return "recommendation"
		case applicationLogLineIdentifier["search"]:
			return "search"
		case applicationLogLineIdentifier["weather"]:
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
