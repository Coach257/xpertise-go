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
func QueryAColumnByID(authorID string) (col []model.SpecialColumn) {
	global.DB.Where("author_id = ? ", authorID).Find(&col)
	//notFound = global.DB.Where("author_id = ? ", authorID).First(&col).RecordNotFound()
	return col
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
func CreateARecommend(authorID string, authorName string, paperID string, citation uint64, reason string) (err error) {
	recommend := model.Recommend{
		AuthorID:   authorID,
		AuthorName: authorName,
		PaperID:    paperID,
		Citation:   citation,
		Reason:     reason,
	}
	if global.DB.Create(&recommend).Error != nil {
		return err
	}
	return
}

// 查看是否在推荐统计表中，返回查到的Paper
func QueryARecommendInPaperRecommend(paperID string) (paperRecommend model.PaperRecommend, notFound bool) {
	notFound = global.DB.Where("paper_id = ? ", paperID).First(&paperRecommend).RecordNotFound()
	return paperRecommend, notFound
}

// CalculateScore 计算推荐指数
func CalculateScore(citation uint64, hindex uint64) (value uint64) {
	value = citation + hindex
	return value
}

// 加入至论文推荐统计表
func AddToPaperRecommend(paperID string, paperTitle string, citation uint64, hindex int64) (err error) {
	//value:=CalculateScore(citation,hindex)
	value := int64(citation) + hindex
	paperRecommend := model.PaperRecommend{PaperID: paperID, PaperTitle: paperTitle, Value: value}
	err = global.DB.Create(&paperRecommend).Error
	return err
}

// 更新论文在论文推荐统计表中的数据
func UpdatePaperRecommend(paperRecommend *model.PaperRecommend, hindex int64) (err error) {
	paperRecommend.Value += hindex
	err = global.DB.Save(paperRecommend).Error
	return err
}

// 查看是否在推荐统计表中，返回查到的Paper
func QueryARecommendInCsPaperRecommend(paperID string) (paperRecommend model.CsPaperRecommend, notFound bool) {
	notFound = global.DB.Where("paper_id = ? ", paperID).First(&paperRecommend).RecordNotFound()
	return paperRecommend, notFound
}

// 加入至论文推荐统计表
func AddToCsPaperRecommend(paperID string, paperTitle string, citation uint64, hindex int64) (err error) {
	//value:=CalculateScore(citation,hindex)
	value := int64(citation) + hindex
	paperRecommend := model.CsPaperRecommend{PaperID: paperID, PaperTitle: paperTitle, Value: value}
	err = global.DB.Create(&paperRecommend).Error
	return err
}

// 更新论文在论文推荐统计表中的数据
func UpdateCsPaperRecommend(paperRecommend *model.CsPaperRecommend, hindex int64) (err error) {
	paperRecommend.Value += hindex
	err = global.DB.Save(paperRecommend).Error
	return err
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
func QueryTopSevenPapers() (results []model.Result) {
	global.DB.Table("paper_recommends").Order("value desc").Limit(7).Scan(&results)
	return results
}

func QueryTopSevenCsPapers() (results []model.Result) {
	global.DB.Table("cs_paper_recommends").Order("value desc").Limit(7).Scan(&results)
	return results
}

func FindPortalByAuthorID(authorID string) (portal model.Portal, notFound bool) {
	notFound = global.DB.Where("author_id = ?", authorID).First(&portal).RecordNotFound()
	return portal, notFound
}

func FindPortalByUserID(userID uint64) (portal model.Portal, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).First(&portal).RecordNotFound()
	return portal, notFound
}

func FindDirectConnectedAuthors(authorID string) (connections []model.Connection, err error) {
	err = global.DB.Where("author1_id = ?", authorID).Or("author2_id = ?", authorID).Find(&connections).Error
	return connections, err
}

func FindAuthorConnections(authorID string, tot int) (connections []model.Connection, err error) {
	var tmpConnections []model.Connection
	err = global.DB.Where("author1_id = ?", authorID).Or("author2_id = ?", authorID).Find(&tmpConnections).Error
	connections = append(connections, tmpConnections...)
	// for _, _ := range tmpConnections {
	// 	if tmpConnections[0].Author2ID != authorID {
	// 		authorID = tmpConnections[0].Author2ID
	// 	} else {
	// 		authorID = tmpConnections[0].Author1ID
	// 	}
	// }

	return connections, err
}
