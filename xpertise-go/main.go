package main

import (
	"xpertise-go/dao"
	"xpertise-go/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	r := router.SetupRouter()
	r.Run()
}
