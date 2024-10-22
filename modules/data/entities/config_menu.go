package entities

type ConfigMenu struct {
	Id        int             `gorm:"column:id;primaryKey" json:"id"`
	Name      string          `gorm:"column:name;unique" json:"name"`
	Status    string          `gorm:"column:status" json:"status"`
	SideBar   bool            `gorm:"column:side_bar;type:bit" json:"sideBar"`
	CompanyId string          `gorm:"column:company_id" json:"companyId"`
	SubMenu   []ConfigSubMenu `gorm:"foreignKey:MenuID"`
}

func (ConfigMenu) TableName() string {
	return "Menu" // Specify the actual table name here
}

type ConfigSubMenu struct {
	Id        int    `gorm:"column:id;primaryKey" json:"id"`
	MenuID    int    `gorm:"column:menu_id" json:"menuId"`
	SideBar   bool   `gorm:"column:side_bar;type:bit " json:"sideBar"`
	Name      string `gorm:"column:name;unique" json:"name"`
	Status    string `gorm:"column:status" json:"status"`
	CompanyId string `gorm:"column:company_id" json:"companyId"`
}

func (ConfigSubMenu) TableName() string {
	return "Sub_Menu" // Specify the actual table name here
}
