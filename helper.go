package spaces

import (
	"regexp"
	"strings"
)

//func helps to create slug 
func CreateSlug(str string) (string, error) {

	if str == "" {

		return "", ErrorSlug
	}

	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")

	trimstr := strings.ReplaceAll(str, " ", "-")

	return trimstr, nil

}
