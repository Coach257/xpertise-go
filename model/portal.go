package model

type SpecialColumn struct {
	ColumnID   uint64 `gorm:"primary_key; not null;" json:"column_id"`
	AuthorID   string `gorm:"type:varchar(10);primary_key;" json:"author_id"`
	ColumnName string `gorm:"type:varchar(100)" json:"column_name"`
}

type ColumnPaper struct {
	ColumnID uint64 `gorm:"primary_key; not null;" json:"column_id"`
	PaperID  string `gorm:"type:varchar(10);primary_key; not null;" json:"paper_id"`
}

// Portal 记录认证用户专栏列表
type Portal struct {
	//PortalID       uint64          `gorm:"primary_key" json:"authuseraff_id"`
	UserID         uint64          `gorm:"primary_key" json:"user_id"`
	AuthorID       string          `gorm:"type:varchar(10);primary_key" json:"author_id"`
	SpecialColumns []SpecialColumn `gorm:"ForeignKey:SpecialColumn;AssociationForeignKey:SpecialColumn" json:"special_column"`
}
