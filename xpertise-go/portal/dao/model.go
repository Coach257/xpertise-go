package dao

import (
	"database/sql/driver"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// customized type for string array store & read
type StrList []string

func (l StrList) Value() (driver.Value, error) {
	data, err := jsoniter.Marshal(l)
	return string(data), err
}

func (l *StrList) Scan(data interface{}) error {
	return jsoniter.Unmarshal(data.([]byte), l)
}

// customized type for uint array store & read
type UintList []uint

func (l UintList) Value() (driver.Value, error) {
	data, err := jsoniter.Marshal(l)
	return string(data), err
}

func (l *UintList) Scan(data interface{}) error {
	err := jsoniter.Unmarshal(data.([]byte), l)
	for i := 0; i < len(*l); i++ {
		print((*l)[i])
	}
	return err
}

// Organization Model.
type Organization struct {
	OrgID        uint64 `gorm:"primary_key;"`
	OrgNameCn    string `gorm:"size:20;not null;unique"`
	OrgNameEn    string `gorm:"size:40;not null;unique"`
	OrgNameEns   string `gorm:"size:10;"`
	BasicInfo    string `gorm:"size:200;"`
	SumOfDocs    uint32
	SumOfPatents uint32
	AvgScore     float32
	AdminID      uint64   `gorm:"not null;"`
	PortalList   UintList `gorm:"type:varchar(200)" json:"portal_list"`
}

// Document Model.
type Document struct {
	DocID       uint64  `gorm:"primary_key;"`
	AuthorList  StrList `gorm:"type:varchar(200)" json:"author_list"`
	TypeList    StrList `gorm:"type:varchar(200)" json:"type_list"`
	Abstract    string  `gorm:"size:150"`
	CiteList    StrList `gorm:"type:varchar(500)" json:"cite_list"`
	PublishTime time.Time
	Source      string `gorm:"size:30"`
	Original    string `gorm:"size:100"`
}

// OrgDocument Model.
type OrgDocument struct {
	OrgID uint64 `gorm:"ForeignKey:OrgID"`
	DocID uint64 `gorm:"ForeignKey:DocID"`
}

// OrgPatent Model.
type OrgPatent struct {
	PatentID     uint64  `gorm:"primary_key;"`
	OrgID        uint64  `gorm:"ForeignKey:OrgID;not null;"`
	InventorList StrList `gorm:"type:varchar(200);not null;" json:"inventor_list"`
	PatentType   int
	PatentCode   string `gorm:"size:20"`
	PatentClass  string `gorm:"size:40"`
	AgentName    string `gorm:"size:20"`
}
