package xstrings

import (
	"pkgs/xset"
	"strings"
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

func SnakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	//snake_case to camelCase
	isToUpper := false
	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return
}

// func SnakeCase(camelStr string) string {

// 	r := make([]rune, 0)
// 	for _, v := range camelStr {

// 	}
// }
