package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	commonfolderpath "github.com/yehpattana/api-yehpattana/modules/commons/common_folder_path"
	commonhelpers "github.com/yehpattana/api-yehpattana/modules/commons/common_helpers"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	ordersdto "github.com/yehpattana/api-yehpattana/modules/orders/orders_dto"
	ordersresponses "github.com/yehpattana/api-yehpattana/modules/orders/orders_response"
	productsdto "github.com/yehpattana/api-yehpattana/modules/products/products_dto"
	"gorm.io/gorm"
)

type OrdersRepositoryInterface interface {
	GetAllOrder() ([]*ordersresponses.OrderResponse, error)
	GetByCustomerId(customerId string) ([]*ordersresponses.OrderResponse, error)
	GetByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error)
	DeleteByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error)
	// GetAllOrders() ([]*ordersresponses.OrderResponse, error)
	CreateOrder(req *ordersdto.OrderRequest) (*ordersresponses.OrderCreateResponse, error)
	UpdateOrderTracking(req *ordersdto.OrderTrackingRequest) (*ordersresponses.OrderCreateResponse, error)
	UpdateOrderPayment(req *ordersdto.OrderPaymentRequest) (*ordersresponses.OrderCreateResponse, error)
	AttachPackingListByOrderId(orderId string, req *ordersdto.UploadPackingListRequest) (*ordersresponses.OrderCreateResponse, error)
}

type ordersRepositoryImpl struct {
	DB *gorm.DB
}

func OrdersRepositoryImpl(db *gorm.DB) OrdersRepositoryInterface {
	return &ordersRepositoryImpl{
		DB: db,
	}
}

func (ordersRepository *ordersRepositoryImpl) GetAllOrder() ([]*ordersresponses.OrderResponse, error) {
	var orders []entities.Order
	var orderResponses []*ordersresponses.OrderResponse

	result := ordersRepository.DB.Preload("CustomerDetail").Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, order := range orders {
		orderResponses = append(orderResponses, &ordersresponses.OrderResponse{
			OrderNo:         order.OrderNo,
			OrderID:         order.OrderID,
			ShippingAddress: order.ShippingAddress,
			TrackingNo:      order.TrackingNo,
			PackingList:     order.PackingList,
			OrderDetail:     string(order.OrderDetail), // Convert byte slice to string
			TotalAmount:     order.TotalAmount,
			CustomerDetail: ordersresponses.CustomerDetail{
				ContractName: order.CustomerDetail.ContactName,
				CustomerID:   order.CustomerDetail.CustomerID,
				CompanyName:  order.CustomerDetail.CompanyName,
			},
			Status:        order.Status,
			PaymentStatus: order.PaymentStatus,
			PaymentID:     order.PaymentID,
			CreatedAt:     order.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     order.UpdatedAt.Format(time.RFC3339),
		})
	}

	return orderResponses, nil

}

func (ordersRepository *ordersRepositoryImpl) GetByCustomerId(customerId string) ([]*ordersresponses.OrderResponse, error) {
	var orders []entities.Order
	var orderResponses []*ordersresponses.OrderResponse

	result := ordersRepository.DB.Preload("CustomerDetail").Where("customer_id=?", customerId).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, order := range orders {
		orderResponses = append(orderResponses, &ordersresponses.OrderResponse{
			OrderNo:         order.OrderNo,
			OrderID:         order.OrderID,
			ShippingAddress: order.ShippingAddress,
			TrackingNo:      order.TrackingNo,
			PackingList:     order.PackingList,
			OrderDetail:     string(order.OrderDetail),
			TotalAmount:     order.TotalAmount,
			CustomerDetail: ordersresponses.CustomerDetail{
				ContractName: order.CustomerDetail.ContactName,
				CustomerID:   order.CustomerDetail.CustomerID,
				CompanyName:  order.CustomerDetail.CompanyName,
			},
			Status:        order.Status,
			PaymentStatus: order.PaymentStatus,
			PaymentID:     order.PaymentID,
			CreatedAt:     order.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     order.UpdatedAt.Format(time.RFC3339),
		})
	}

	return orderResponses, nil
}

func (ordersRepository *ordersRepositoryImpl) GetByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error) {
	var orders []entities.Order
	var orderResponses []*ordersresponses.OrderResponse

	result := ordersRepository.DB.Preload("CustomerDetail").Where("order_id=?", orderId).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, order := range orders {
		orderResponses = append(orderResponses, &ordersresponses.OrderResponse{
			OrderNo:         order.OrderNo,
			OrderID:         order.OrderID,
			ShippingAddress: order.ShippingAddress,
			TrackingNo:      order.TrackingNo,
			PackingList:     order.PackingList,
			OrderDetail:     string(order.OrderDetail),
			TotalAmount:     order.TotalAmount,
			CustomerDetail: ordersresponses.CustomerDetail{
				ContractName: order.CustomerDetail.ContactName,
				CustomerID:   order.CustomerDetail.CustomerID,
				CompanyName:  order.CustomerDetail.CompanyName,
			},
			Status:        order.Status,
			PaymentStatus: order.PaymentStatus,
			PaymentID:     order.PaymentID,
			CreatedAt:     order.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     order.UpdatedAt.Format(time.RFC3339),
		})
	}
	return orderResponses, nil
}

