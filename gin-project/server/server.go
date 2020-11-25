package server

import "gin-project/dao"

// CreateAUser : Create a table user.
func CreateAUser(user *dao.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

// CreateAStudent : Create a table user.
func CreateAStudent(student *dao.Student) (err error) {
	if err = dao.DB.Create(&student).Error; err != nil {
		return err
	}
	return
}

// DeleteAStudentByID : Create a table user.
func DeleteAStudentByID(StudentID uint) (err error) {
	dao.DB.Where("ID = ?", StudentID).Delete(&dao.Student{})
	return
}

// GetFirstInStudent returns the first message of the table student.
func GetFirstInStudent() (err error) {
	dao.DB.First(&dao.Student{})
	return
}
