package teamroles

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

type Type string

const ( //for permission check
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

// In Default superadmin or roleid 1 have all permissions
type Config struct {
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     *auth.Auth
	DataBaseType     Type
	DB               *gorm.DB
}

type PermissionConfig struct {
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     *auth.Auth
	DB               *gorm.DB
	DataBaseType     Type
	AuthFlg          bool
	PermissionFlg    bool
	UserId           int
	DataAccess       int
}
