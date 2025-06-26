package comparelist

import (
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

func ContainsString(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func ContainsStringByRegex(slice []string, element string) bool {
	for _, v := range slice {
		regex, err := regexp.Compile(v)
		if err != nil {
			logx.Error(err)
			return false
		}
		if regex.MatchString(element) {
			return true
		}
	}
	return false
}
