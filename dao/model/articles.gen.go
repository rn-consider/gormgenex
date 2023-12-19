// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameArticle = "articles"

// Article mapped from table <articles>
type Article struct {
	ID     int32  `gorm:"column:id;type:int;primaryKey" json:"id"`
	Title  string `gorm:"column:title;type:varchar(255)" json:"title"`
	UserID int32  `gorm:"column:user_id;type:int;index:user_id,priority:1" json:"user_id"`
	User   User   // 文章所属的用户
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}