package models

type CartResponse struct {
	TotalPrice float64
	Cart       []Cart
}
type CartTotal struct {
	TotalPrice     float64 `json:"total_price"`
	FinalPrice     float64 `json:"final_price"`
}
type Cart struct {
	ProductID   uint    `json:"product_id"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type Carts struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}
type GetCartResponse struct {
	ID   int
	Data []GetCart
}
type GetCart struct {
	ID          int     `json:"id" gorm:"primarykey;not null"`
	ProductId   int     `json:"product_id"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}
