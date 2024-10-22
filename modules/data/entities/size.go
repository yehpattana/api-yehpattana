package entities

type Size struct {
	Id   int    `gorm:"primaryKey;autoIncrement;column:id"`
	Size string `gorm:"column:size" json:"size"`
}

func (Size) TableName() string {
	return "Size"
}
