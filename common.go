package spaces

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorSlug       = errors.New("slug value empty")
	CurrentTime, _  = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	Empty           string
)

func TruncateDescription(description string, limit int) string {
	if len(description) <= limit {
		return description
	}

	truncated := description[:limit] + "..."
	return truncated
}

func AuthandPermission(space *Spaces) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if space.AuthEnable && !space.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if space.PermissionEnable && !space.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

// func helps to create slug
func CreateSlug(str string) (string, error) {

	if str == "" {

		return "", ErrorSlug
	}

	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")

	trimstr := strings.ReplaceAll(str, " ", "-")

	return trimstr, nil

}

func convertarrayintToString(a []int, delim string) string {

	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
