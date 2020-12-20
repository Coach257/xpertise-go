package model

type SpecialColumn struct {
	ColumnID   uint64 `gorm:"primary_key; not null;" json:"column_id"`
	AuthorID   string `gorm:"type:varchar(30);primary_key;" json:"author_id"`
	ColumnName string `gorm:"type:varchar(100)" json:"column_name"`
}

type ColumnPaper struct {
	ColumnID   uint64 `gorm:"primary_key; not null;" json:"column_id"`
	PaperID    string `gorm:"type:varchar(30);primary_key; not null;" json:"paper_id"`
	PaperTitle string `gorm:"type:varchar(400);not null;" json:"paper_title"`
}

// Portal 记录认证用户专栏列表
type Portal struct {
	//PortalID       uint64          `gorm:"primary_key" json:"authuseraff_id"`
	UserID         uint64          `gorm:"primary_key" json:"user_id"`
	AuthorID       string          `gorm:"type:varchar(30);primary_key" json:"author_id"`
	SpecialColumns []SpecialColumn `gorm:"ForeignKey:SpecialColumn;AssociationForeignKey:SpecialColumn" json:"special_column"`
}

// 推荐表，记录认证用户推荐的内容
type Recommend struct {
	AuthorID   string `gorm:"type:varchar(30);primary_key" json:"author_id"`
	PaperID    string `gorm:"type:varchar(30);primary_key" json:"paper_id"`
	AuthorName string `gorm:"type:varchar(100)" json:"author_name"`
	Citation   uint64 `json:"n_citation"`
	Reason     string `gorm:"type:varchar(255)" json:"reason"`
}

type PaperRecommend struct {
	PaperID    string `gorm:"type:varchar(30);primary_key" json:"paper_id"`
	PaperTitle string `gorm:"type:varchar(400)" json:"paper_title"`
	Value      int64  `json:"value"`
}

type Result struct {
	PaperID string `json:"paper_id"`
	Total   uint64 `json:"total"`
}
