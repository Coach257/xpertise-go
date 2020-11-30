package server

import "xpertise-go/user/dao"

// CreateAUser : Create a table user.
func CreateAUser(user *dao.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

func DeleteAStudentByID(StudentID uint64) {
	dao.DB.Where("ID = ?", StudentID).Delete(&dao.Student{})
	return
}

func UpdateAStudentByAge(student *dao.Student, age uint64) {
	dao.DB.Model(&student).Update("Age", age)
}

func QueryAllStudents() (students []*dao.Student) {
	dao.DB.Find(&students)
	return students
}

func QueryStudentByID(StudentID uint64) (student []*dao.Student) {
	dao.DB.First(&student, StudentID)
	return student
}

func QueryStudentsByAge(age uint64) (students []*dao.Student) {
	dao.DB.Where("Age = ?", age).Find(&students)
	return students
}
