package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type tblmodule struct {
	Id                  int `gorm:"primaryKey;auto_increment"`
	ModuleName          string
	IsActive            int
	CreatedBy           int
	CreatedOn           time.Time
	CreatedDate         string `gorm:"-:migration;<-:false"`
	DefaultModule       int
	ParentId            int
	IconPath            string
	TblModulePermission []tblmodulepermission `gorm:"-:migration;<-:false; foreignKey:ModuleId"`
	Description         string
	OrderIndex          int
}

type tblmodulepermission struct {
	Id                   int `gorm:"primaryKey;auto_increment"`
	RouteName            string
	DisplayName          string
	SlugName             string
	Description          string
	ModuleId             int
	CreatedBy            int
	CreatedOn            time.Time
	CreatedDate          string    `gorm:"-"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"DEFAULT:NULL"`
	ModuleName           string    `gorm:"-:migration;<-:false"`
	FullAccessPermission int
	ParentId             int
	AssignPermission     int
	BreadcrumbName       string
	TblRolePermission    []TblRolePermission `gorm:"-:migration;<-:false; foreignKey:PermissionId"`
	OrderIndex           int
}

type tblrolepermission struct {
	Id           int `gorm:"primaryKey;auto_increment"`
	RoleId       int
	PermissionId int
	CreatedBy    int
	CreatedOn    time.Time
	CreatedDate  string `gorm:"-:migration;<-:false"`
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

type LoginCheck struct {
	Username string
	Password string
}

/*bulk creation*/
func (as ModelStruct) CreateRolePermission(roleper *[]tblrolepermission, DB *gorm.DB) error {

	if err := DB.Table("tbl_role_permissions").Create(&roleper).Error; err != nil {

		return err

	}

	return nil
}

func (as ModelStruct) CheckPermissionIdNotExist(roleperm *[]tblrolepermission, roleid int, permissionid []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_role_permissions").Where("role_id=? and permission_id not in(?)", roleid, permissionid).Find(&roleperm).Error; err != nil {

		return err

	}
	return nil
}

/*Delete Role Permission by id*/
func (as ModelStruct) Deleterolepermission(TblRolePermission *tblrolepermission, id int, DB *gorm.DB) error {

	if err := DB.Where("permission_id=?", id).Delete(&TblRolePermission).Error; err != nil {

		return err
	}

	return nil
}

func (as ModelStruct) DeleteRolePermissionById(roleper *[]tblrolepermission, roleid int, DB *gorm.DB) error {

	if err := DB.Where("role_id=?", roleid).Delete(&roleper).Error; err != nil {

		return err

	}
	return nil
}

func (as ModelStruct) CheckPermissionIdExist(roleperm *[]TblRolePermission, roleid int, permissionid []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_role_permissions").Where("role_id=? and permission_id in(?)", roleid, permissionid).Find(&roleperm).Error; err != nil {

		return err

	}
	return nil
}

/**/
func (as ModelStruct) GetAllParentModules1(mod *[]tblmodule, DB *gorm.DB) (err error) {

	if err := DB.Model(TblModule{}).Where("parent_id=0").Find(&mod).Error; err != nil {

		return err
	}

	return nil
}

/**/
func (as ModelStruct) GetAllSubModules(mod *[]tblmodule, ids []int, DB *gorm.DB) (err error) {

	if err := DB.Model(TblModule{}).Where("(tbl_modules.parent_id in (?) or id in(?)) and tbl_modules.assign_permission=1", ids, ids).Order("order_index").Preload("TblModulePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("assign_permission =0").Order("order_index asc")
	}).Find(&mod).Error; err != nil {

		return err
	}

	return nil
}

/*This is for assign permission*/
func (as ModelStruct) GetAllModules(mod *[]tblmodule, limit, offset, id int, filter Filter, DB *gorm.DB) (count int64) {

	query := DB.Table("tbl_modules").Where("parent_id!=0 or assign_permission=1").Preload("TblModulePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("assign_permission =0")
	}).Preload("TblModulePermission.TblRolePermission", func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id = ?", id)
	})

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(module_name)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("id asc").Find(&mod)

	} else {

		query.Find(&mod).Count(&count)

		return count
	}

	return 0
}

/*Get PermissionId By RoleId*/
func (as ModelStruct) GetPermissionId(perm *[]tblrolepermission, roleid int, DB *gorm.DB) error {

	if err := DB.Model(tblrolepermission{}).Where("role_id=?", roleid).Find(&perm).Error; err != nil {

		return err
	}

	return nil
}
