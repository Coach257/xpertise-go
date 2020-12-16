package service

import (
	"fmt"
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

func QueryAPaperByID(paperID string) (paper model.Paper, err error) {
	err = global.DB.Where("paper_id = ?", paperID).Find(&paper).Error
	fmt.Println(paper.PaperID, paper.Title)
	return paper, err
}

func QueryAnAuthorByID(authorID string) (author model.Author, notFound bool) {
	notFound = global.DB.Where("author_id = ? ", authorID).First(&author).RecordNotFound()
	return author, notFound
}

func CreateAColumn(authorID string, columnName string) (err error) {
	specialColumn := model.SpecialColumn{AuthorID: authorID, ColumnName: columnName}
	if err = global.DB.Create(&specialColumn).Error; err != nil {
		return err
	}
	return
}

func AddPaperToColumn(columnID uint64, paperID string) (err error) {
	columnPaper := model.ColumnPaper{ColumnID: columnID, PaperID: paperID}
	if err = global.DB.Create(&columnPaper).Error; err != nil {
		return err
	}
	return
}

// 查询ColumnPaper表，判断专栏是否存在、文献是否收藏在该专栏中
func QueryItemFromColumnPaper(columnID uint64, paperID string) (columnPaper model.ColumnPaper, notFound bool) {
	notFound = global.DB.Where("column_id = ?", columnID).Where("paper_id = ?", paperID).First(&columnPaper).RecordNotFound()
	return columnPaper, notFound
}

// 列出某专栏中的所有内容
func QueryAllFromAColumn(columnID uint64) (columnPapers []model.ColumnPaper) {
	global.DB.Where("column_id = ?", columnID).Find(&columnPapers)
	return columnPapers
}

// 从某一专栏中删除某文献
func DeleteOnePaperFromAColumn(columnID uint64, paperID string) (err error) {
	var columnPaper model.ColumnPaper
	notFound := global.DB.Where("column_id = ?", columnID).Where("paper_id = ?", paperID).First(&columnPaper).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&columnPaper).Error
	return err
}
