package customstrings

import "strings"

func TrimSpace(value string) string {
	if value == "" {
		return value
	}

	return strings.TrimSpace(value)
}

func RemoveEndPeriod(value string) string {
	if value == "" {
		return value
	}

	lastCharacter := value[len(value)-1:]
	if lastCharacter == "." {
		return value[0 : len(value)-1]
	}

	return value
}

func RemoveSpecialCharacters(value string) string {
	if value == "" {
		return value
	}

	// Remove double quotes
	if strings.Contains(value, "\"") {
		value = strings.Replace(value, "\"", "", -1)
	}

	// Remove single quote
	if strings.Contains(value, "'") {
		value = strings.Replace(value, "'", "", -1)
	}

	// Remove open clasp
	if strings.Contains(value, "[") {
		value = strings.Replace(value, "[", "", -1)
	}

	// Remove closed clasp
	if strings.Contains(value, "]") {
		value = strings.Replace(value, "]", "", -1)
	}

	value = strings.ToLower(value)

	return strings.TrimSpace(value)
}
