package repositories

import (
	"fmt"

	commonfolderpath "github.com/natersland/b2b-e-commerce-api/modules/commons/common_folder_path"
	commonhelpers "github.com/natersland/b2b-e-commerce-api/modules/commons/common_helpers"
	commonimages "github.com/natersland/b2b-e-commerce-api/modules/commons/common_images"
	companiesdto "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_dto"
	companiesresponses "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_responses"
	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	"gorm.io/gorm"
)

type CompaniesRepositoryInterface interface {
	GetAllCompanies(role string, companyId string) (*companiesresponses.GetCompaniesResponse, error)
	GetCompanyById(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error)
	GetCompanyByName(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error)
	CreateCompany(req *companiesdto.CreateCompanyRequest) (*companiesresponses.CreateCompanyResponse, error)
	EditCompanyById(req *companiesdto.EditCompanyRequest) (*companiesresponses.EditCompanyResponse, error)
	DeleteCompany(req *companiesdto.DeleteCompanyRequest) (*companiesresponses.DeleteCompanyResponse, error)
}

func CompaniesRepositoryImpl(db *gorm.DB) CompaniesRepositoryInterface {
	return &companiesRepositoryImpl{
		DB: db,
	}
}

type companiesRepositoryImpl struct {
	*gorm.DB
}

func (c *companiesRepositoryImpl) GetAllCompanies(role string, companyId string) (*companiesresponses.GetCompaniesResponse, error) {
	var companies []*entities.Company
	var result *gorm.DB
	fmt.Println(role)
	if role == "SuperAdmin" {
		result = c.DB.Table("Companies").Order("company_name ASC").Find(&companies)
	} else {
		// filter if same master code then not appear in the list
		result = c.DB.Table("Companies").Where("id = ?", companyId).First(&companies)
	}
	// Adding Order clause to sort by company_name alphabetically

	if result.Error != nil {
		return &companiesresponses.GetCompaniesResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Failed to fetch companies: " + result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	var companyFields []*companiesresponses.CompanyFieldResponse
	for _, company := range companies {
		companyFields = append(companyFields, &companiesresponses.CompanyFieldResponse{
			Id:                       company.Id,
			CompanyCode:              company.CompanyCode,
			CompanyName:              company.CompanyName,
			Currency:                 company.Currency,
			MinimumCostAvoidShipping: company.MinimumCostAvoidShipping,
			Logo:                     company.Logo,
		})
	}

	return &companiesresponses.GetCompaniesResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Companies fetched successfully.",
		},
		Data: companyFields,
	}, nil
}

func (c *companiesRepositoryImpl) GetCompanyById(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error) {
	var company entities.Company
	result := c.DB.Table("Companies").Where("id = ?", req.CompanyId).First(&company)
	if result.Error != nil {
		return &companiesresponses.GetCompanyDetailResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Failed to fetch company.",
			},
			Data: nil,
		}, result.Error
	}

	return &companiesresponses.GetCompanyDetailResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Company fetched successfully.",
		},
		Data: &companiesresponses.CompanyFieldResponse{
			Id:                       company.Id,
			CompanyCode:              company.CompanyCode,
			CompanyName:              company.CompanyName,
			Currency:                 company.Currency,
			MinimumCostAvoidShipping: company.MinimumCostAvoidShipping,
			Logo:                     company.Logo,
		},
	}, nil
}
func (c *companiesRepositoryImpl) GetCompanyByName(req *companiesdto.GetCompanyDetailRequest) (*companiesresponses.GetCompanyDetailResponse, error) {
	var company entities.Company
	result := c.DB.Table("Companies").Where("company_name = ?", req.CompanyId).First(&company)
	if result.Error != nil {
		return &companiesresponses.GetCompanyDetailResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Failed to fetch company." + result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	return &companiesresponses.GetCompanyDetailResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Company fetched successfully.",
		},
		Data: &companiesresponses.CompanyFieldResponse{
			Id:                       company.Id,
			CompanyCode:              company.CompanyCode,
			CompanyName:              company.CompanyName,
			Currency:                 company.Currency,
			MinimumCostAvoidShipping: company.MinimumCostAvoidShipping,
			Logo:                     company.Logo,
		},
	}, nil
}

