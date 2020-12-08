package model

type User struct {
	UserID    uint64 `gorm:"primary_key; not null" json:"user_id"`
	Username  string `gorm:"size:15; not null; unique" json:"username"`
	Password  string `gorm:"size:20; not null" json:"password"`
	Email     string `gorm:"size:20; not null; unique" json:"email"`
	UserType  int    `gorm:"not null;default:1" json:"user_type"`
	BasicInfo string `gorm:"size:100" json:"basic_info"`
	Ban       bool   `gorm:"default:false" json:"ban"`

	Folders []Folder `gorm:"ForeignKey:UserID" json:"folders"` //User包含多个Folders
}

type Folder struct {
	FolderID   uint64     `gorm:"primary_key" json:"folder_id"`
	FolderName string     `gorm:"size:15;not null" json:"folder_name"`
	FolderInfo string     `gorm:"size:100" json:"folder_info"`
	UserID     uint64     `json:"user_id" json:"user_id"`               //外键
	Favorites  []Favorite `gorm:"ForeignKey:FolderID" json:"favorites"` //Folder包含多个Favorites
}

type Favorite struct {
	FavorID  uint64   `gorm:"primary_key" json:"favor_id"`
	FolderID uint64   `json:"folder_id"` //外键
	Document Document `json:"document"`
}
