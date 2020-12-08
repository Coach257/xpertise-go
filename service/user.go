package service

import (
	"xpertise-go/global"
	"xpertise-go/model"
)

func CreateAUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}
func QueryAUserByID(userID uint64) (user model.User, notFound bool) {
	notFound = global.DB.First(&user, userID).RecordNotFound()
	return user, notFound
}
func QueryAUserByUsername(username string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, notFound
}
func QueryAUserByEmail(email string) (user model.User, notFound bool) {
	notFound = global.DB.Where("email = ?", email).First(&user).RecordNotFound()
	return user, notFound
}
func UpdateAUserPassword(user *model.User, newPassword string) error {
	user.Password = newPassword
	err := global.DB.Save(user).Error
	return err
}

func UpdateAUser(user *model.User, username string, password string, email string, info string) error {
	user.Username = username
	user.Password = password
	user.Email = email
	user.BasicInfo = info
	err := global.DB.Save(user).Error
	return err
}
func CountUsersByUsername(username string) (count int, err error) {
	global.DB.Where("username = ? ", username).Count(&count)
	return count, err
}

func CountUsersByEmail(email string) (count int, err error) {
	global.DB.Where("email = ? ", email).Count(&count)
	return count, err
}

func CreateAFolder(folderName string, folderInfo string, userID uint64) (uint64, error) {
	folder := model.Folder{FolderName: folderName, FolderInfo: folderInfo, UserID: userID}
	err := global.DB.Create(&folder).Error
	return folder.FolderID, err
}

func QueryAFolderByID(folderID string) (folder model.Folder, notFound bool) {
	notFound = global.DB.First(&folder, folderID).RecordNotFound()
	return folder, notFound
}

func CreateAFavorite(folderID uint64, doc model.Document) (uint64, error) {
	favorite := model.Favorite{FolderID: folderID, Document: doc}
	err := global.DB.Create(&favorite).Error
	return favorite.FavorID, err
}
