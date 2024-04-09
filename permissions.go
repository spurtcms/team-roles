package teamroles

import (
	"log"
	"time"
)

type Action string

const ( //for permission check
	Create Action = "Create"

	Read Action = "View"

	Update Action = "Update"

	Delete Action = "Delete"

	CRUD Action = "CRUD"
)

// Check User Permission
func (permission PermissionConfig) IsGranted(modulename string, permisison Action) (bool, error) {

	if permission.RoleId != 1 || permission.RoleName != "Super Admin" { //if not an admin user

		var modid int

		var module TblModule

		var modpermissions TblModulePermission

		if err := permission.DB.Model(TblModule{}).Where("module_name=? and parent_id !=0", modulename).Find(&module).Error; err != nil {

			return false, err
		}

		if err1 := permission.DB.Model(TblModulePermission{}).Where("display_name=?", modulename).Find(&modpermissions).Error; err1 != nil {

			return false, err1
		}

		if module.Id != 0 {

			modid = module.Id

		} else {

			modid = modpermissions.Id
		}

		var modulepermission []TblModulePermission

		if permisison == "CRUD" {

			if err := permission.DB.Model(TblModulePermission{}).Where("id=? and (full_access_permission=1 or display_name='View' or display_name='Update' or  display_name='Create' or display_name='Delete')", modid).Find(&modulepermission).Error; err != nil {

				return false, err
			}

		} else {

			if err := permission.DB.Model(TblModulePermission{}).Where("module_id=? and display_name=?", modid, permisison).Find(&modulepermission).Error; err != nil {

				return false, err
			}

		}

		for _, val := range modulepermission {

			var rolecheck TblRolePermission

			if err := permission.DB.Model(TblRolePermission{}).Where("permission_id=? and role_id=?", val.Id, permission.RoleId).First(&rolecheck).Error; err != nil {

				return false, err
			}

		}

		permission.PermissionFlg = true

	}

	permission.PermissionFlg = true

	return true, nil

}

