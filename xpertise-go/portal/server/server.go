package server

import "xpertise-go/portal/dao"

// CreateADocument is a test service
func CreateADocument(doc *dao.Document) (err error) {
	print(doc.Abstract)
	if err = dao.DB.Create(&doc).Error; err != nil {
		return err
	}
	return
}

// QueryDocument is a test service
func QueryDocument(id uint64) (doc []*dao.Document) {
	dao.DB.First(&doc, id)
	return doc
}
