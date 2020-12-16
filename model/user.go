package model

type User struct {
	UserID    uint64     `gorm:"primary_key; not null;" json:"user_id"`
	Username  string     `gorm:"size:15; not null; unique" json:"username"`
	Password  string     `gorm:"size:20; not null" json:"password"`
	Email     string     `gorm:"size:20; not null; unique" json:"email"`
	UserType  int        `gorm:"not null;default:1" json:"user_type"`
	BasicInfo string     `gorm:"size:100" json:"basic_info"`
	Ban       bool       `gorm:"default:false" json:"ban"`
	Favorites []Favorite `gorm:"ForeignKey:UserID;AssociationForeignKey:UserID" json:"favorites"`
	Wishes    []Wish     `gorm:"ForeignKey:UserID;AssociationForeignKey:UserID" json:"wishes"`
}

type Favorite struct {
	FavorID   uint64 `gorm:"primary_key;" json:"favor_id"`
	UserID    uint64 `gorm:"not null" json:"user_id"`
	PaperID   string `gorm:"size:10;not null" json:"paper_id"`
	PaperInfo string `gorm:"size:2560" json:"paper_info"`
}

type Wish struct {
	WishID  uint64 `gorm:"primary_key" json:"wish_id"`
	UserID  uint64 `gorm:"not null" json:"user_id"`
	Title   string `gorm:"not null" json:"title"`
	PaperID string `gorm:"size:10;not null" json:"paper_id"`
	Url     string `gorm:"not null" json:"url"`
}
