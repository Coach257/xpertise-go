package dao

import (
	"github.com/jinzhu/gorm"
)

// ...
var (
	DB *gorm.DB
)

// InitMySQL means initialize database.
func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:@buaa21@tcp(101.132.227.56:3306)/user_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// 为模型`User`创建表
	DB.CreateTable(&User{})

	user := User{UserID: 18373059, Username: "IAmParasite", Password: "123", Email: "1004181396@qq.com",
		Usertype: 1}
	DB.Create(&user)
	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	return DB.DB().Ping()
}

// Close the connection of the database.
func Close() {
	DB.Close()
}
