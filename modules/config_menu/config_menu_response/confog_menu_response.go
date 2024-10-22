package configmenuresponses

type ConfigMenuResponse struct {
	Id      int                     `json:"id"`
	Name    string                  `json:"name"`
	Status  string                  `json:"status"`
	SubMenu []ConfigSubMenuResponse `json:"subMenus"`
	SideBar bool                    `json:"sideBar"`
}
type ConfigSubMenuResponse struct {
	Id      int    `json:"id"`
	MenuId  int    `json:"menuId"`
	Name    string `json:"name"`
	SideBar bool   `json:"sideBar"`
	Status  string `json:"status"`
}

type ConfigMenuCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
