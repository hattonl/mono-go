package xstrings

import (
	"utils-go/xset"
)

func ArrayContains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func ArrayContainsAny(arr, s []string) bool {
	iset := xset.NewSetString()
	for _, a := range arr {
		iset.Insert(a)
	}
	for _, a := range s {
		if iset.Contains(a) {
			return true
		}
	}
	return false
}

func IsOneOf(s string, arr []string) bool {
	return ArrayContains(arr, s)
}
