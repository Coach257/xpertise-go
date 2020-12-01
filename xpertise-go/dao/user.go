package dao

type User struct {
	UserID       uint64 `gorm:"primary_key; not null"`
	Username     string `gorm:"size:15; not null; unique"`
	Password     string `gorm:"size:20; not null"`
	Email        string `gorm:"size:20; not null; unique"`
	Usertype     int    `gorm:"not null;default:1"`
	BasicInfo    string `gorm:"size:100"`
	Interdiction bool   `gorm:"default:false"`

	Folders    []Folder	`gorm:"ForeignKey:UserID"`//User包含多个Folders
	//User_type Profile `gorm:"ForeignKey:UserRefer"` 指定外键
}

type Folder struct{
	FolderID 	uint64 `gorm:"primary_key"`
	Foldername	string `gorm:"size:15;not null"`
	Folderinfo  string `gorm:"size:100"`
	UserID      uint64  //外键

	Favorites []Favorite `gorm:"ForeignKey:FolderID"`//Folder包含多个Favorites
}

type Favorite struct{
	FavorID 	uint64 `gorm:"primary_key"`
	FolderID	uint64 //外键
	DocID		uint64 //外键
	Docinfo		string `gorm:"size:100"`
}


type Student struct {
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255"`
	Age  uint64
}
