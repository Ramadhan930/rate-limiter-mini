package detector

import (
	"strings"
)

func IsSQLInjection(input string) bool {
	input = strings.ToLower(input)

	patterns := []string{
		"or 1=1",
		"or '1'='1",
		"union select",
		"' or '",
		"--",
		"drop table",
	}

	for _, pattern := range patterns {
		if strings.Contains(input, pattern) {
			return true
		}
	}
	return false
}
