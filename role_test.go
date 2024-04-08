package teamroles

import (
	"fmt"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

// test role list
func TestRoleList(t *testing.T) {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: false,
		DB:               &gorm.DB{},
		Authenticate:     Auth,
	})

	rolelist, count, err := PermissionFlg.RoleList(rolelist{Limit: 1, Offset: 0})

	if err != nil {

		panic(err)
	}

	fmt.Println(rolelist, count)

}

// test create role
func TestCreateRole(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       false,
		PermissionEnable: false,
		DB:               &gorm.DB{},
		Authenticate:     &auth.Auth{},
	})

	rolecreate, err := PermissionFlg.CreateRole(RoleCreation{Name: "Manager", Description: "", CreatedBy: 1})

	if err != nil {

		panic(err)
	}

	fmt.Println(rolecreate)

}

// test update role
func TestUpdateRole(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       false,
		PermissionEnable: false,
		DB:               &gorm.DB{},
		Authenticate:     &auth.Auth{},
	})

	roleupdate, err := PermissionFlg.UpdateRole(RoleCreation{Name: "Manager", Description: "", CreatedBy: 1}, 1)

	if err != nil {

		panic(err)
	}

	fmt.Println(roleupdate)

}

// test create role
func TestDeleteRole(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       false,
		PermissionEnable: false,
		DB:               &gorm.DB{},
		Authenticate:     &auth.Auth{},
	})

	flg, err := PermissionFlg.DeleteRole(1)

	if err != nil {

		panic(err)
	}

	fmt.Println(flg)

}

// test create role
func TestCheckRole(t *testing.T) {

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       false,
		PermissionEnable: false,
		DB:               &gorm.DB{},
		Authenticate:     &auth.Auth{},
	})

	flg, err := PermissionFlg.CheckRoleAlreadyExists(1, "Manager")

	if err != nil {

		panic(err)
	}

	fmt.Println(flg)

}
