package main

import "strings"

func Find(a string) bool {
	if strings.Contains(a, "W") {
		return true
	}
	return false

}
