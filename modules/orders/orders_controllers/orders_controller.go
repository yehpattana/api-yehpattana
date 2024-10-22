package orderscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	ordersdto "github.com/yehpattana/api-yehpattana/modules/orders/orders_dto"
	orderservices "github.com/yehpattana/api-yehpattana/modules/orders/orders_services"
)

type OrdersControllerInterface interface {
	GetAllOrder(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrderTracking(c *fiber.Ctx) error
	UpdateOrderPayment(c *fiber.Ctx) error
	GetByCustomerId(c *fiber.Ctx) error
	GetByOrderId(c *fiber.Ctx) error
	DeleteByOrderId(c *fiber.Ctx) error
	AttachPackingListByOrderId(c *fiber.Ctx) error
}

type ordersControllerImpl struct {
	config        configs.ConfigInterface
	orderServices orderservices.OrderServiceInterface
}

func OrdersControllerImpl(config configs.ConfigInterface, orderservices orderservices.OrderServiceInterface) OrdersControllerInterface {
	return &ordersControllerImpl{
		config:        config,
		orderServices: orderservices,
	}
}

func (config *ordersControllerImpl) GetAllOrder(c *fiber.Ctx) error {
	result, err := config.orderServices.GetAllOrder()
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllOrder", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *ordersControllerImpl) CreateOrder(c *fiber.Ctx) error {
	req := new(ordersdto.OrderRequest)

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "CreateOrder", err.Error()).Res()
	}
	result, err := config.orderServices.CreateOrder(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "CreateOrder", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *ordersControllerImpl) UpdateOrderTracking(c *fiber.Ctx) error {
	req := new(ordersdto.OrderTrackingRequest)

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "UpdateOrderTracking", err.Error()).Res()
	}
	result, err := config.orderServices.UpdateOrderTracking(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "UpdateOrderTracking", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *ordersControllerImpl) GetByCustomerId(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	result, err := config.orderServices.GetByCustomerId(customerId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetByCustomerId", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
func (config *ordersControllerImpl) GetByOrderId(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	result, err := config.orderServices.GetByOrderId(orderId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetByCustomerId", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *ordersControllerImpl) DeleteByOrderId(c *fiber.Ctx) error {
	customerId := c.Params("orderId")
	result, err := config.orderServices.DeleteByOrderId(customerId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "DeleteByOrderId", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
func (config *ordersControllerImpl) AttachPackingListByOrderId(c *fiber.Ctx) error {
	req := new(ordersdto.UploadPackingListRequest)

	file, err := c.FormFile("file")
	if err != nil {
		file = nil
	}

	req.File = file
	orderId := c.Params("orderId")
	result, err := config.orderServices.AttachPackingListByOrderId(orderId, req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "AttachPackingListByOrderId", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *ordersControllerImpl) UpdateOrderPayment(c *fiber.Ctx) error {
	req := new(ordersdto.OrderPaymentRequest)

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "UpdateOrderTracking", err.Error()).Res()
	}
	result, err := config.orderServices.UpdateOrderPayment(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "UpdateOrderTracking", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
