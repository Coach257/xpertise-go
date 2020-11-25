package dao

// User Model.
type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
}

// Student model for test.
type Student struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
	Age  uint
}
