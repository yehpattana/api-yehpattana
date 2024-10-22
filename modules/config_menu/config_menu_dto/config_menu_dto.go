package configmenudto

type ConfigMenuRequest struct {
	Name      string `db:"name" json:"name" form:"name"`
	SideBar   bool   `db:"side_bar" json:"sidebar" form:"sideBar"`
	Status    string `db:"status" json:"status" form:"status"`
	CompanyId string `db:"company_id" json:"companyId" form:"companyId"`
}

type UpdateConfigMenuRequest struct {
	Id      int    `db:"id" json:"id" form:"id"`
	Name    string `db:"name" json:"name" form:"name"`
	SideBar bool   `db:"side_bar" json:"sidebar" form:"sideBar"`
	Status  string `db:"status" json:"status" form:"status"`
}

// sub menu
type ConfigSubMenuRequest struct {
	MenuID    int    `db:"menu_id" json:"menuId" form:"menuId"`
	Name      string `db:"name" json:"name" form:"name"`
	Status    string `db:"status" json:"status" form:"status"`
	SideBar   bool   `db:"side_bar" json:"sideBar" form:"sideBar"`
	CompanyId string `db:"company_id" json:"companyId" form:"companyId"`
}

type UpdateConfigSubMenuRequest struct {
	Id      int    `db:"id" json:"id" form:"id"`
	MenuID  int    `db:"menu_id" json:"menuId" form:"menuId"`
	Name    string `db:"name" json:"name" form:"name"`
	Status  string `db:"status" json:"status" form:"status"`
	SideBar bool   `db:"side_bar" json:"sideBar" form:"sideBar"`
}
