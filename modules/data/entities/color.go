package entities

type Colors struct {
	Id       int    `gorm:"primaryKey;autoIncrement;column:id"`
	Name     string `gorm:"column:name" json:"name"`
	Code     string `gorm:"column:code" json:"code"`
	CodeName string `gorm:"column:code_name" json:"code_name"`
}
