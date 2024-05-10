package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(s string) string {
	return cases.Title(language.English, cases.NoLower).String(s)
}
