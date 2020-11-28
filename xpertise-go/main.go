package main

import (
	adminDao "xpertise-go/admin/dao"
	adminRouter "xpertise-go/admin/router"
	branchDao "xpertise-go/branch/dao"
	branchRouter "xpertise-go/branch/router"
	portalDao "xpertise-go/portal/dao"
	portalRouter "xpertise-go/portal/router"
	userDao "xpertise-go/user/dao"
	userRouter "xpertise-go/user/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	adminErr := adminDao.InitMySQL()
	if adminErr != nil {
		panic(adminErr)
	}
	defer adminDao.Close()
	adminR := adminRouter.SetupRouter()
	adminR.Run()

	branchErr := branchDao.InitMySQL()
	if branchErr != nil {
		panic(branchErr)
	}
	defer branchDao.Close()
	branchR := branchRouter.SetupRouter()
	branchR.Run()

	portalErr := portalDao.InitMySQL()
	if portalErr != nil {
		panic(portalErr)
	}
	defer portalDao.Close()
	portalR := portalRouter.SetupRouter()
	portalR.Run()

	userErr := userDao.InitMySQL()
	if userErr != nil {
		panic(adminErr)
	}
	defer userDao.Close()
	userR := userRouter.SetupRouter()
	userR.Run()
}
