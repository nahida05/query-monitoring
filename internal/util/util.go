package util

import "strings"

func Exists(value string, options ...string) bool {
	lowerValue := strings.ToLower(value)
	for _, option := range options {
		if lowerValue == strings.ToLower(option) {
			return true
		}
	}

	return false
}

// PageCount returns the number of pages for a given total and items per page
func PageCount(total int, limit int, defaultLimit int) int {
	if limit == 0 {
		limit = defaultLimit
	}

	pages := total / limit

	if total%limit > 0 {
		pages++
	}

	return pages
}
