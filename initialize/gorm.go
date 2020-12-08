package initialize

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

func InitMySQL() (error error) {
	global.DB, error = gorm.Open("mysql", "root:@buaa21@tcp(101.132.227.56:3306)/xpertise_db?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		panic(error)
	}

	// migration usage see: [https://gorm.io/docs/migration.html]
	// 仅仅新增新增的字段，不会进行修改已有的字段类型，删除字段的操作
	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		// user
		&model.User{},
		&model.Folder{},
		&model.Favorite{},

		// potal
		&model.Document{},

		// branch
		&model.Document{},

		// admin
	)

	return global.DB.DB().Ping()
}

// Close the connection of the database.
func Close() {
	global.DB.Close()
}
