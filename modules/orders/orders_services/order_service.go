package orderservices

import (
	"github.com/natersland/b2b-e-commerce-api/configs"
	"github.com/natersland/b2b-e-commerce-api/modules/data/repositories"
	ordersdto "github.com/natersland/b2b-e-commerce-api/modules/orders/orders_dto"
	ordersresponses "github.com/natersland/b2b-e-commerce-api/modules/orders/orders_response"
)

type OrderServiceInterface interface {
	GetAllOrder() ([]*ordersresponses.OrderResponse, error)
	GetByCustomerId(customerId string) ([]*ordersresponses.OrderResponse, error)
	GetByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error)
	CreateOrder(req *ordersdto.OrderRequest) (*ordersresponses.OrderCreateResponse, error)
	UpdateOrderTracking(req *ordersdto.OrderTrackingRequest) (*ordersresponses.OrderCreateResponse, error)
	UpdateOrderPayment(req *ordersdto.OrderPaymentRequest) (*ordersresponses.OrderCreateResponse, error)
	DeleteByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error)
	AttachPackingListByOrderId(orderId string, req *ordersdto.UploadPackingListRequest) (*ordersresponses.OrderCreateResponse, error)
}

func OrderServiceImpl(cfg configs.ConfigInterface, ordersRepository repositories.OrdersRepositoryInterface) OrderServiceInterface {
	return &orderServiceImpl{
		config:           cfg,
		ordersRepository: ordersRepository,
	}
}

type orderServiceImpl struct {
	config           configs.ConfigInterface
	ordersRepository repositories.OrdersRepositoryInterface
}

func (c *orderServiceImpl) GetAllOrder() ([]*ordersresponses.OrderResponse, error) {
	result, err := c.ordersRepository.GetAllOrder()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *orderServiceImpl) GetByCustomerId(customerId string) ([]*ordersresponses.OrderResponse, error) {
	result, err := c.ordersRepository.GetByCustomerId(customerId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *orderServiceImpl) GetByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error) {
	result, err := c.ordersRepository.GetByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *orderServiceImpl) DeleteByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error) {
	result, err := c.ordersRepository.DeleteByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *orderServiceImpl) CreateOrder(req *ordersdto.OrderRequest) (*ordersresponses.OrderCreateResponse, error) {
	result, err := c.ordersRepository.CreateOrder(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *orderServiceImpl) AttachPackingListByOrderId(orderId string, req *ordersdto.UploadPackingListRequest) (*ordersresponses.OrderCreateResponse, error) {
	result, err := c.ordersRepository.AttachPackingListByOrderId(orderId, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *orderServiceImpl) UpdateOrderTracking(req *ordersdto.OrderTrackingRequest) (*ordersresponses.OrderCreateResponse, error) {
	result, err := c.ordersRepository.UpdateOrderTracking(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *orderServiceImpl) UpdateOrderPayment(req *ordersdto.OrderPaymentRequest) (*ordersresponses.OrderCreateResponse, error) {
	result, err := c.ordersRepository.UpdateOrderPayment(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
