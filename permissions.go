package teamroles

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// create permission
func (permission *PermissionConfig) CreatePermission(Perm MultiPermissin) error {

	if autherr := AuthandPermission(permission); autherr != nil {

		return autherr
	}

	var createrolepermission []TblRolePermission

	for _, roleperm := range Perm.Ids {

		var createmod TblRolePermission

		createmod.PermissionId = roleperm

		createmod.RoleId = Perm.RoleId

		createmod.CreatedBy = Perm.CreatedBy

		createmod.TenantId = Perm.TenantId

		createmod.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		createrolepermission = append(createrolepermission, createmod)

	}

	var err error

	if len(createrolepermission) != 0 {

		err = AS.CreateRolePermission(&createrolepermission, permission.DB)

	}

	return err
}

// create permission
func (permission *PermissionConfig) CreatePermission1(permissions CreatePermissions) error {

	if autherr := AuthandPermission(permission); autherr != nil {

		return autherr
	}

	module, err := AS.CheckModuleExists(permissions.ModuleName, permission.DB, permissions.TenantId)

	if err == gorm.ErrRecordNotFound {

		return ErrorModuleNotFound
	}

	modper, moderr := AS.CheckModulePemissionExists(module.Id, permissions.Permission, permission.DB, permissions.TenantId)

	if moderr == gorm.ErrRecordNotFound {

		return ErrorModulePermissionNotFound
	}

	var createmod TblRolePermission

	createmod.PermissionId = modper.Id

	createmod.RoleId = permissions.RoleId

	createmod.CreatedBy = permissions.CreatedBy

	createmod.TenantId = permissions.TenantId

	createmod.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	return nil
}

// update permission
func (permission *PermissionConfig) CreateUpdatePermission(Perm MultiPermissin) error {

	if autherr := AuthandPermission(permission); autherr != nil {

		return autherr
	}

	checknotexist, cnerr := AS.CheckPermissionIdNotExist(Perm.RoleId, Perm.Ids, permission.DB, Perm.TenantId)

	if len(Perm.Ids) == 0 {

		AS.Deleterolepermission(Perm.RoleId, permission.DB, Perm.TenantId)
	}

	if cnerr != nil {

		fmt.Println(cnerr)

	} else if len(checknotexist) != 0 {

		AS.DeleteRolePermissionById(Perm.RoleId, permission.DB, Perm.TenantId)
	}

	checkexist, cerr := AS.CheckPermissionIdExist(Perm.RoleId, Perm.Ids, permission.DB, Perm.TenantId)

	if cerr != nil {

		fmt.Println(cerr)

	}

	var existid []int

	for _, exist := range checkexist {

		existid = append(existid, exist.PermissionId)

	}

	pid := Difference(Perm.Ids, existid)

	var createrolepermission []TblRolePermission

	for _, roleperm := range pid {

		var createmod TblRolePermission

		createmod.PermissionId = roleperm

		createmod.RoleId = Perm.RoleId

		createmod.TenantId = Perm.TenantId

		createmod.CreatedBy = Perm.CreatedBy

		createmod.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		createrolepermission = append(createrolepermission, createmod)

	}

	var err error

	if len(createrolepermission) != 0 {

		err = AS.CreateRolePermission(&createrolepermission, permission.DB)

	}

	return err

}

func (permission *PermissionConfig) PermissionList(tenantid string) (menu []MenuMod, err error) {

	if autherr := AuthandPermission(permission); autherr != nil {

		return []MenuMod{}, autherr
	}

	parentModules, err := AS.GetAllParentModule(permission.DB, tenantid) //get parent default main module only

	var Final []MenuMod

	for _, val := range parentModules {

		var fin MenuMod

		fin.Id = val.Id

		fin.ModuleName = val.ModuleName

		fin.IconPath = val.IconPath

		fin.TenantId = val.TenantId

		subModules, _ := AS.GetAllSubModule(val.Id, permission.DB, tenantid)

		var Suball []SubModule

		for _, tab := range subModules {

			if tab.ParentId == val.Id {

				var sub SubModule

				sub.Id = tab.Id

				sub.ModuleName = tab.ModuleName

				sub.IconPath = tab.IconPath

				sub.TenantId = tab.TenantId

				rout, _ := AS.GetModulePermissions(tab.Id, []int{}, permission.DB, tenantid)

				var Url []URL

				for _, url := range rout {

					var urll URL

					urll.Id = url.Id

					urll.DisplayName = url.DisplayName

					urll.RouteName = url.RouteName

					Url = append(Url, urll)

				}

				sub.Routes = Url

				Suball = append(Suball, sub)

			}

		}
		if len(Suball) > 0 && len(Suball[0].Routes) > 0 {

			fin.Route = Suball[0].Routes[0].RouteName
		}

		fin.SubModule = Suball

		Final = append(Final, fin)
	}

	return Final, nil

}

