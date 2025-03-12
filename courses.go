package courses

import (
	"fmt"

	"github.com/spurtcms/auth/migration"
)

func CoursesSetup(config Config) *Courses {

	migration.AutoMigration(config.DB, config.DataBaseType)

	fmt.Println("hello")

	return &Courses{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}
