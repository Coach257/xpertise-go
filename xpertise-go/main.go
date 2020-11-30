package main

import (
	adminDao "xpertise-go/admin/dao"
	branchDao "xpertise-go/branch/dao"
	portalDao "xpertise-go/portal/dao"
	"xpertise-go/router"
	userDao "xpertise-go/user/dao"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	userErr := userDao.InitMySQL()
	if userErr != nil {
		panic(userErr)
	}
	defer userDao.Close()

	adminErr := adminDao.InitMySQL()
	if adminErr != nil {
		panic(adminErr)
	}
	defer adminDao.Close()

	branchErr := branchDao.InitMySQL()
	if branchErr != nil {
		panic(branchErr)
	}
	defer branchDao.Close()

	portalErr := portalDao.InitMySQL()
	if portalErr != nil {
		panic(portalErr)
	}
	defer portalDao.Close()

	r := router.SetupRouter()
	r.Run()
}
