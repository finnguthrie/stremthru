package util

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	query = normalizeForFuzzySearch(query)
	if result, _, err := transform.String(t, query); err == nil {
		query = result
	}
	input = normalizeForFuzzySearch(input)
	if result, _, err := transform.String(t, input); err == nil {
		input = result
	}
	return fuzzy.TokenSetRatio(query, input)
}
