// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID   int32  `gorm:"column:id;type:int;primaryKey" json:"id"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
	Articles []Article `gorm:"many2many:article_tags;"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}