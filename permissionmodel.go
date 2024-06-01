package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type Tblmodule struct {
	Id                  int                   `gorm:"column:id"`
	ModuleName          string                `gorm:"column:module_name"`
	IsActive            int                   `gorm:"column:is_active"`
	CreatedBy           int                   `gorm:"column:created_by"`
	CreatedOn           time.Time             `gorm:"column:created_on;DEFAULT:NULL"`
	CreatedDate         string                `gorm:"-:migration;<-:false"`
	DefaultModule       int                   `gorm:"column:default_module"`
	ParentId            int                   `gorm:"column:parent_id"`
	IconPath            string                `gorm:"column:icon_path"`
	TblModulePermission []TblModulePermission `gorm:"-:migration;<-:false; foreignKey:ModuleId"`
	Description         string                `gorm:"column:description"`
	OrderIndex          int                   `gorm:"column:order_index"`
}

type TblModule struct {
	Id               int       `gorm:"column:id"`
	ModuleName       string    `gorm:"column:module_name"`
	IsActive         int       `gorm:"column:is_active"`
	DefaultModule    int       `gorm:"column:default_module"`
	ParentId         int       `gorm:"column:parent_id"`
	IconPath         string    `gorm:"column:icon_path"`
	AssignPermission int       `gorm:"column:assign_permission"`
	Description      string    `gorm:"column:description"`
	OrderIndex       int       `gorm:"column:order_index"`
	CreatedBy        int       `gorm:"column:created_by"`
	CreatedOn        time.Time `gorm:"column:created_on"`
	MenuType         string    `gorm:"column:menu_type"`
	GroupFlg         int       `gorm:"column:group_flg"`
}

// type TblRolePermission struct {
// 	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
// 	RoleId       int       `gorm:"type:integer"`
// 	PermissionId int       `gorm:"type:integer"`
// 	CreatedBy    int       `gorm:"type:integer"`
// 	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
// }