// create permission
func (permission PermissionConfig) CreatePermission(Perm MultiPermissin) error {

	if permission.AuthEnable && !permission.Authenticate.AuthFlg {

		return ErrorAuth
	}

	if permission.PermissionEnable && !permission.Authenticate.PermissionFlg {

		return ErrorPermission
	}

	var createrolepermission []tblrolepermission

	for _, roleperm := range Perm.Ids {

		var createmod tblrolepermission

		createmod.PermissionId = roleperm

		createmod.RoleId = Perm.RoleId

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

// update permission
func (permission PermissionConfig) CreateUpdatePermission(Perm MultiPermissin) error {

	if permission.AuthEnable && !permission.Authenticate.AuthFlg {

		return ErrorAuth
	}

	if permission.PermissionEnable && !permission.Authenticate.PermissionFlg {

		return ErrorPermission
	}

	var checknotexist []tblrolepermission

	cnerr := AS.CheckPermissionIdNotExist(&checknotexist, Perm.RoleId, Perm.Ids, permission.DB)

	if len(Perm.Ids) == 0 {

		AS.Deleterolepermission(&tblrolepermission{}, Perm.RoleId, permission.DB)
	}

	if cnerr != nil {

		log.Println(cnerr)

	} else if len(checknotexist) != 0 {

		AS.DeleteRolePermissionById(&checknotexist, Perm.RoleId, permission.DB)
	}

	var checkexist []TblRolePermission

	cerr := AS.CheckPermissionIdExist(&checkexist, Perm.RoleId, Perm.Ids, permission.DB)

	if cerr != nil {

		log.Println(cerr)

	}

	var existid []int

	for _, exist := range checkexist {

		existid = append(existid, exist.PermissionId)

	}

	pid := Difference(Perm.Ids, existid)

	var createrolepermission []tblrolepermission

	for _, roleperm := range pid {

		var createmod tblrolepermission

		createmod.PermissionId = roleperm

		createmod.RoleId = Perm.RoleId

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

// permission List
func (permission PermissionConfig) PermissionListRoleId(limit, offset, roleid int, filter Filter) (Module []tblmodule, count int64, err error) {

	if permission.AuthEnable && !permission.Authenticate.AuthFlg {

		return []tblmodule{}, 0, ErrorAuth
	}

	if permission.PermissionEnable && !permission.Authenticate.PermissionFlg {

		return []tblmodule{}, 0, ErrorPermission
	}

	var allmodule []tblmodule

	var allmodules []tblmodule

	var parentid []int //all parentid

	AS.GetAllParentModules1(&allmodule, permission.DB)

	for _, val := range allmodule {

		parentid = append(parentid, val.Id)
	}

	var submod []tblmodule

	AS.GetAllSubModules(&submod, parentid, permission.DB)

	for _, val := range allmodule {

		if val.ModuleName == "Settings" {

			var newmod tblmodule

			newmod.Id = val.Id

			newmod.Description = val.Description

			newmod.CreatedBy = val.CreatedBy

			newmod.ModuleName = val.ModuleName

			newmod.IsActive = val.IsActive

			newmod.IconPath = val.IconPath

			newmod.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

			for _, sub := range submod {

				if sub.ParentId == val.Id {

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == sub.Id {

							var modper tblmodulepermission

							modper.Id = getmod.Id

							modper.Description = getmod.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")

							modper.FullAccessPermission = getmod.FullAccessPermission

							newmod.TblModulePermission = append(newmod.TblModulePermission, modper)
						}

					}
				}

			}

			allmodules = append(allmodules, newmod)

		} else if val.ModuleName == "Spaces" {

			var newmod tblmodule

			newmod.Id = val.Id

			newmod.Description = val.Description

			newmod.CreatedBy = val.CreatedBy

			newmod.ModuleName = val.ModuleName

			newmod.IsActive = val.IsActive

			newmod.IconPath = val.IconPath

			newmod.CreatedDate = val.CreatedOn.Format("02 Jan 2006 03:04 PM")

			for _, sub := range submod {

				if sub.Id == val.Id {

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == val.Id {

							var modper tblmodulepermission

							modper.Id = getmod.Id

							modper.Description = getmod.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

							modper.CreatedDate = val.CreatedOn.Format("2006-01-02 15:04:05")

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

					var newmod tblmodule

					newmod.Id = sub.Id

					newmod.Description = sub.Description

					newmod.CreatedBy = sub.CreatedBy

					newmod.ModuleName = sub.ModuleName

					newmod.IsActive = sub.IsActive

					newmod.IconPath = sub.IconPath

					newmod.CreatedDate = sub.CreatedOn.Format("02 Jan 2006 03:04 PM")

					for _, getmod := range sub.TblModulePermission {

						if getmod.ModuleId == sub.Id {

							var modper tblmodulepermission

							modper.Id = getmod.Id

							modper.Description = sub.Description

							modper.DisplayName = getmod.DisplayName

							modper.ModuleName = getmod.ModuleName

							modper.RouteName = getmod.RouteName

							modper.CreatedBy = getmod.CreatedBy

							modper.Description = getmod.Description

							modper.TblRolePermission = getmod.TblRolePermission

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

	var allmodul []tblmodule

	Totalcount := AS.GetAllModules(&allmodul, 0, 0, roleid, filter, permission.DB)

	return allmodules, Totalcount, nil

}

// permission List
func (permission PermissionConfig) GetPermissionDetailsById(roleid int) (rolepermissionid []int, err error) {

	if permission.AuthEnable && !permission.Authenticate.AuthFlg {

		return []int{}, ErrorAuth
	}

	if permission.PermissionEnable && !permission.Authenticate.PermissionFlg {

		return []int{}, ErrorPermission
	}

	var permissionid []int

	var roleper []tblrolepermission

	AS.GetPermissionId(&roleper, roleid, permission.DB)

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
