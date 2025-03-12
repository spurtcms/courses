package courses

import (
	"errors"
	"os"
	"strconv"
	"time"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorFormName   = errors.New("given some values is empty")
	TZONE, _        = time.LoadLocation(os.Getenv("TIME_ZONE"))
	TenantId, _     = strconv.Atoi(os.Getenv("Tenant_ID"))
)

func AuthandPermission(Course *Courses) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if Course.AuthEnable && !Course.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if Course.PermissionEnable && !Course.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}