// permission List
func (permission *PermissionConfig) PermissionListRoleId(limit, offset, roleid int, filter Filter, tenantid string) (Module []Tblmodule, count int64, err error) {

	if autherr := AuthandPermission(permission); autherr != nil {

		return []Tblmodule{}, 0, autherr
	}

	var allmodules []Tblmodule

	var parentid []int //all parentid

	allmodule, err := AS.GetAllParentModules1(permission.DB, tenantid)

	for _, val := range allmodule {

		parentid = append(parentid, val.Id)
	}

	submod, err := AS.GetAllSubModules(parentid, permission.DB, tenantid)

	for _, val := range allmodule {

		if val.ModuleName == "Settings" {

			var newmod Tblmodule

			newmod.Id = val.Id

			newmod.Description = val.Description

			newmod.CreatedBy = val.CreatedBy

			newmod.ModuleName = val.ModuleName

			newmod.IsActive = val.IsActive

			newmod.IconPath = val.IconPath

			newmod.TenantId = val.TenantId

			newmod.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

			for _, sub := range submod {

				if sub.ParentId == val.Id {

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == sub.Id {

							var modper TblModulePermission

							modper.Id = getmod.Id

							modper.Description = getmod.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

							modper.TenantId = getmod.TenantId

							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")

							modper.FullAccessPermission = getmod.FullAccessPermission

							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}
				}

			}

			allmodules = append(allmodules, newmod)

		} else if val.ModuleName == "Spaces" {

			var newmod Tblmodule

			newmod.Id = val.Id

			newmod.Description = val.Description

			newmod.CreatedBy = val.CreatedBy

			newmod.ModuleName = val.ModuleName

			newmod.IsActive = val.IsActive

			newmod.IconPath = val.IconPath

			newmod.TenantId = val.TenantId

			newmod.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

			for _, sub := range submod {

				if sub.Id == val.Id {

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == val.Id {

							var modper TblModulePermission

							modper.Id = getmod.Id

							modper.Description = getmod.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")

							modper.TenantId = val.TenantId

							modper.FullAccessPermission = getmod.FullAccessPermission

							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}
				}

			}

			allmodules = append(allmodules, newmod)

		} else {

			for _, sub := range submod {

				if sub.ParentId == val.Id {

					var newmod Tblmodule

					newmod.Id = sub.Id

					newmod.Description = sub.Description

					newmod.CreatedBy = sub.CreatedBy

					newmod.ModuleName = sub.ModuleName

					newmod.IsActive = sub.IsActive

					newmod.IconPath = sub.IconPath

					newmod.TenantId = sub.TenantId

					newmod.CreatedDate = sub.CreatedOn.Format("02 Jan 2006 03:04 PM")

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == sub.Id {

							var modper TblModulePermission

							modper.Id = getmod.Id

							modper.Description = sub.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

							modper.TenantId = getmod.TenantId

							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")

							modper.FullAccessPermission = getmod.FullAccessPermission

							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}

					allmodules = append(allmodules, newmod)

				}
			}

		}

	}

	_, Totalcount := AS.GetAllModules(0, 0, roleid, filter, permission.DB, tenantid)

	return allmodules, Totalcount, nil

}

// permission List
func (permission *PermissionConfig) GetPermissionDetailsById(roleid int, tenantid string) (rolepermissionid []int, err error) {

	if autherr := AuthandPermission(permission); autherr != nil {

		return []int{}, autherr
	}

	var permissionid []int

	roleper, err := AS.GetPermissionId(roleid, permission.DB, tenantid)

	for _, val := range roleper {

		permissionid = append(permissionid, val.PermissionId)

	}

	return permissionid, nil

}

// Set Difference: A - B
func Difference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
