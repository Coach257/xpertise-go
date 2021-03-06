package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
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

// UpdateAUserUserType 更新用户状态
func UpdateAUserUserType(user *model.User, userType int) error {
	user.UserType = userType
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

func CreateAFavorite(userID uint64, paperID string, paperInfo string) (err error) {
	favorite := model.Favorite{UserID: userID, PaperID: paperID, PaperInfo: paperInfo}
	if err = global.DB.Create(&favorite).Error; err != nil {
		return err
	}
	return
}

func DeleteAFavorite(favorID uint64) (err error) {
	var favorite model.Favorite
	notFound := global.DB.First(&favorite, favorID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&favorite).Error
	return err
}

func QueryAllFavorites(userID uint64) (favorites []model.Favorite) {
	global.DB.Where("user_id = ?", userID).Find(&favorites)
	return favorites
}

func CreateAWish(userID uint64, paperID string, paperType string, title string, url string, year string, citation uint64) error {
	wish := model.Wish{UserID: userID, PaperID: paperID, PaperType: paperType, Title: title, URL: url, Citation: citation, PaperPublishYear: year}
	err := global.DB.Create(&wish).Error
	return err
}

func QueryAWish(userID uint64, paperID string) (wish model.Wish, notFound bool) {
	notFound = global.DB.Where("user_id = ? AND paper_id = ?", userID, paperID).First(&wish).RecordNotFound()
	return wish, notFound
}

func DeleteAWish(wishID uint64) error {
	var wish model.Wish
	if notFound := global.DB.First(&wish, wishID).RecordNotFound(); notFound {
		return gorm.ErrRecordNotFound
	}
	err := global.DB.Delete(&wish).Error
	return err
}

func QueryAllWishes(userID uint64) (wishes []model.Wish) {
	global.DB.Where("user_id = ?", userID).Find(&wishes)
	return wishes
}

func DeleteAUserByID(userID uint64) (err error) {
	user, _ := QueryAUserByID(userID)
	err = global.DB.Delete(&user).Error
	return err
}

func FindFavByUPID(userID string, paperID string) (favorite model.Favorite, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).Where("paper_id = ?", paperID).First(&favorite).RecordNotFound()
	return favorite, notFound
}
