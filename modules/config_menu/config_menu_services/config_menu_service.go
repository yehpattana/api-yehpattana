package configmenuservices

import (
	"github.com/yehpattana/api-yehpattana/configs"
	configmenudto "github.com/yehpattana/api-yehpattana/modules/config_menu/config_menu_dto"
	configmenuresponses "github.com/yehpattana/api-yehpattana/modules/config_menu/config_menu_response"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
)

type ConfigMenuServiceInterface interface {
	GetAllConfigMenu(companyId string) ([]*configmenuresponses.ConfigMenuResponse, error)
	CreateConfigMenu(req *configmenudto.ConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	UpdateConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	DeleteConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	GetAllConfigSubMenu(companyId string) ([]*configmenuresponses.ConfigSubMenuResponse, error)
	CreateConfigSubMenu(req *configmenudto.ConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	UpdateConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
	DeleteConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error)
}

func ConfigMenuServiceImpl(cfg configs.ConfigInterface, configMenuRepository repositories.ConfigMenuRepositoryInterface) ConfigMenuServiceInterface {
	return &configMenuServiceImpl{
		config:               cfg,
		configMenuRepository: configMenuRepository,
	}
}

type configMenuServiceImpl struct {
	config               configs.ConfigInterface
	configMenuRepository repositories.ConfigMenuRepositoryInterface
}

func (c *configMenuServiceImpl) GetAllConfigMenu(companyId string) ([]*configmenuresponses.ConfigMenuResponse, error) {
	result, err := c.configMenuRepository.GetAllConfigMenu(companyId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *configMenuServiceImpl) CreateConfigMenu(req *configmenudto.ConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.CreateConfigMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *configMenuServiceImpl) UpdateConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.UpdateConfigMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *configMenuServiceImpl) DeleteConfigMenu(req *configmenudto.UpdateConfigMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.DeleteConfigMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// sub menu
func (c *configMenuServiceImpl) GetAllConfigSubMenu(companyId string) ([]*configmenuresponses.ConfigSubMenuResponse, error) {
	result, err := c.configMenuRepository.GetAllConfigSubMenu(companyId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *configMenuServiceImpl) CreateConfigSubMenu(req *configmenudto.ConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.CreateConfigSubMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *configMenuServiceImpl) UpdateConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.UpdateConfigSubMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *configMenuServiceImpl) DeleteConfigSubMenu(req *configmenudto.UpdateConfigSubMenuRequest) (*configmenuresponses.ConfigMenuCreateResponse, error) {
	result, err := c.configMenuRepository.DeleteConfigSubMenu(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
