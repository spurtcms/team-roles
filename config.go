package teamroles

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

// In Default superadmin or roleid 1 have all permissions
type Config struct {
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     *auth.Auth
	RoleId           int
	RoleName         string
	DB               *gorm.DB
}

type PermissionConfig struct {
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     *auth.Auth
	RoleId           int
	RoleName         string
	DB               *gorm.DB
	AuthFlg          bool
	PermissionFlg    bool
}
