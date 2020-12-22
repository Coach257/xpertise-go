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
func CalculateScore(citation uint64, h_index uint64) (value uint64) {
	value = citation + h_index
	return value
}

// 加入至论文推荐统计表
func AddToPaperRecommend(paperID string, paperTitle string, citation uint64, h_index int64) (err error) {
	//value:=CalculateScore(citation,h_index)
	value := int64(citation) + h_index
	paperRecommend := model.PaperRecommend{PaperID: paperID, PaperTitle: paperTitle, Value: value}
	err = global.DB.Create(&paperRecommend).Error
	return err
}

// 更新论文在论文推荐统计表中的数据
func UpdatePaperRecommend(paperRecommend *model.PaperRecommend, h_index int64) (err error) {
	paperRecommend.Value += h_index
	err = global.DB.Save(paperRecommend).Error
	return err
}

// 查看是否在推荐统计表中，返回查到的Paper
func QueryARecommendInCsPaperRecommend(paperID string) (paperRecommend model.CsPaperRecommend, notFound bool) {
	notFound = global.DB.Where("paper_id = ? ", paperID).First(&paperRecommend).RecordNotFound()
	return paperRecommend, notFound
}

// 加入至论文推荐统计表
func AddToCsPaperRecommend(paperID string, paperTitle string, citation uint64, h_index int64) (err error) {
	//value:=CalculateScore(citation,h_index)
	value := int64(citation) + h_index
	paperRecommend := model.CsPaperRecommend{PaperID: paperID, PaperTitle: paperTitle, Value: value}
	err = global.DB.Create(&paperRecommend).Error
	return err
}

// 更新论文在论文推荐统计表中的数据
func UpdateCsPaperRecommend(paperRecommend *model.CsPaperRecommend, h_index int64) (err error) {
	paperRecommend.Value += h_index
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

func AddIntoUniqueConnections(curTot int, connections []model.Connection, connection model.Connection) (int, []model.Connection) {
	// stat := false
	for _, v := range connections {
		if connection == v {
			// stat = true
			// fmt.Println("!!!", connection)
			return curTot, connections
		}
	}
	// fmt.Println("unique_____________", connection)
	connections = append(connections, connection)
	curTot++
	return curTot, connections
}

func FormatConnections(res []model.Connection) (a model.A, err error) {
	var b []model.B
	var c []model.C
	for _, s := range res {
		x := model.B{Name: s.Author1Name, Value: s.Author1HIndex}
		stat := false
		for _, p := range b {
			if p.Name == x.Name {
				stat = true
				break
			}
		}
		if stat == false {
			b = append(b, x)
		}
		x = model.B{Name: s.Author2Name, Value: s.Author2HIndex}
		stat = false
		for _, p := range b {
			if p.Name == x.Name {
				stat = true
				break
			}
		}
		if stat == false {
			b = append(b, x)
		}
	}

	for _, s := range res {
		x := model.C{Source: s.Author1Name, Target: s.Author2Name, Num: s.CoNum}
		c = append(c, x)
	}
	a = model.A{Bs: b, Cs: c}
	return a, err
}

func FindAuthorConnections(tot int, authorID string) (a model.A, err error) {
	curTot := 0
	var level1Connections, level2Connections, level3Connections, connections []model.Connection
	err = global.DB.Where("author1_id = ?", authorID).Or("author2_id = ?", authorID).Find(&level1Connections).Error
	connections = append(connections, level1Connections...)
	curTot += len(level1Connections)
	for _, v := range level1Connections {
		var tmpAuthorID string
		if v.Author2ID != authorID {
			tmpAuthorID = v.Author2ID
		} else {
			tmpAuthorID = v.Author1ID
		}
		// fmt.Println(tmpAuthorID)
		err = global.DB.Where("author1_id = ?", tmpAuthorID).Or("author2_id = ?", tmpAuthorID).Find(&level2Connections).Error
		// fmt.Println(level2Connections)
		for _, e := range level2Connections {
			curTot, connections = AddIntoUniqueConnections(curTot, connections, e)
			if curTot >= tot {
				a, err = FormatConnections(connections)
				return a, err
			}
			//fmt.Println(e)
		}
		for _, e := range level2Connections {
			var tmpAuthorID string
			if e.Author2ID != authorID {
				tmpAuthorID = e.Author2ID
			} else {
				tmpAuthorID = e.Author1ID
			}
			// fmt.Println(tmpAuthorID)
			err = global.DB.Where("author1_id = ?", tmpAuthorID).Or("author2_id = ?", tmpAuthorID).Find(&level3Connections).Error
			for _, s := range level3Connections {
				curTot, connections = AddIntoUniqueConnections(curTot, connections, s)
				if curTot >= tot {
					a, err = FormatConnections(connections)
					return a, err
				}
			}
		}
	}
	a, err = FormatConnections(connections)
	return a, err
}
