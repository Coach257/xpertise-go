package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ...
var (
	DB  *gorm.DB
	DB1 *gorm.DB
)

// InitMySQL means initialize database.
func InitMySQL() (err1 error) {

	DB1, err1 = gorm.Open("mysql", "root:@buaa21@tcp(101.132.227.56:3306)/branch_db?charset=utf8&parseTime=True&loc=Local")
	if err1 != nil {
		panic(err1)
	}
	// 为模型`User`创建表
	//DB.CreateTable(&User{})
	//DB1.DropTable("comments")
	//DB1.CreateTable(&Comment{})
	//DB1.CreateTable(&Document{})
	//DB1.CreateTable(&LikeDislikeRecord{})
	var doc1 = Document{PublishTime: time.Now(), AuthorList: "[ztx]", Abstract: "这是一篇文章"}
	DB1.Create(&doc1)
	var com1 = Comment{DocID: doc1.ID, Content: "这是第一篇文章的第一篇评论"}
	DB1.Create(&com1)
	var thumbup1 = LikeDislikeRecord{ComID: com1.ID, UserID: 3, IsLikeOrDis: true}
	DB1.Create((&thumbup1))
	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	// 自动表迁移
	// DB.AutoMigrate(&Student{})

	return DB1.DB().Ping()
}

// Close the connection of the database.
func Close() {
	DB1.Close()
}
