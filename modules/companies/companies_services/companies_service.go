package companiesservices

import (
	"github.com/natersland/b2b-e-commerce-api/configs"
	companiesdto "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_dto"
	companiesresponses "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_responses"
	"github.com/natersland/b2b-e-commerce-api/modules/data/repositories"
)

type CompaniesServiceInterface interface {
	GetAllCompanies(role string, companyId string) (*companiesresponses.GetCompaniesResponse, error)
	GetCompanyById(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error)
	CreateCompany(req *companiesdto.CreateCompanyRequest) (*companiesresponses.CreateCompanyResponse, error)
	EditCompanyById(req *companiesdto.EditCompanyRequest) (*companiesresponses.EditCompanyResponse, error)
	DeleteCompany(req *companiesdto.DeleteCompanyRequest) (*companiesresponses.DeleteCompanyResponse, error)
}

func CompaniesServiceImpl(cfg configs.ConfigInterface, companiesRepository repositories.CompaniesRepositoryInterface) CompaniesServiceInterface {
	return &companiesServiceImpl{
		config:              cfg,
		companiesRepository: companiesRepository,
	}
}

type companiesServiceImpl struct {
	config              configs.ConfigInterface
	companiesRepository repositories.CompaniesRepositoryInterface
}

func (c *companiesServiceImpl) GetAllCompanies(role string, companyId string) (*companiesresponses.GetCompaniesResponse, error) {
	result, err := c.companiesRepository.GetAllCompanies(role, companyId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *companiesServiceImpl) GetCompanyById(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error) {
	result, err := c.companiesRepository.GetCompanyById(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *companiesServiceImpl) CreateCompany(req *companiesdto.CreateCompanyRequest) (*companiesresponses.CreateCompanyResponse, error) {
	result, err := c.companiesRepository.CreateCompany(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *companiesServiceImpl) EditCompanyById(req *companiesdto.EditCompanyRequest) (*companiesresponses.EditCompanyResponse, error) {
	result, err := c.companiesRepository.EditCompanyById(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *companiesServiceImpl) DeleteCompany(req *companiesdto.DeleteCompanyRequest) (*companiesresponses.DeleteCompanyResponse, error) {
	result, err := c.companiesRepository.DeleteCompany(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
