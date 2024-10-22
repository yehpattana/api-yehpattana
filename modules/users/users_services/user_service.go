package usersservices

import (
	"github.com/natersland/b2b-e-commerce-api/configs"
	"github.com/natersland/b2b-e-commerce-api/modules/data/repositories"
	usersdto "github.com/natersland/b2b-e-commerce-api/modules/users/users_dto"
	usersresponses "github.com/natersland/b2b-e-commerce-api/modules/users/users_responses"
)

type UserServiceInterface interface {
	GetAllCustomers(role string, companyId string) ([]*usersresponses.CustomerResponse, error)
	GetCustomerDetail(userId string) (*usersresponses.GetCustomerDetialResponse, error)
	GetUserIdByEmail(userEmail string) (*usersresponses.GetUserIdByEmailResponse, error)
	ChangePassword(userId string, newPassword string) (*usersresponses.ChangePasswordResponse, error)
	VerifyUser(userId string) (*usersresponses.VerifyUserResponse, error)
	BanUser(userId string) (*usersresponses.BanUserResponse, error)
	UpdateCustomer(userRequest *usersdto.UpdateCustomerRequest, userId string) (*usersresponses.UpdateUserResponse, error)
	DeleteCustomer(userId string, customerId string) (*usersresponses.BanUserResponse, error)
}

func UserServiceImpl(cfg configs.ConfigInterface, usersRepository repositories.UsersRepositoryInterface) UserServiceInterface {
	return &userServiceImpl{
		config:          cfg,
		usersRepository: usersRepository,
	}
}

type userServiceImpl struct {
	config          configs.ConfigInterface
	usersRepository repositories.UsersRepositoryInterface
}

func (u *userServiceImpl) GetAllCustomers(role string, companyId string) ([]*usersresponses.CustomerResponse, error) {
	result, err := u.usersRepository.FindAllCustomers(role, companyId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceImpl) GetCustomerDetail(userId string) (*usersresponses.GetCustomerDetialResponse, error) {
	result, err := u.usersRepository.GetCustomerDetail(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (u *userServiceImpl) GetUserIdByEmail(userEmail string) (*usersresponses.GetUserIdByEmailResponse, error) {
	result, err := u.usersRepository.GetUserIdByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceImpl) ChangePassword(userId string, newPassword string) (*usersresponses.ChangePasswordResponse, error) {
	result, err := u.usersRepository.ChangePassword(userId, newPassword)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceImpl) VerifyUser(userId string) (*usersresponses.VerifyUserResponse, error) {
	result, err := u.usersRepository.VerifyUser(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceImpl) BanUser(userId string) (*usersresponses.BanUserResponse, error) {
	result, err := u.usersRepository.BanUser(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceImpl) UpdateCustomer(userRequest *usersdto.UpdateCustomerRequest, userId string) (*usersresponses.UpdateUserResponse, error) {
	result, err := u.usersRepository.UpdateCustomer(userRequest, userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (u *userServiceImpl) DeleteCustomer(userId string, customerId string) (*usersresponses.BanUserResponse, error) {
	result, err := u.usersRepository.DeleteCustomer(userId, customerId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
