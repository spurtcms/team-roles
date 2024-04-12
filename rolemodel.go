package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type RoleCreation struct {
	Name        string
	Description string
	CreatedBy   int
}

type tblrole struct {
	Id          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Slug        string    `gorm:"column:slug"`
	IsActive    int       `gorm:"column:is_active"`
	IsDeleted   int       `gorm:"column:is_deleted"`
	CreatedOn   time.Time `gorm:"column:created_on"`
	CreatedBy   int       `gorm:"column:created_by"`
	ModifiedOn  time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"column:modified_by;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
	User        []tbluser `gorm:"-"`
}

type tblroleuser struct {
	Id           int       `gorm:"column:id"`
	RoleId       int       `gorm:"column:role_id"`
	UserId       int       `gorm:"column:user_id"`
	CreatedBy    int       `gorm:"column:created_by"`
	CreatedOn    time.Time `gorm:"column:created_on"`
	ModifiedBy   int       `gorm:"column:modified_by;DEFAULT:NULL"`
	ModifiedOn   time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModuleName   string    `gorm:"-"`
	RouteName    string    `gorm:"-:migration;<-:false"`
	DisplayName  string    `gorm:"-:migration;<-"`
	Description  string    `gorm:"-"`
	ModuleId     int       `gorm:"-:migration;<-"`
	PermissionId int       `gorm:"-"`
}

type rolelist struct {
	Limit      int
	Offset     int
	filter     Filter
	GetAllData bool
}

type ModelStruct struct{} //Just Group the all model using this struct

// Get all roles list with limit and offset
func (as ModelStruct) GetAllRoles(limit, offset int, filter Filter, getalldata bool, DB *gorm.DB) (role []tblrole, rolecount int64, err error) {

	query := DB.Table("tbl_roles").Where("is_deleted = 0").Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if getalldata {

		query.Find(&role)

	}

	if limit != 0 && !getalldata {

		query.Limit(limit).Offset(offset).Find(&role)

	} else {

		query.Find(&role).Count(&rolecount)

		return role, rolecount, nil
	}

	return []tblrole{}, 0, nil
}

/*Get role by id*/
func (as ModelStruct) GetRoleById(id int, DB *gorm.DB) (role tblrole, err error) {

	if err := DB.Model(tblrole{}).Where("id=?", id).First(&role).Error; err != nil {

		return tblrole{}, err

	}

	return role, nil
}

// Roels Insert
func (as ModelStruct) RoleCreate(role *tblrole, DB *gorm.DB) error {

	if err := DB.Model(tblrole{}).Create(role).Error; err != nil {

		return err
	}

	return nil
}

/**/
func (as ModelStruct) RoleUpdate(role *tblrole, DB *gorm.DB) error {

	if err := DB.Model(tblrole{}).Where("id=?", role.Id).Updates(tblrole{Name: role.Name, Description: role.Description, Slug: role.Slug, IsActive: role.IsActive, IsDeleted: role.IsDeleted, ModifiedOn: role.ModifiedOn, ModifiedBy: role.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

// Delete the role data
func (as ModelStruct) RoleDelete(role *tblrole, id int, DB *gorm.DB) error {

	if err := DB.Model(tblrole{}).Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {

		return err

	}

	return nil
}

/*Check role*/
func (as ModelStruct) CheckRoleExists(role *tblrole, id int, name string, DB *gorm.DB) error {

	if id == 0 {
		if err := DB.Model(tblrole{}).Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and is_deleted = 0 ", name).First(&role).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Model(tblrole{}).Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and id not in(?) and is_deleted= 0 ", name, id).First(&role).Error; err != nil {

			return err
		}
	}
	return nil

}

func (as ModelStruct) GetRolesData(roles *[]tblrole, DB *gorm.DB) error {

	if err := DB.Where("is_deleted=? and is_active=1", 0).Order("name").Find(&roles).Error; err != nil {

		return err

	}

	return nil
}
