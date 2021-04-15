package helpers

import (
	"strings"
)

func ExtractToken(bearer string) string {
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearer, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
