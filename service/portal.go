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
func QueryAColumnByID(authorID string) (col model.SpecialColumn, notFound bool) {
	notFound = global.DB.Where("author_id = ? ", authorID).First(&col).RecordNotFound()
	return col, notFound
}
func CreateAColumn(authorID string, columnName string) (err error) {
	specialColumn := model.SpecialColumn{AuthorID: authorID, ColumnName: columnName}
	if err = global.DB.Create(&specialColumn).Error; err != nil {
		return err
	}
	return
}

func AddPaperToColumn(columnID uint64, paperID string, paperTitle string) (err error) {
	columnPaper := model.ColumnPaper{ColumnID: columnID, PaperID: paperID, PaperTitle: paperTitle}
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

// CreateAPortal 初始化门户
func CreateAPortal(userID uint64, authorID string) (err error) {
	portal := model.Portal{
		UserID:   userID,
		AuthorID: authorID,
	}
	if global.DB.Create(&portal).Error != nil {
		return err
	}
	return
}

// 创建一条推荐
func CreateARecommend(authorID string, authorName string, paperID string, reason string) (err error) {
	recommend := model.Recommend{
		AuthorID:   authorID,
		AuthorName: authorName,
		PaperID:    paperID,
		Reason:     reason,
	}
	if global.DB.Create(&recommend).Error != nil {
		return err
	}
	return
}

// 删除某条评论
func DeleteRecommend(authorID string, paperID string) (err error) {
	var recommend model.Recommend
	notFound := global.DB.Where(&model.Recommend{AuthorID: authorID, PaperID: paperID}).First(&recommend).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&recommend).Error
	return err
}

// 列出某个作者的所有推荐
func QueryRecommendsFromOneAuthor(authorID string) (recommends []model.Recommend) {
	global.DB.Where("author_id = ?", authorID).Find(&recommends)
	return recommends
}

// 列出某个文献的所有推荐
func QueryRecommendsFromOnePaper(paperID string) (recommends []model.Recommend) {
	global.DB.Where("paper_id = ?", paperID).Find(&recommends)
	return recommends
}

// 列出推荐数目最多的前七篇文献
//db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
func QueryTopSevenPapers() (results []model.Result) {
	global.DB.Table("recommends").Select("paper_id as paper_id, count(author_id) as total").Group("paper_id").Order("total desc").Limit(7).Scan(&results)
	return results // db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)

}

func FindPortalByID(authorID string) (portal model.Portal, notFound bool) {
	notFound = global.DB.Where("author_id = ?", authorID).First(&portal).RecordNotFound()
	return portal, notFound
}