func (c *companiesRepositoryImpl) CreateCompany(req *companiesdto.CreateCompanyRequest) (*companiesresponses.CreateCompanyResponse, error) {

	defaultLogoImage := commonimages.DefaultCompanyLogoImage
	logoImageFolder := commonfolderpath.CompanyLogoFolderPath
	logoImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.Logo, defaultLogoImage, logoImageFolder)
	if err != nil {
		return &companiesresponses.CreateCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company registration failed." + err.Error(),
			},
			Data: nil,
		}, err
	}

	company := &entities.Company{
		Id:                       commonhelpers.GenerateUUID(),
		CompanyCode:              req.CompanyCode,
		CompanyName:              req.CompanyName,
		Currency:                 req.Currency,
		MinimumCostAvoidShipping: req.MinimumCostAvoidShipping,
		Logo:                     logoImageUrl,
	}

	// Save the Company to the database
	result := c.DB.Table("Companies").Create(company)
	if result.Error != nil {
		return &companiesresponses.CreateCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company registration failed." + result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	return &companiesresponses.CreateCompanyResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Company registration successful.",
		},
		Data: &companiesresponses.CompanyFieldResponse{
			Id:                       company.Id,
			CompanyCode:              company.CompanyCode,
			CompanyName:              company.CompanyName,
			Currency:                 company.Currency,
			MinimumCostAvoidShipping: company.MinimumCostAvoidShipping,
			Logo:                     company.Logo,
		},
	}, nil
}

func (c *companiesRepositoryImpl) EditCompanyById(req *companiesdto.EditCompanyRequest) (*companiesresponses.EditCompanyResponse, error) {
	// fetch company by id
	var currentCompanyData entities.Company
	currentCompanyResult := c.DB.Table("Companies").Where("id = ?", req.CompanyId).First(&currentCompanyData)
	if currentCompanyResult.Error != nil {
		return &companiesresponses.EditCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company update failed." + currentCompanyResult.Error.Error(),
			},
			Data: nil,
		}, currentCompanyResult.Error
	}

	isCompanyCodeIsBlank := req.CompanyCode == ""
	if isCompanyCodeIsBlank {
		req.CompanyCode = currentCompanyData.CompanyCode
	}

	isCompanyNameIsBlank := req.CompanyName == ""
	if isCompanyNameIsBlank {
		req.CompanyName = currentCompanyData.CompanyName
	}

	isCompanyLogoIsBlank := req.Logo == nil
	currentLogoImageUrl := currentCompanyData.Logo
	defaultLogoImage := commonimages.DefaultCompanyLogoImage
	logoImageFolder := commonfolderpath.CompanyLogoFolderPath
	logoImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.Logo, defaultLogoImage, logoImageFolder)
	if err != nil {
		return &companiesresponses.EditCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company update failed." + err.Error(),
			},
			Data: nil,
		}, err
	}

	if isCompanyLogoIsBlank {
		logoImageUrl = currentLogoImageUrl
	}

	company := &entities.Company{
		Id:                       req.CompanyId,
		CompanyCode:              req.CompanyCode,
		CompanyName:              req.CompanyName,
		Currency:                 req.Currency,
		MinimumCostAvoidShipping: req.MinimumCostAvoidShipping,
		Logo:                     logoImageUrl,
	}

	// Save the Company to the database
	result := c.DB.Table("Companies").Where("id = ?", req.CompanyId).Updates(company)
	if result.Error != nil {
		return &companiesresponses.EditCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company update failed." + result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	return &companiesresponses.EditCompanyResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Company update successful.",
		},
		Data: &companiesresponses.CompanyFieldResponse{
			Id:                       company.Id,
			CompanyCode:              company.CompanyCode,
			CompanyName:              company.CompanyName,
			Currency:                 company.Currency,
			MinimumCostAvoidShipping: company.MinimumCostAvoidShipping,
			Logo:                     company.Logo,
		},
	}, nil
}

func (c *companiesRepositoryImpl) DeleteCompany(req *companiesdto.DeleteCompanyRequest) (*companiesresponses.DeleteCompanyResponse, error) {
	// Delete the Company from the database
	resultDeleteSubmenu := c.DB.Table("Sub_Menu").Where("company_id = ?", req.CompanyId).Delete(&entities.ConfigSubMenu{})
	if resultDeleteSubmenu.Error != nil {
		return &companiesresponses.DeleteCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company deletion failed." + resultDeleteSubmenu.Error.Error(),
			},
		}, resultDeleteSubmenu.Error
	}
	resultDeleteMenu := c.DB.Table("Menu").Where("company_id = ?", req.CompanyId).Delete(&entities.ConfigMenu{})
	if resultDeleteMenu.Error != nil {
		return &companiesresponses.DeleteCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company deletion failed." + resultDeleteMenu.Error.Error(),
			},
		}, resultDeleteMenu.Error
	}
	result := c.DB.Table("Companies").Where("id = ?", req.CompanyId).Delete(&entities.Company{})
	if result.Error != nil {
		return &companiesresponses.DeleteCompanyResponse{
			Status: &companiesresponses.CompanyStatusResponse{
				Success: false,
				Message: "Company deletion failed." + result.Error.Error(),
			},
		}, result.Error
	}

	return &companiesresponses.DeleteCompanyResponse{
		Status: &companiesresponses.CompanyStatusResponse{
			Success: true,
			Message: "Company deletion successful.",
		},
	}, nil
}
