package utils

import (
	"regexp"
	"strings"
	"fmt"
)

// Return a map from a named capture group regex
// Keys are the names, values are the captured regex
func ReSubMatchMap(r *regexp.Regexp, str string) (map[string]string) {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
return subMatchMap
}

// Delete an empty string from a slice
func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Return first n lines of a multi-line string
// Returns both list and string for convenience
func GetHeader(lines string, keepNLines int) ([]string, string) {
	var header []string
	var headerString string
	for i, line := range strings.Split(lines, "\n") {
		if i <= keepNLines {
			header = append(header, line, "\n")
		}
	}
	headerString = strings.Join(header, "")
	return header, headerString
}

func KeepSliceReMatches(s []string, re *regexp.Regexp) []string {
	var matchLines []string
	for _, l := range s {
		if re.MatchString(l) {
			matchLines = append(matchLines, l)
		}
	}
	return matchLines
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// Generic error checker
func CheckErr(err error) {
	if err != nil {
		fmt.Println("ERRROR:", err)
	}
}
