package spaces

import "errors"

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
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
	if space.PermissionEnable && !space.Permissions.PermissionFlg {

		return ErrorPermission

	}

	return nil
}
