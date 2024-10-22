package entities

type Log struct {
	Id          int    `gorm:"column:id;primaryKey" json:"id"`
	EndPoint    string `gorm:"column:end_point" json:"end_point"`
	Description string `gorm:"column:description" json:"description"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	UpdatedBy   string `gorm:"column:updated_by" json:"updated_by"`
}
