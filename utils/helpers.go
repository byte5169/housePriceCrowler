package utils

import "strings"

// It returns true if the string str is in the slice s, and false otherwise
//
// Args:
//   s ([]string): The slice of strings to search through.
//   str (string): The string to search for.
//
// Returns:
//   A slice of strings
func Contains(s []string, str string) bool {

	newStr := strings.ReplaceAll(str, "/", " ")
	for _, v := range s {
		if v == newStr {
			return true
		}
	}

	return false
}
