package dao

import (
	"github.com/jinzhu/gorm"
)

// ...
var (
	DB *gorm.DB
)

// InitMySQL means initialize database.
func InitMySQL() (error error) {
	DB, error = gorm.Open("mysql", "root:@buaa21@tcp(101.132.227.56:3306)/xpertise_db?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		panic(error)
	}

	// migration usage see: [https://gorm.io/docs/migration.html]
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Folder{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Favorite{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Comment{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Document{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&LikeDislikeRecord{})

	return DB.DB().Ping()
}

// Close the connection of the database.
func Close() {
	DB.Close()
}
