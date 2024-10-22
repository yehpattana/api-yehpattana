package repositories

import (
	configmenudto "github.com/natersland/b2b-e-commerce-api/modules/config_menu/config_menu_dto"
	configmenuresponses "github.com/natersland/b2b-e-commerce-api/modules/config_menu/config_menu_response"
	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	"gorm.io/gorm"
)

type ConfigMenuRepositoryInterface interface {
	GetAllConfigMenu(companyId string) ([]*configmenuresponses.ConfigMenuResponse, error)
	CreateConfigMenu(req *configmenudto.ConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	UpdateConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	DeleteConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	GetAllConfigSubMenu(companyId string) ([]*configmenuresponses.ConfigSubMenuResponse, error)
	CreateConfigSubMenu(req *configmenudto.ConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	UpdateConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	DeleteConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
}

type configMenuRepositoryImpl struct {
	*gorm.DB
}

func ConfigMenuRepositoryImpl(db *gorm.DB) ConfigMenuRepositoryInterface {
	return &configMenuRepositoryImpl{
		DB: db,
	}
}

func (configMenuRepository *configMenuRepositoryImpl) GetAllConfigMenu(companyId string) ([]*configmenuresponses.ConfigMenuResponse, error) {
	var configMenus []*configmenuresponses.ConfigMenuResponse
	var menus []*entities.ConfigMenu

	result := configMenuRepository.DB.Preload("SubMenu").Where("company_id = ?", companyId).Find(&menus)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, menu := range menus {
		var subMenus []configmenuresponses.ConfigSubMenuResponse
		for _, subMenu := range menu.SubMenu {
			subMenus = append(subMenus, configmenuresponses.ConfigSubMenuResponse{
				Id:      subMenu.Id,
				MenuId:  subMenu.MenuID,
				Name:    subMenu.Name,
				SideBar: subMenu.SideBar,
				Status:  subMenu.Status,
			})
		}

		configMenus = append(configMenus, &configmenuresponses.ConfigMenuResponse{
			Id:      menu.Id,
			Name:    menu.Name,
			Status:  menu.Status,
			SubMenu: subMenus,
			SideBar: menu.SideBar,
		})
	}

	return configMenus, nil
}

func (configMenuRepository *configMenuRepositoryImpl) CreateConfigMenu(req *configmenudto.ConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configMenu := &entities.ConfigMenu{
		Name:      req.Name,
		SideBar:   req.SideBar,
		Status:    req.Status,
		CompanyId: req.CompanyId,
	}

	result := configMenuRepository.DB.Table("Menu").Omit("id").Create(configMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Create Menu successful.",
	}, nil
}

func (configMenuRepository *configMenuRepositoryImpl) UpdateConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configMenu := &entities.ConfigMenu{
		Id:      req.Id,
		Name:    req.Name,
		SideBar: req.SideBar,
		Status:  req.Status,
	}

	result := configMenuRepository.DB.Table("Menu").Where("id = ?", configMenu.Id).Updates(configMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Update Menu successful.",
	}, nil
}

func (configMenuRepository *configMenuRepositoryImpl) DeleteConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configMenu := &entities.ConfigMenu{
		Id: req.Id,
	}

	result := configMenuRepository.DB.Table("Menu").Where("id = ?", configMenu.Id).Delete(configMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Delete Menu successful.",
	}, nil
}

// sub menu
func (configMenuRepository *configMenuRepositoryImpl) GetAllConfigSubMenu(companyId string) ([]*configmenuresponses.ConfigSubMenuResponse, error) {
	var configSubMenu []*configmenuresponses.ConfigSubMenuResponse

	// Use the correct placeholder '?' for the company_id condition
	result := configMenuRepository.DB.Table("Sub_menu").Where("company_id = ?", companyId).Find(&configSubMenu)

	if result.Error != nil {
		return nil, result.Error
	}

	return configSubMenu, nil

}

func (configMenuRepository *configMenuRepositoryImpl) CreateConfigSubMenu(req *configmenudto.ConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configSubMenu := &entities.ConfigSubMenu{
		MenuID:    req.MenuID,
		Name:      req.Name,
		Status:    req.Status,
		SideBar:   req.SideBar,
		CompanyId: req.CompanyId,
	}

	result := configMenuRepository.DB.Table("Sub_menu").Omit("id").Create(configSubMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Create Sub Menu successful.",
	}, nil
}

func (configMenuRepository *configMenuRepositoryImpl) UpdateConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configSubMenu := &entities.ConfigSubMenu{
		Id:      req.Id,
		MenuID:  req.MenuID,
		Name:    req.Name,
		SideBar: req.SideBar,
		Status:  req.Status,
	}

	result := configMenuRepository.DB.Table("sub_menu").Where("id = ?", configSubMenu.Id).Updates(configSubMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Update Sub Menu successful.",
	}, nil
}

func (configMenuRepository *configMenuRepositoryImpl) DeleteConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	configMenu := &entities.ConfigSubMenu{
		Id: req.Id,
	}

	result := configMenuRepository.DB.Table("Sub_menu").Where("id = ?", configMenu.Id).Delete(configMenu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &configmenuresponses.ConfigMenuCreateResponse{
		Success: true,
		Message: "Delete Sub Menu successful.",
	}, nil
}