type TblModulePermission struct {
	Id                   int                 `gorm:"column:id"`
	RouteName            string              `gorm:"column:route_name"`
	DisplayName          string              `gorm:"column:display_name"`
	SlugName             string              `gorm:"column:slug_name"`
	Description          string              `gorm:"column:description"`
	ModuleId             int                 `gorm:"column:module_id"`
	CreatedBy            int                 `gorm:"column:created_by"`
	CreatedOn            time.Time           `gorm:"column:created_on;DEFAULT:NULL"`
	CreatedDate          string              `gorm:"-:all"`
	ModifiedBy           int                 `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time           `gorm:"column:modified_by;DEFAULT:NULL"`
	ModuleName           string              `gorm:"-:migration;<-:false"`
	FullAccessPermission int                 `gorm:"column:full_access_permission"`
	ParentId             int                 `gorm:"column:parent_id"`
	AssignPermission     int                 `gorm:"column:assign_permission"`
	BreadcrumbName       string              `gorm:"column:breadcrumb_name"`
	TblRolePermission    []TblRolePermission `gorm:"-:migration;<-:false; foreignKey:PermissionId"`
	OrderIndex           int                 `gorm:"column:order_index"`
}

type TblRolePermission struct {
	Id           int       `gorm:"column:id"`
	RoleId       int       `gorm:"column:role_id"`
	PermissionId int       `gorm:"column:permission_id"`
	CreatedBy    int       `gorm:"column:created_by"`
	CreatedOn    time.Time `gorm:"column:created_on;DEFAULT:NULL"`
	CreatedDate  string    `gorm:"-:migration;<-:false"`
}

type tbluser struct {
	Id                   int `gorm:"primaryKey;auto_increment"`
	Uuid                 string
	FirstName            string
	LastName             string
	RoleId               int
	Email                string
	Username             string
	Password             string
	MobileNo             string
	IsActive             int
	ProfileImage         string
	ProfileImagePath     string
	DataAccess           int
	CreatedOn            time.Time
	CreatedBy            int
	ModifiedOn           time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	LastLogin            time.Time `gorm:"DEFAULT:NULL"`
	IsDeleted            int
	DeletedOn            time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy            int       `gorm:"DEFAULT:NULL"`
	ModuleName           string    `gorm:"-"`
	RouteName            string    `gorm:"-:migration;<-:false"`
	DisplayName          string    `gorm:"-:migration;<-:false"`
	Description          string    `gorm:"-"`
	ModuleId             int       `gorm:"-:migration;<-:false"`
	PermissionId         int       `gorm:"-"`
	FullAccessPermission int       `gorm:"-:migration;<-:false"`
	Roles                []TblRole `gorm:"-"`
}

type Filter struct {
	Keyword  string
	Category string
	Status   string
	FromDate string
	ToDate   string
}

type Permission struct {
	ModuleName string
	Action     []string //create,View,update,delete

}

type MultiPermissin struct {
	RoleId    int
	Ids       []int
	CreatedBy int
	// Permissions []Permission
}

type Action string

type CreatePermissions struct {
	RoleId     int
	ModuleName string
	Permission Action
	CreatedBy  int
}

type LoginCheck struct {
	Username string
	Password string
}

type SubModule struct {
	Id         int
	ModuleName string
	IconPath   string
	Routes     []URL
	// FullAccessPermission bool
	Action bool
}

type URL struct {
	Id          int
	DisplayName string
	RouteName   string
}

type MenuMod struct {
	Id         int
	ModuleName string
	IconPath   string
	Routes     []URL // this is for single menu multiple permissions arr
	HrefRoute  URL
	Route      string      //this is for a tag href route
	SubModule  []SubModule // this is for submodules
	// FullAccessPermission bool
	EmptyCheck bool //this is flg for mainmenu hide if submodule is empty
}

/*bulk creation*/
func (as ModelStruct) CreateRolePermission(roleper *[]TblRolePermission, DB *gorm.DB) error {

	if err := DB.Table("tbl_role_permissions").Create(&roleper).Error; err != nil {

		return err

	}

	return nil
}

func (as ModelStruct) CheckPermissionIdNotExist(roleid int, permissionid []int, DB *gorm.DB) (roleperm []TblRolePermission, err error) {

	if err := DB.Table("tbl_role_permissions").Where("role_id=? and permission_id not in(?)", roleid, permissionid).Find(&roleperm).Error; err != nil {

		return roleperm, err

	}
	return roleperm, nil
}

/*Delete Role Permission by id*/
func (as ModelStruct) Deleterolepermission(id int, DB *gorm.DB) (TblRolePermission TblRolePermission, err error) {

	if err := DB.Where("permission_id=?", id).Delete(&TblRolePermission).Error; err != nil {

		return TblRolePermission, err
	}

	return TblRolePermission, nil
}

func (as ModelStruct) DeleteRolePermissionById(roleid int, DB *gorm.DB) (roleper []TblRolePermission, err error) {

	if err := DB.Where("role_id=?", roleid).Delete(&roleper).Error; err != nil {

		return roleper, err

	}
	return roleper, nil
}

func (as ModelStruct) CheckPermissionIdExist(roleid int, permissionid []int, DB *gorm.DB) (roleperm []TblRolePermission, err error) {

	if err := DB.Table("tbl_role_permissions").Where("role_id=? and permission_id in(?)", roleid, permissionid).Find(&roleperm).Error; err != nil {

		return roleperm, err

	}
	return roleperm, nil
}

/**/
func (as ModelStruct) GetAllParentModules1(DB *gorm.DB) (mod []Tblmodule, err error) {

	if err := DB.Model("tbl_modules").Where("parent_id=0").Find(&mod).Error; err != nil {

		return mod, err
	}

	return mod, nil
}

/**/
func (as ModelStruct) GetAllSubModules(ids []int, DB *gorm.DB) (mod []Tblmodule, err error) {

	if err := DB.Model("tbl_modules").Where("(tbl_modules.parent_id in (?) or id in(?)) and tbl_modules.assign_permission=1", ids, ids).Order("order_index").Preload("TblModulePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("assign_permission =0").Order("order_index asc")
	}).Find(&mod).Error; err != nil {

		return mod, err
	}

	return mod, nil
}

/*This is for assign permission*/
func (as ModelStruct) GetAllModules(limit, offset, id int, filter Filter, DB *gorm.DB) (mod []Tblmodule, count int64) {

	query := DB.Table("tbl_modules").Where("parent_id!=0 or assign_permission=1").Preload("TblModulePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("assign_permission =0")
	}).Preload("TblModulePermission.TblRolePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id = ?", id)
	})

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(module_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("id asc").Find(&mod)

	} else {

		query.Find(&mod).Count(&count)

		return mod, count
	}

	return mod, 0
}

/*Get PermissionId By RoleId*/
func (as ModelStruct) GetPermissionId(roleid int, DB *gorm.DB) (perm []TblRolePermission, err error) {

	if err := DB.Table("tbl_role_permissions").Where("role_id=?", roleid).Find(&perm).Error; err != nil {

		return perm, err
	}

	return perm, nil
}

func (as ModelStruct) GetAllSubModule(moduleid int, DB *gorm.DB) (modules []TblModule, err error) {

	if err := DB.Table("tbl_modules").Where("parent_id = (?) and menu_type='tab'", moduleid).Order("tbl_modules.id asc").Find(&modules).
		Error; err != nil {

		return modules, err
	}

	return modules, nil
}

func (as ModelStruct) GetAllParentModule(DB *gorm.DB) (modules []TblModule, err error) {

	if err := DB.Table("tbl_modules").Where("default_module = 0 and parent_id = 0").Order("tbl_modules.id asc").Find(&modules).
		Error; err != nil {

		return modules, err
	}

	return modules, nil
}

func (as ModelStruct) GetModulePermissions(modid int, ids []int, DB *gorm.DB) (permission []TblModulePermission, err error) {

	query := DB.Table("tbl_module_permissions").Select("tbl_module_permissions.*,tbl_modules.module_name").Joins("inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id").Order("tbl_modules.order_index asc,tbl_module_permissions.order_index asc")

	if len(ids) > 0 {

		query = query.Where("tbl_module_permissions.id in (?)", ids)

	}

	if modid != 0 && modid > -1 {

		query = query.Where("module_id = (?)", modid)
	}

	query.Find(&permission)

	if err := query.Error; err != nil {

		return permission, err
	}

	return permission, nil
}

func (as ModelStruct) CheckModuleExists(modulename string, DB *gorm.DB) (tblmod Tblmodule, err error) {

	if qerr := DB.Model("tbl_modules").Where("module_name =? and parent_id != 0 ").First(tblmod).Error; err != nil {

		return Tblmodule{}, qerr
	}

	return tblmod, nil

}

func (as ModelStruct) CheckModulePemissionExists(moduleid int, permissions Action, DB *gorm.DB) (tblmod TblModulePermission, err error) {

	if permissions == "CRUD" {

		if qerr := DB.Model("tbl_modules").Where("module_id =? and full_access_permission= 1  ", moduleid).First(tblmod).Error; qerr != nil {

			return TblModulePermission{}, qerr
		}

	} else {

		if qerr := DB.Model("tbl_modules").Where("module_id =? and display_name = ?  ", moduleid, permissions).First(tblmod).Error; qerr != nil {

			return TblModulePermission{}, qerr
		}

	}

	return tblmod, nil

}

/*Get Id by RouteName*/
func (as ModelStruct) GetIdByRouteName(id string, DB *gorm.DB) (tblmodper TblModulePermission, err error) {

	if err := DB.Table("tbl_module_permissions").Where("route_name=?", "/channel/entrylist/"+id).First(&tblmodper).Error; err != nil {

		return tblmodper, err
	}

	return tblmodper, nil
}

func (as ModelStruct) DeleteModulePermissioninEntries(id string, DB *gorm.DB) (tblmodper TblModulePermission, err error) {

	if err := DB.Where("route_name=?", "/channel/entrylist/"+id).Delete(&tblmodper).Error; err != nil {

		return tblmodper, err
	}

	return tblmodper, nil
}

/*create json module permission*/
func (as ModelStruct) CreateModulePermission(modpermission *TblModulePermission, DB *gorm.DB) (modper *TblModulePermission, err error) {

	if err := DB.Model(TblModulePermission{}).Create(&modpermission).Error; err != nil {

		return &TblModulePermission{}, err

	}

	return modpermission, nil
}
