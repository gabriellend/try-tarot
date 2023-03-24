package cards

import "strings"

// ToBase lowercases and replaces spaces with hyphens in a string to
// use it as the last element (base) in a url path
func ToBase(name string) string {
	b := strings.NewReplacer(" ", "-")
	return strings.ToLower(b.Replace(name))
}
