package teamroles

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// role&permission setup config
func RoleSetup(config Config) *PermissionConfig {

	return &PermissionConfig{
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		RoleId:           config.RoleId,
		RoleName:         config.RoleName,
		DB:               config.DB,
		Authenticate:     config.Authenticate,
	}

}

var AS ModelStruct

// create role
func (RoleConf PermissionConfig) RoleList(rolelist rolelist) (roles []tblrole, rolecount int64, err error) {

	fmt.Println("auth", RoleConf.Authenticate)

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return []tblrole{}, 0, ErrorAuth
	}	

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return []tblrole{}, 0, ErrorPermission
	}

	role, _, _ := AS.GetAllRoles(rolelist.Limit, rolelist.Offset, rolelist.filter, rolelist.GetAllData, RoleConf.DB)

	_, rolecounts, _ := AS.GetAllRoles(0, 0, rolelist.filter, rolelist.GetAllData, RoleConf.DB)

	return role, rolecounts, nil

}

// get role by id
func (RoleConf PermissionConfig) GetRoleById(roleid int) (tblrol tblrole, err error) {

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return tblrole{}, ErrorAuth
	}

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return tblrole{}, ErrorPermission
	}

	var AS ModelStruct

	role, err := AS.GetRoleById(roleid, RoleConf.DB)

	if err != nil {

		return tblrole{}, err
	}

	return role, nil

}

// create role
func (RoleConf PermissionConfig) CreateRole(rolec RoleCreation) (tblrole, error) {

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return tblrole{}, ErrorAuth
	}

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return tblrole{}, ErrorPermission
	}

	if rolec.Name == "" {

		return tblrole{}, errors.New("can't store role name is empty")
	}

	var role tblrole

	role.Name = rolec.Name

	role.Description = rolec.Description

	role.Slug = strings.ToLower(role.Name)

	role.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	role.CreatedBy = rolec.CreatedBy

	err := AS.RoleCreate(&role, RoleConf.DB)

	if err != nil {

		return tblrole{}, err
	}

	return role, nil

}

// update role
func (RoleConf PermissionConfig) UpdateRole(rolec RoleCreation, roleid int) (updaterole tblrole, err error) {

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return tblrole{}, ErrorAuth
	}

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return tblrole{}, ErrorPermission
	}

	if rolec.Name == "" {

		return tblrole{}, errors.New("empty value")
	}

	var role tblrole

	role.Id = roleid

	role.Name = rolec.Name

	role.Description = rolec.Description

	role.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	role.ModifiedBy = rolec.CreatedBy

	err1 := AS.RoleUpdate(&role, RoleConf.DB)

	if err1 != nil {

		return tblrole{}, err1
	}

	return role, nil

}

// delete role
func (RoleConf PermissionConfig) DeleteRole(roleid int) (bool, error) {

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return false, ErrorAuth
	}

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return false, ErrorPermission
	}

	if roleid <= 0 {

		return false, errors.New("invalid roleid cannot delete")
	}

	var role tblrole

	err1 := AS.RoleDelete(&role, roleid, RoleConf.DB)

	var permissions []tblrolepermission

	AS.DeleteRolePermissionById(&permissions, roleid, RoleConf.DB)

	if err1 != nil {

		return false, err1
	}

	return true, nil

}

/*Check Role Already Exists*/
func (RoleConf PermissionConfig) CheckRoleAlreadyExists(roleid int, rolename string) (bool, error) {

	if RoleConf.AuthEnable && !RoleConf.Authenticate.AuthFlg {

		return false, ErrorAuth
	}

	if RoleConf.PermissionEnable && !RoleConf.Authenticate.PermissionFlg {

		return false, ErrorPermission
	}

	var role tblrole

	err1 := AS.CheckRoleExists(&role, roleid, rolename, RoleConf.DB)

	if err1 != nil {

		return false, err1
	}

	if role.Id == 0 {

		return false, nil
	}

	return true, nil
}
