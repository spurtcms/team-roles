package teamroles

import (
	"fmt"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// test permission working or not
func TestIsGranted(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		DB:           &gorm.DB{},
		Authenticate: &auth.Auth{},
		RoleId:       1,
	})

	flg, err := PermissionFlg.IsGranted("spaces", CRUD)

	fmt.Println(err, flg)

}

// test create permissions
func TestCreatePermissions(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		DB:           &gorm.DB{},
		Authenticate: &auth.Auth{},
	})

	permis := MultiPermissin{
		RoleId:    2,
		Ids:       []int{1, 2}, //permission_module primary key
		CreatedBy: 1,
	}

	err := PermissionFlg.CreatePermission(permis)

	fmt.Println(err)

}

//test update permissions
func TestUpdatePermissions(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		DB:           &gorm.DB{},
		Authenticate: &auth.Auth{},
	})

	permis := MultiPermissin{
		RoleId:    2,
		Ids:       []int{1, 2}, //permission module primary key
		CreatedBy: 1,
	}

	err := PermissionFlg.CreateUpdatePermission(permis)

	fmt.Println(err)

}

// test permissions list
func TestPermissionsList(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		DB:           &gorm.DB{},
		Authenticate: &auth.Auth{},
	})

	module, count, err := PermissionFlg.PermissionListRoleId(100, 0, 1, Filter{})

	if err != nil {

		panic(err)
	}

	fmt.Println(module, count)

}

// test permissions list
func TestPermissionsRole(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		DB:           &gorm.DB{},
		Authenticate: &auth.Auth{},
	})

	module, err := PermissionFlg.GetPermissionDetailsById(1)

	if err != nil {

		panic(err)
	}

	fmt.Println(module)

}
