package model

import "github.com/ajitirto/go-restaurant-app/internal/model/constant"

type Order struct {
	ID            string               `gorm:"primaryKey" json:"id"`
	Status        constant.OrderStatus `json:"status"`
	ProductOrders []ProductOrder       `json:"product_orders"`
	ReferenceID   string               `gorm:"unique" json:"reference_id"`
}

type ProductOrder struct {
	ID         string                      `gorm:"primaryKey" json:"id"`
	OrderId    string                      `json:"order_id"`
	OrderCode  string                      `json:"order_code"`
	Quantity   int                         `json:"quantity"`
	TotalPrice int64                       `json:"total_price"`
	Status     constant.ProductOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	OrderID string `json:"order_id"`
}
