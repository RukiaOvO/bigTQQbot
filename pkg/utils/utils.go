package utils

import (
	"regexp"
	"strings"
)

func OrderCheck(str string) (string, bool) {
	if !strings.HasPrefix(str, "/") {
		return "", false
	}
	hasEmpty := strings.Contains(str, " ")
	if hasEmpty {
		str = strings.Split(str, " ")[0]
	}
	re := regexp.MustCompile("\\/[a-zA-Z][a-zA-Z_]*")
	result := re.FindString(str)
	if result == str {
		return str, true
	} else {
		return "", false
	}
}

func PixivUrlRebuild(url string) string {
	url = strings.Join(strings.Split(url, "/")[5:], "/")
	return "https://i.pixiv.cat/" + url
}

func StrIsContains(str string, array []string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}
