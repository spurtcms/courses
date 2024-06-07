package spaces

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

type Type string

const (
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	DataBaseType     Type
	Auth             *auth.Auth
}

type Spaces struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Auth             *auth.Auth
}
