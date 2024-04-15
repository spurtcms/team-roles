package teamroles

import "errors"

var (
	ErrorAuthentication           = errors.New("auth enabled not initialised")
	ErrorAuthorization            = errors.New("permissions enabled not initialised")
	Error500Status                = errors.New("internal server error")
	Error403                      = errors.New("unauthorized")
	Error404                      = errors.New("not found")
	ErrorModuleNotFound           = errors.New("modulename not found")
	ErrorModulePermissionNotFound = errors.New("module permission not found")
	ErrorRoleNameEmpty            = errors.New("role name is empty")
	ErrorInvalidroleid            = errors.New("invalid roleid cannot delete")
)

func AuthandPermission(permission *PermissionConfig) error {

	if permission.AuthEnable && !permission.Authenticate.AuthFlg {

		return ErrorAuthentication
	}

	if permission.PermissionEnable && !permission.Authenticate.PermissionFlg {

		return ErrorAuthorization
	}

	return nil
}
