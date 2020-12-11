package service

import (
	"xpertise-go/global"
	"xpertise-go/model"
)

func QueryAPaperByID(paperID string) (papers model.Paper, err error) {
	err = global.DB.Where("paper_id = ?", papers).Error
	return papers, err
}
