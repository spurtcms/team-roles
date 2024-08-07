package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:integer"`
	TenantId    int       `gorm:"type:integer"`
}

type TblModule struct {
	Id               int       `gorm:"primaryKey;type:serial"`
	ModuleName       string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	DefaultModule    int       `gorm:"type:integer"`
	ParentId         int       `gorm:"type:integer"`
	IconPath         string    `gorm:"type:character varying"`
	AssignPermission int       `gorm:"type:integer"`
	Description      string    `gorm:"type:character varying"`
	OrderIndex       int       `gorm:"type:integer"`
	CreatedBy        int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	MenuType         string    `gorm:"type:character varying"`
	GroupFlg         int       `gorm:"type:integer"`
	TenantId         int       `gorm:"type:integer"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey;type:serial"`
	RouteName            string    `gorm:"type:character varying;unique"`
	DisplayName          string    `gorm:"type:character varying"`
	SlugName             string    `gorm:"type:character varying"`
	Description          string    `gorm:"type:character varying"`
	ModuleId             int       `gorm:"type:integer"`
	FullAccessPermission int       `gorm:"type:integer"`
	ParentId             int       `gorm:"type:integer"`
	AssignPermission     int       `gorm:"type:integer"`
	BreadcrumbName       string    `gorm:"type:character varying"`
	OrderIndex           int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId             int       `gorm:"type:integer"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	TenantId     int       `gorm:"type:integer"`
}

type TblRoleUser struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId     int       `gorm:"type:integer"`
	UserId     int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId   int       `gorm:"type:integer"`
}

func MigrationTables(db *gorm.DB) {

	err := db.AutoMigrate(
		&TblRole{},
		&TblModule{},
		&TblModulePermission{},
		&TblRolePermission{},
		&TblRoleUser{},
	)

	if err != nil {

		panic(err)
	}
}