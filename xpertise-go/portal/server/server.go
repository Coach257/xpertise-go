package server

import "xpertise-go/dao"

// DocCreate is a test service
func DocCreate(doc *dao.Document) (err error) {
	print(doc.Abstract)
	if err = dao.DB.Create(&doc).Error; err != nil {
		return err
	}
	return
}

// DocQueryID get a document by its id
func DocQueryID(id uint64) (doc dao.Document) {
	//println(id)
	dao.DB.First(&doc, id)
	return doc
}

// DocQueryTitle get an document by its en/cn name
// brute force fuzzy query
func DocQueryTitle(title string) (docs []*dao.Document) {
	dao.DB.Where("title LIKE ?", "%"+title+"%").Find(&docs)
	return docs
}

// OrgCreate
func OrgCreate(org *dao.Organization) (err error) {
	print(org.OrgNameCn)
	if err = dao.DB.Create(&org).Error; err != nil {
		return err
	}
	return
}

// OrgQueryID get an organization by its id
func OrgQueryID(id uint64) (org dao.Organization) {
	dao.DB.First(&org, id)
	return org
}

// OrgQueryName get an organization by its en/cn name
// brute force fuzzy query
func OrgQueryName(name string) (orgs []*dao.Organization) {
	dao.DB.Where("org_name_en LIKE ? OR org_name_cn LIKE ?", "%"+name+"%", "%"+name+"%").Find(&orgs)
	return orgs
}
