package util

import (
	"errors"
)

func ParseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "yes", "YES", "Yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "no", "NO", "No":
		return false, nil
	}
	return false, errors.New("strconv.ParseBool: " + "parsing " + str)
}

// FormatBool returns "true" or "false" according to the value of b.
func FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func FormatBoolYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
