package teamroles

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

// In Default superadmin and roleid 1 have all permissions
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
}

type Permissions struct {
	DBString         *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	AuthFlg          bool
	PermissionFlg    bool
}
