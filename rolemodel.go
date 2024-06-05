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

type Tblrole struct {
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

type TblRole struct {
	Id          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Slug        string    `gorm:"column:slug"`
	IsActive    int       `gorm:"column:is_active"`
	IsDeleted   int       `gorm:"column:is_deleted"`
	CreatedOn   time.Time `gorm:"column:created_on"`
	CreatedBy   int       `gorm:"column:created_by"`
	ModifiedOn  time.Time `gorm:"column:modified_on"`
	ModifiedBy  int       `gorm:"column:modified_by"`
}

type Rolelist struct {
	Limit      int
	Offset     int
	Filter     Filter
	GetAllData bool
}

// Just Group the all model using this struct
type ModelStruct struct {
	UserId     int
	DataAccess int
}

// Get all roles list with limit and offset
func (as ModelStruct) GetAllRoles(limit, offset int, filter Filter, getalldata bool, DB *gorm.DB) (role []Tblrole, rolecount int64, err error) {

	query := DB.Table("tbl_roles").Where("is_deleted = 0").Order("id desc")

	if as.DataAccess == 1 {

		query = query.Where("tbl_roles.created_by = ?", as.UserId)

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if getalldata {

		query.Find(&role)

	}

	if limit != 0 && !getalldata {

		query.Limit(limit).Offset(offset).Find(&role)

		return role, rolecount, nil

	}

	query.Find(&role).Count(&rolecount)

	return role, rolecount, nil

}

/*Get role by id*/
func (as ModelStruct) GetRoleById(id int, DB *gorm.DB) (role Tblrole, err error) {

	if err := DB.Table("tbl_roles").Where("id=?", id).First(&role).Error; err != nil {

		return Tblrole{}, err

	}

	return role, nil
}

// Roels Insert
func (as ModelStruct) RoleCreate(role *Tblrole, DB *gorm.DB) error {

	if err := DB.Table("tbl_roles").Create(role).Error; err != nil {

		return err
	}

	return nil
}

/**/
func (as ModelStruct) RoleUpdate(role *Tblrole, DB *gorm.DB) error {

	if err := DB.Table("tbl_roles").Where("id=?", role.Id).Updates(Tblrole{Name: role.Name, Description: role.Description, Slug: role.Slug, IsActive: role.IsActive, IsDeleted: role.IsDeleted, ModifiedOn: role.ModifiedOn, ModifiedBy: role.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

// Delete the role data
func (as ModelStruct) RoleDelete(id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_roles").Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {

		return err

	}

	return nil
}

/*Check role*/
func (as ModelStruct) CheckRoleExists(role *TblRole, id int, name string, DB *gorm.DB) error {

	if id == 0 {
		if err := DB.Table("tbl_roles").Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and is_deleted = 0 ", name).First(&role).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_roles").Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and id not in(?) and is_deleted= 0 ", name, id).First(&role).Error; err != nil {

			return err
		}
	}
	return nil

}

func (as ModelStruct) GetRolesData(roles *[]Tblrole, DB *gorm.DB) error {

	if err := DB.Where("is_deleted=? and is_active=1", 0).Order("name").Find(&roles).Error; err != nil {

		return err

	}

	return nil
}

// delete multiple roles
func (as ModelStruct) MultiSelectRoleDelete(role *TblRole, ids []int, id int, DB *gorm.DB) error {

	if id != 0 {

		if err := DB.Table("tbl_roles").Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {

			return err

		}

	} else {

		if err := DB.Table("tbl_roles").Where("id in (?)", ids).Update("is_deleted", 1).Error; err != nil {

			return err

		}

	}
	return nil

}

// delete multiple permission
func (as ModelStruct) MultiSelectDeleteRolePermissionById(roleper *[]TblRolePermission, roleids []int, roleid int, DB *gorm.DB) error {

	if roleid != 0 {

		if err := DB.Where("role_id=?", roleid).Delete(&roleper).Error; err != nil {

			return err

		}
		return nil

	} else {

		if err := DB.Where("role_id in (?)", roleids).Delete(&roleper).Error; err != nil {

			return err

		}
	}
	return nil
}

// update selected role status
func (as ModelStruct) MultiSelectRoleIsActive(role *TblRole, id []int, val int, DB *gorm.DB) error {

	if err := DB.Table("tbl_roles").Where("id in (?)", id).UpdateColumns(map[string]interface{}{"is_active": val, "modified_by": role.ModifiedBy, "modified_on": role.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*update role status*/
func (as ModelStruct) RoleIsActive(role *TblRole, id, val int, DB *gorm.DB) error {

	if err := DB.Table("tbl_roles").Where("id=?", id).UpdateColumns(map[string]interface{}{"is_active": val, "modified_by": role.ModifiedBy, "modified_on": role.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
