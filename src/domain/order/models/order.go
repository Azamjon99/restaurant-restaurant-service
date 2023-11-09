package models

import "time"

type PaymentInfo struct {
	Method      string `json:"method"` // cash, card
	DeliveryFee int    `json:"delivery_fee"`
	TotalPrice  int    `json:"total_price"`
	CardID      string `json:"card_id"` // required if method == card
}
type DeliveryInfo struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type EaterInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type OrderItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}
type Order struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	RestaurantID string        `json:"restaurant_id" gorm:"index:idx_restaurant_order"` // restaurant id
	Status       string        `json:"status"`
	Eater        *EaterInfo    `json:"eater" gorm:"serializer:json"`
	Delivery     *DeliveryInfo `json:"delivery" gorm:"serializer:json"` // address info
	Payment      *PaymentInfo  `json:"payment" gorm:"serializer:json"`
	Items        []*OrderItem  `json:"items" gorm:"serializer:json"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