func (ordersRepository *ordersRepositoryImpl) DeleteByOrderId(orderId string) ([]*ordersresponses.OrderResponse, error) {
	var orderResponses []*ordersresponses.OrderResponse

	result := ordersRepository.DB.Delete(&entities.Order{}, orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderResponses, nil
}

func (ordersRepository *ordersRepositoryImpl) getNextSequence() (int, error) {
	var nextSeq int
	err := ordersRepository.DB.Raw("SELECT NEXT VALUE FOR dbo.OrderSeq").Scan(&nextSeq).Error
	if err != nil {
		return 0, err
	}
	return nextSeq, nil
}

func (ordersRepository *ordersRepositoryImpl) CreateOrder(req *ordersdto.OrderRequest) (*ordersresponses.OrderCreateResponse, error) {
	// Generate the order_no
	currentDate := time.Now()
	year := currentDate.Format("06")       // Last two digits of the year
	date := currentDate.Format("20060102") // YYYYMMDD format

	sequence, err := ordersRepository.getNextSequence()
	if err != nil {
		return nil, err
	}
	var products []productsdto.Product
	err = json.Unmarshal([]byte(req.OrderDetail), &products)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	// Loop through products to get the first product variant
	for _, product := range products {
		for _, stock := range product.ProductVariants[0].Stock {
			// Loop through products again to check sizes and quantities
			for size, quantity := range product.Quantity {
				// Check if stock size matches product quantity size and quantity is more than 0
				if stock.Size == size && quantity > 0 {
					err = ordersRepository.DecreaseStock(stock.ID, quantity)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	fmt.Printf("%+v\n", products)

	orderNo := fmt.Sprintf("%s%s%s%03d", req.CompanyName, date, year, sequence)

	// Convert the JSON string to a byte slice
	orderDetailBytes := []byte(req.OrderDetail)
	order := &entities.Order{
		OrderNo:         orderNo,
		OrderDetail:     orderDetailBytes,
		ShippingAddress: req.ShippingAddress,
		TrackingNo:      req.TrackingNo,
		CustomerID:      req.CustomerID,
		TotalAmount:     req.TotalAmount,
		Status:          "pending",
		PaymentStatus:   "pending",
	}

	result := ordersRepository.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ordersresponses.OrderCreateResponse{
		Success: true,
		Message: "Create Order successful.",
	}, nil
}

func (ordersRepository *ordersRepositoryImpl) UpdateOrderTracking(req *ordersdto.OrderTrackingRequest) (*ordersresponses.OrderCreateResponse, error) {

	// Update only the Status and TrackingNo fields where OrderNo matches
	result := ordersRepository.DB.Model(&entities.Order{}).Where("order_no = ?", req.OrderNo).Updates(map[string]interface{}{
		"tracking_no": req.TrackingNo,
		"status":      req.Status,
	})

	if result.Error != nil {
		return nil, result.Error
	}
	return &ordersresponses.OrderCreateResponse{
		Success: true,
		Message: "Order update successful.",
	}, nil
}

func (ordersRepository *ordersRepositoryImpl) AttachPackingListByOrderId(orderId string, req *ordersdto.UploadPackingListRequest) (*ordersresponses.OrderCreateResponse, error) {
	productFolder := commonfolderpath.ProductFolderPath
	attachment, err := commonhelpers.UploadAttachment(req.File, productFolder)
	if err != nil {
		return &ordersresponses.OrderCreateResponse{
			Success: true,
			Message: "Order update attachment error.",
		}, err
	}
	// Update only the Status and TrackingNo fields where OrderNo matches
	result := ordersRepository.DB.Model(&entities.Order{}).Where("order_id = ?", orderId).Updates(map[string]interface{}{
		"packing_list": attachment,
	})

	if result.Error != nil {
		return nil, result.Error
	}
	return &ordersresponses.OrderCreateResponse{
		Success: true,
		Message: "Order update attachment successful.",
	}, nil
}

// Assume these functions are defined elsewhere in your codebase
func commonhelpersGetCurrentTimeISO() string {
	return time.Now().Format(time.RFC3339)
}

func assignVariantStatus(isOutOfStock bool) string {
	if isOutOfStock {
		return "out_of_stock"
	}
	return "available"
}

func checkOutOfStock(quantity int) bool {
	return quantity <= 0
}

// decreaseStock decreases the quantity of stock with a given stock ID.
func (ordersRepository *ordersRepositoryImpl) DecreaseStock(stockId string, decreaseQuantity int) error {
	// find stock by stock id and assign to stock object
	stock := &entities.Stock{}
	result := ordersRepository.DB.Table("Stock").Where("id = ?", stockId).First(stock)
	if result.Error != nil {
		return result.Error
	}
	// Calculate updated stock values
	newQuantity := stock.Quantity - decreaseQuantity
	if decreaseQuantity >= stock.Quantity {
		newQuantity = 0
	}

	newPreQuantity := stock.PreQuantity
	if decreaseQuantity > stock.Quantity {
		newPreQuantity = stock.PreQuantity - (decreaseQuantity - stock.Quantity)
	}
	// Prepare update data
	updates := map[string]interface{}{
		"quantity":     newQuantity,
		"pre_quantity": newPreQuantity,
		"item_status":  assignVariantStatus(checkOutOfStock(newQuantity)),
		"updated_at":   commonhelpersGetCurrentTimeISO(),
	}
	// Update the stock record in the database
	result = ordersRepository.DB.Table("Stock").Where("id = ?", stockId).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	// If stockId not found
	return nil
}

func (ordersRepository *ordersRepositoryImpl) UpdateOrderPayment(req *ordersdto.OrderPaymentRequest) (*ordersresponses.OrderCreateResponse, error) {

	// Update only the Status and TrackingNo fields where OrderNo matches
	result := ordersRepository.DB.Model(&entities.Order{}).Where("order_no = ?", req.OrderNo).Updates(map[string]interface{}{
		"payment_status": req.PaymentStatus,
		"payment_id":     req.PaymentID,
	})

	if result.Error != nil {
		return nil, result.Error
	}
	return &ordersresponses.OrderCreateResponse{
		Success: true,
		Message: "Order update payment status successful.",
	}, nil
}
