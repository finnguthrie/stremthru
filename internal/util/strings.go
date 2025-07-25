package util

import (
	"regexp"
	"strconv"
	"strings"

	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
)

func RepeatJoin(s string, count int, sep string) string {
	if count == 0 {
		return ""
	}
	return strings.Repeat(s+sep, count-1) + s
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

var numericRegex = regexp.MustCompile(`^[0-9]+$`)

func IsNumericString(s string) bool {
	return numericRegex.MatchString(s)
}

var quoteRegex = regexp.MustCompile(`['"]+`)
var separatorRegex = regexp.MustCompile(`[._-]+`)

func normalizeForFuzzySearch(s string) string {
	s = strings.ToLower(s)
	s = quoteRegex.ReplaceAllLiteralString(s, "")
	s = separatorRegex.ReplaceAllLiteralString(s, " ")
	return fuzzy.Cleanse(s, false)
}

func FuzzyTokenSetRatio(query, input string) int {
	return fuzzy.TokenSetRatio(normalizeForFuzzySearch(query), normalizeForFuzzySearch(input))
}
