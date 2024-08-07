package teamroles

import (
	"time"

	"gorm.io/gorm"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId    int       `gorm:"type:int"`
}

type TblModule struct {
	Id                   int       `gorm:"primaryKey"`
	ModuleName           string    `gorm:"type:varchar(255)"`
	IsActive             int       `gorm:"type:int"`
	DefaultModule        int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	IconPath             string    `gorm:"type:varchar(255)"`
	AssignPermission     int       `gorm:"type:int"`
	Description          string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	MenuType             string    `gorm:"type:varchar(255)"`
	FullAccessPermission int       `gorm:"type:int"`
	GroupFlg             int       `gorm:"type:int"`
	TenantId             int       `gorm:"type:int"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey"`
	RouteName            string    `gorm:"type:varchar(255)"`
	DisplayName          string    `gorm:"type:varchar(255)"`
	SlugName             string    `gorm:"type:varchar(255)"`
	Description          string    `gorm:"type:varchar(255)"`
	ModuleId             int       `gorm:"type:int"`
	FullAccessPermission int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	AssignPermission     int       `gorm:"type:int"`
	BreadcrumbName       string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn           time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId             int       `gorm:"type:int"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId     int       `gorm:"type:int"`
}

type TblRoleUser struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	RoleId     int       `gorm:"type:int"`
	UserId     int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId   int       `gorm:"type:int"`
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
