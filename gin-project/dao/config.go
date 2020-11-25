package dao

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("szl:123@unix(/%s/test-project-296508:asia-east2:test)/test?parseTime=true", socketDir)

	DB, err = gorm.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	}

	// 为模型`User`创建表
	//DB.CreateTable(&User{})

	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	//DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
