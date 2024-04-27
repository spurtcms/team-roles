package teamroles

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "root",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurt-cms-apr3",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

// test role list
func TestRoleList(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {

		rolelist, count, err := PermissionFlg.RoleList(rolelist{Limit: 10, Offset: 0})

		if err != nil {

			panic(err)
		}

		fmt.Println(rolelist, count)
	} else {

		log.Println("permissions enabled not initialised")
	}

}

// test create role
func TestCreateRole(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {

		rolecreate, err := PermissionFlg.CreateRole(RoleCreation{Name: "Manager", Description: "", CreatedBy: 1})

		if err != nil {

			panic(err)
		}

		fmt.Println(rolecreate)
	} else {

		log.Println("permissions enabled not initialised")
	}
}

// test update role
func TestUpdateRole(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {

		rolecreate, err := PermissionFlg.UpdateRole(RoleCreation{Name: "Manager", Description: "deportment of marketting", CreatedBy: 1}, 3)

		if err != nil {

			panic(err)
		}

		fmt.Println(rolecreate)
	} else {

		log.Println("permissions enabled not initialised")
	}

}

// test Delete role
func TestDeleteRole(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {

		flg, err := PermissionFlg.DeleteRole(1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")
	}

}

// test checkrole
func TestCheckRole(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {
		flg, err := PermissionFlg.CheckRoleAlreadyExists(3, "Manager")

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")
	}

}

// test GetroleByid
func TestGetRoleById(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		DB:         db,
		RoleId:     1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Roles & Permissions", auth.CRUD)

	PermissionFlg := RoleSetup(Config{
		AuthEnable:       true,
		PermissionEnable: true,
		DB:               db,
		Authenticate:     Auth,
	})

	if permisison {
		role, err := PermissionFlg.GetRoleById(3)

		if err != nil {

			panic(err)
		}

		fmt.Println(role)
	} else {

		log.Println("permissions enabled not initialised")
	}

}