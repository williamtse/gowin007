package src

import (
	//	"net/http"
	"regexp"
)

func GetMatches(pattern string, str string) []string {
	reg := regexp.MustCompile(pattern)
	matche := reg.FindStringSubmatch(str)
	return matche
}
