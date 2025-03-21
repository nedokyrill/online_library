package Utils

import "strings"

func SplitVerses(text string) []string {
	if text == "" {
		return []string{}
	}
	cleanText := strings.ReplaceAll(text, "\\n", "\n")
	return strings.Split(cleanText, "\n")
}
