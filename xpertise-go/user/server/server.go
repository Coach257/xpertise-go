package server

import (
	"xpertise-go/dao"
)

// CreateAUser : Create a table user.
func CreateAUser(user *dao.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

func QueryAUerById(userid uint64)(user dao.User,NotFound bool)  {
	NotFound = dao.DB.First(&user,userid).RecordNotFound()
	return user, NotFound
}

func QueryAUserByUsername(username string) (user dao.User, NotFound bool) {
	NotFound = dao.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, NotFound
}

func QueryAUserByEmail(email string) (user dao.User, NotFound bool) {
	NotFound = dao.DB.Where("email = ?", email).First(&user).RecordNotFound()
	return user, NotFound
}

func UpdateAUserPassword(user *dao.User, newPassword string) error {
	user.Password = newPassword
	err := dao.DB.Save(user).Error
	return err
}

func UpdateAUser(user *dao.User,username string,email string,info string)error{
	user.Username=username
	user.Email=email
	user.BasicInfo=info
	err:=dao.DB.Save(user).Error
	return err
}

func CountUsersByUsername(username string)(count int,err error){
	dao.DB.Where("username = ? ",username).Count(&count)
	return count,err
}

func CountUsersByEmail(email string)(count int,err error){
	dao.DB.Where("email = ? ",email).Count(&count)
	return count,err
}

func CreateAFolder(folderName string,folderInfo string,userId uint64)(error,uint64){
	folder := dao.Folder{FolderName: folderName,FolderInfo: folderInfo,UserID: userId}
	err:=dao.DB.Create(&folder).Error
	return err,folder.FolderID
}

func QueryAFolderByID(folderId string)(folder dao.Folder,NotFound bool){
	NotFound=dao.DB.First(&folder,folderId).RecordNotFound()
	return folder,NotFound
}

func CreateAFavorite(folderId uint64,docId uint64,docInfo string)(error,uint64){
	favorite :=dao.Favorite{FolderID: folderId,DocID: docId,DocInfo: docInfo}
	err :=dao.DB.Create(&favorite).Error
	return err,favorite.FavorID
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
