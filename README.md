# team-roles Package

The 'team-roles' package empowers administrators to create specialized teams with diverse users and roles tailored for distinct domains. 

## Features

- Creates access control settings for members, specifying the level of access they have to specific resources or functionalities within the system.
- Updates the existing control settings with new configuration of permissions.


# Installation

``` bash
go get github.com/spurtcms/team-roles
```
# Usage Example


```bash
import (
		"github.com/spurtcms/auth"
		role "github.com/spurtcms/team-roles"
	)

	func main(){
		
		NewAuth := auth.AuthSetup(newauth.Config{
			SecretKey: os.Getenv("JWT_SECRET"),
			DB:        DB,
		})
		
		token, _ := Auth.CreateToken()

		Auth.VerifyToken(token, SecretKey)

		permission, _ := Auth.IsGranted("Roles & Permission", auth.CRUD)
		
		NewRoleWP := role.RoleSetup(role.Config{
			DB:               DB,
			AuthEnable:       true,
			PermissionEnable: false,
			Authenticate:     NewAuth,
		})
		
		if permission{
		
			//list role
			roles, rolecount, err := NewRole.RoleList(role.Rolelist{Limit: 10, Offset: 0})
			if err!=nil{
				//handle error
				fmt.Println(err)
			}
			fmt.Println(roles,rolecount)
			
			//create role
			rolecreate, err := PermissionFlg.CreateRole(RoleCreation{Name: "Manager", Description: "", CreatedBy: 1})
			if err != nil {
				//handle error
				fmt.Println(err)
			}
			fmt.Println(rolecreate)
			
			//update role
			roleupdate, err := PermissionFlg.UpdateRole(RoleCreation{Name: "Manager", Description: "deportment of marketting", CreatedBy: 1},3)
			if err != nil {
				//handle error
				fmt.Println(err)
			}
			fmt.Println(roleupdate)
			
			//delete role 1st param multiple delete , 2nd param single delete
			flg, err := PermissionFlg.DeleteRole([]int{}, 1)
			if err != nil {
				panic(err)
			}
			fmt.Println(flg)
		
		}

	}	

	```




	# Getting help
	If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/team-roles/issues]. 

