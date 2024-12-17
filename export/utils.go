package export

import (
	"slices"
	"strings"
)

func GenerateShorthands(args []string) []rune {
	var shorthands []rune

	for _, arg := range args {
		if len(arg) < 1 {
			continue
		}
		for _, short := range arg {
			found := slices.Contains(shorthands, rune(short))
			if !found {
				shorthands = append(shorthands, short)
				break
			}
		}
	}

	return shorthands
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s // Return as is if string is empty
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func GenerateDefaultValues(s string) string {
	if s == "string" {
		return `""`
	}

	if s == "int" {
		return `0`
	}
	return `nil`
}
