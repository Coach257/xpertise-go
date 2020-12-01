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
	// 为模型`User`创建表
	DB.AutoMigrate(&User{})
	DB.DropTable("comments")
	DB.CreateTable(&Comment{})
	DB.CreateTable(&Document{})
	DB.CreateTable(&LikeDislikeRecord{})

	//@ztx Document表的字段有变化，需要再斟酌下里面要有什么，document放在portal.go里
	// doc1 := Document{PublishTime: time.Now(), AuthorList: "[ztx]", Abstract: "这是一篇文章"}
	// DB.Create(&doc1)
	// com1 := Comment{DocID: doc1.ID, Content: "这是第一篇文章的第一篇评论"}
	// DB.Create(&com1)
	// thumbup1 := LikeDislikeRecord{ComID: com1.ID, UserID: 3, IsLikeOrDis: true}
	// DB.Create((&thumbup1))
	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	// 自动表迁移
	// DB.AutoMigrate(&Student{})

	return DB.DB().Ping()
}

// Close the connection of the database.
func Close() {
	DB.Close()
}
