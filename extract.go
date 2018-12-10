package main

import (
	"regexp"
)

var translatables = regexp.MustCompile(`(?m)(?:text = "|[^s]tr\(["'])(.*?)(?:"|["']\))`)

// extract extracts the translatable keys for
func extract(fileContent string) []string {
	translations := []string{}
	matches := translatables.FindAllStringSubmatch(fileContent, -1)
	for _, match := range matches {
		translations = append(translations, match[1])
	}

	return translations
}
