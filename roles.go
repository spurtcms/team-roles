package teamroles

import (
	"fmt"
	"strings"
	"time"

	"github.com/spurtcms/team-roles/migration"
)

// role&permission setup config
func RoleSetup(config Config) *PermissionConfig {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &PermissionConfig{
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		DB:               config.DB,
		DataBaseType:     config.DataBaseType,
		Authenticate:     config.Authenticate,
	}

}

var AS ModelStruct

// role list
func (RoleConf *PermissionConfig) RoleList(rolelist Rolelist, tenantid int) (roles []Tblrole, rolecount int64, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return []Tblrole{}, 0, autherr
	}

	AS.DataAccess = RoleConf.DataAccess
	AS.UserId = RoleConf.UserId

	role, _, errr := AS.GetAllRoles(rolelist.Limit, rolelist.Offset, rolelist.Filter, rolelist.GetAllData, RoleConf.DB, tenantid)

	if errr != nil {

		fmt.Println(errr)
	}

	_, rolecounts, _ := AS.GetAllRoles(0, 0, rolelist.Filter, rolelist.GetAllData, RoleConf.DB, tenantid)

	return role, rolecounts, nil

}

// get role by id
func (RoleConf *PermissionConfig) GetRoleById(roleid int, tenantid int) (tblrol Tblrole, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return Tblrole{}, autherr
	}

	var AS ModelStruct

	role, err := AS.GetRoleById(roleid, RoleConf.DB, tenantid)

	if err != nil {

		return Tblrole{}, err
	}

	return role, nil

}

// create role
func (RoleConf *PermissionConfig) CreateRole(rolec RoleCreation) (Tblrole, error) {

	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return Tblrole{}, autherr
	}

	if rolec.Name == "" {

		return Tblrole{}, ErrorRoleNameEmpty
	}

	var role Tblrole

	role.Name = rolec.Name

	role.Description = rolec.Description

	role.Slug = strings.ToLower(role.Name)

	role.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	role.CreatedBy = rolec.CreatedBy

	role.TenantId = rolec.TenantId

	err := AS.RoleCreate(&role, RoleConf.DB)

	if err != nil {

		return Tblrole{}, err
	}

	return role, nil

}

// update role
func (RoleConf *PermissionConfig) UpdateRole(rolec RoleCreation, roleid int, tenantid int) (updaterole Tblrole, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return Tblrole{}, autherr
	}

	if rolec.Name == "" {

		return Tblrole{}, ErrorRoleNameEmpty
	}

	var role Tblrole

	role.Id = roleid

	role.Name = rolec.Name

	role.Description = rolec.Description

	role.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	role.ModifiedBy = rolec.CreatedBy

	err1 := AS.RoleUpdate(&role, RoleConf.DB, tenantid)

	if err1 != nil {

		return Tblrole{}, err1
	}

	return role, nil

}

// delete seleted role
func (RoleConf *PermissionConfig) DeleteRole(roleids []int, roleid int, tenantid int) (bool, error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return false, autherr
	}

	var role TblRole

	err1 := AS.MultiSelectRoleDelete(&role, roleids, roleid, RoleConf.DB, tenantid)

	var permissions []TblRolePermission

	AS.MultiSelectDeleteRolePermissionById(&permissions, roleids, roleid, RoleConf.DB, tenantid)

	if err1 != nil {

		return false, err1
	}

	return true, nil

}

/*Check Role Already Exists*/
func (RoleConf *PermissionConfig) CheckRoleAlreadyExists(roleid int, rolename string, tenantid int) (bool, error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return false, autherr
	}

	var role TblRole

	err1 := AS.CheckRoleExists(&role, roleid, rolename, RoleConf.DB, tenantid)

	if err1 != nil {

		return false, err1
	}

	if role.Id == 0 {

		return false, nil
	}

	return true, nil
}

// update selected role status
func (RoleConf *PermissionConfig) MultiSelectRoleStatus(roleid []int, status int, userid int, tenantid int) (err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return autherr
	}

	var rolestatus TblRole

	rolestatus.ModifiedBy = userid

	rolestatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := AS.MultiSelectRoleIsActive(&rolestatus, roleid, status, RoleConf.DB, tenantid)

	if err1 != nil {

		return err1
	}

	return nil

}

// change role status 0-inactive, 1-active
func (RoleConf *PermissionConfig) RoleStatus(roleid int, status int, userid int, tenantid int) (err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(RoleConf); autherr != nil {

		return autherr
	}

	var rolestatus TblRole

	rolestatus.ModifiedBy = userid

	rolestatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := AS.RoleIsActive(&rolestatus, roleid, status, RoleConf.DB, tenantid)

	if err1 != nil {

		return err1
	}

	return nil

}

func (RoleConf *PermissionConfig) GetRoleByName(tenantid int) (tblrole []TblRole, err error) {

	if autherr := AuthandPermission(RoleConf); autherr != nil {
		return []TblRole{}, autherr
	}

	var role []TblRole

	AS.GetRoleByName(&role, RoleConf.DB, tenantid)

	return role, nil

}
