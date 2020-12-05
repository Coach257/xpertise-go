package dao

import (
	"time"
)

// Organization Model.
type Organization struct {
	OrgID        uint64   `gorm:"primary_key;" json:"org_id"`
	OrgNameCn    string   `gorm:"size:20;not null;unique" json:"org_name_cn" binding:"required"`
	OrgNameEn    string   `gorm:"size:40;not null;unique" json:"org_name_en" binding:"required"`
	OrgNameEns   string   `gorm:"size:10;" json:"org_name_ens"`
	BasicInfo    string   `gorm:"size:200;" json:"basic_info"`
	SumOfDocs    uint32   `json:"sum_of_docs"`
	SumOfPatents uint32   `json:"sum_of_patents"`
	AvgScore     float32  `json:"avg_score"`
	AdminID      uint64   `gorm:"not null;" json:"admin_id"`
	PortalList   UintList `gorm:"type:varchar(200)" json:"portal_list"`
}

// Document Model.
type Document struct {
	DocID       uint64    `gorm:"primary_key;" json:"doc_id"`
	Title       string    `gorm:"size:250;not null;" json:"title" binding:"required"`
	AuthorList  StrList   `gorm:"type:varchar(300)" json:"author_list"`
	TypeList    StrList   `gorm:"type:varchar(300)" json:"type_list"`
	Abstract    string    `gorm:"size:5000" json:"abstract"`
	CiteList    StrList   `gorm:"type:varchar(500)" json:"cite_list"`
	PublishTime time.Time `json:"time"`
	Source      string    `gorm:"size:30" json:"source"`
	Original    string    `gorm:"size:1000" json:"original"` // same as content
	Comments    []Comment `gorm:"ForeignKey:ComID" json:"comments"`
	ComID       uint64
	// comments & content要吗，你们商量下
}

// OrgDocument Model.
type OrgDocument struct {
	Organization []Organization `gorm:"ForeignKey:OrgID" json:"org"`
	Document     []Document     `gorm:"ForeignKey:DocID" json:"doc"`
	// GORM need an id for each foreign key
	OrgID uint64
	DocID uint64
}

// OrgPatent Model.
type OrgPatent struct {
	PatentID     uint64         `gorm:"primary_key;" json:"patent_id"`
	InventorList StrList        `gorm:"type:varchar(200);not null;" json:"inventor_list"`
	PatentType   int            `json:"patent_type"`
	PatentCode   string         `gorm:"size:20" json:"patent_code"`
	PatentClass  string         `gorm:"size:40" json:"patent_class"`
	AgentName    string         `gorm:"size:20" json:"agent_name"`
	Organization []Organization `gorm:"ForeignKey:OrgID;not null;" json:"org"`
	OrgID        uint64
}
