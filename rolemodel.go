package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type RoleCreation struct {
	Name        string
	Description string
	CreatedBy   int
	TenantId    string
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
	TenantId    string    `gorm:"column:tenant_id;"`
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
	TenantId     string    `gorm:"column:tenant_id;"`
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
	TenantId    string    `gorm:"column:tenant_id;"`
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
func (as ModelStruct) GetAllRoles(limit, offset int, filter Filter, getalldata bool, DB *gorm.DB, tenantid string, active bool) (role []Tblrole, rolecount int64, err error) {

	query := DB.Table("tbl_roles").Where("is_deleted = 0 and (slug = ? or tenant_id = ?)", "admin", tenantid).Order("id desc")

	if active {
		query = query.Where("is_active = 1")
	}

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

func (as ModelStruct) RoleId(id int, DB *gorm.DB) (user int, err error) {
	if err := DB.Debug().Table("tbl_users").Where("id=?", id).Select("role_id").Scan(&user).Error; err != nil {
		return 0, err
	}
	return user, nil
}

/*Get role by id*/
func (as ModelStruct) GetRoleById(id int, DB *gorm.DB, tenantid string) (role Tblrole, err error) {

	query := DB.Debug().Table("tbl_roles").Where("id = ?", id)

	if tenantid != "" {

		query = query.Where("tenant_id = ? ", tenantid)

	}

	//please don't remove this tenant_id is NULL from where condition
	if err := query.First(&role).Error; err != nil {

		return Tblrole{}, err

	}

	return role, nil
}

// Roels Insert
func (as ModelStruct) RoleCreate(role *Tblrole, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_roles").Create(role).Error; err != nil {

		return err
	}

	return nil
}

/**/
func (as ModelStruct) RoleUpdate(role *Tblrole, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_roles").Where("id=? and tenant_id = ?", role.Id, tenantid).Updates(Tblrole{Name: role.Name, Description: role.Description, Slug: role.Slug, IsActive: role.IsActive, IsDeleted: role.IsDeleted, ModifiedOn: role.ModifiedOn, ModifiedBy: role.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

// Delete the role data
func (as ModelStruct) RoleDelete(id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_roles").Where("id = ? and tenant_id = ?", id, tenantid).Update("is_deleted", 1).Error; err != nil {

		return err

	}

	return nil
}

/*Check role*/
func (as ModelStruct) CheckRoleExists(role *TblRole, id int, name string, DB *gorm.DB, tenantid string) error {

	if id == 0 {
		if err := DB.Debug().Table("tbl_roles").Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and is_deleted = 0 and (tenant_id = ? or tenant_id is null)", name, tenantid).First(&role).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Debug().Table("tbl_roles").Where("LOWER(TRIM(name))=LOWER(TRIM(?)) and id not in(?) and is_deleted= 0 and  (tenant_id = ? or tenant_id is null)", name, id, tenantid).First(&role).Error; err != nil {

			return err
		}
	}
	return nil

}

func (as ModelStruct) GetRolesData(roles *[]Tblrole, DB *gorm.DB, tenantid string) error {

	if err := DB.Where("is_deleted=? and is_active=1 and tenant_id = ?", 0, tenantid).Order("name").Find(&roles).Error; err != nil {

		return err

	}

	return nil
}

// delete multiple roles
func (as ModelStruct) MultiSelectRoleDelete(role *TblRole, ids []int, id int, DB *gorm.DB, tenantid string) error {

	if id != 0 {

		if err := DB.Table("tbl_roles").Where("id = ? and tenant_id = ?", id, tenantid).Update("is_deleted", 1).Error; err != nil {

			return err

		}

	} else {

		if err := DB.Table("tbl_roles").Where("id in (?) and tenant_id = ?", ids, tenantid).Update("is_deleted", 1).Error; err != nil {

			return err

		}

	}
	return nil

}

// delete multiple permission
func (as ModelStruct) MultiSelectDeleteRolePermissionById(roleper *[]TblRolePermission, roleids []int, roleid int, DB *gorm.DB, tenantid string) error {

	if roleid != 0 {

		if err := DB.Where("role_id=? and tenant_id = ?", roleid, tenantid).Delete(&roleper).Error; err != nil {

			return err

		}
		return nil

	} else {

		if err := DB.Where("role_id in (?) and tenant_id = ?", roleids, tenantid).Delete(&roleper).Error; err != nil {

			return err

		}
	}
	return nil
}

// update selected role status
func (as ModelStruct) MultiSelectRoleIsActive(role *TblRole, id []int, val int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_roles").Where("id in (?) and tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": val, "modified_by": role.ModifiedBy, "modified_on": role.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*update role status*/
func (as ModelStruct) RoleIsActive(role *TblRole, id, val int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_roles").Where("id=? and tenant_id = ?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": val, "modified_by": role.ModifiedBy, "modified_on": role.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (as ModelStruct) GetRoleByName(role *[]TblRole, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_roles").Where("slug IN (?) or tenant_id = ? ", []string{"admin", "super_admin"}, tenantid).Find(&role).Error; err != nil {
		return err
	}

	return nil
}

// get role by slugname
func (as ModelStruct) GetRoleBySlug(slug string, DB *gorm.DB, tenantid string) (role TblRole, err error) {

	if err := DB.Table("tbl_roles").Where("slug =?  and tenant_id = ? and is_deleted=0", slug, tenantid).Find(&role).Error; err != nil {
		return TblRole{}, err
	}

	return role, nil
}
func (as ModelStruct) Checkrolespermission(user *tbluser, id int, ids []int, DB *gorm.DB, tenantid string) error {
	query := DB.Table("tbl_users")
	if id != 0 {
		query = query.Where("role_id=? and tenant_id=? and is_deleted=0", id, tenantid)
	} else {
		query = query.Where("role_id in (?) and tenant_id=? and is_deleted=0", ids, tenantid)
	}
	if err := query.Debug().First(&user).Error; err != nil {
		return err
	}
	return nil
}
