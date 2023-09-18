package domain

type Order struct {
	Order_uid          string   `json:"order_uid" validate:"required"`
	Track_number       string   `json:"track_number" validate:"required"`
	Entry              string   `json:"entry" validate:"required"`
	Locale             string   `json:"locale" validate:"required"`
	Internal_signature string   `json:"internal_signature"`
	Customer_id        string   `json:"customer_id" validate:"required"`
	Delivery_service   string   `json:"delivery_service" validate:"required"`
	Shardkey           string   `json:"shardkey" validate:"required"`
	Sm_id              int      `json:"sm_id" validate:"required"`
	Date_created       string   `json:"date_created" validate:"required"`
	Oof_shard          string   `json:"oof_shard" validate:"required"`
	Delivery           Delivery `json:"delivery" validate:"required"`
	Payment            Payment  `json:"payment" validate:"required"`
	Items              []Item   `json:"items" validate:"required"`
}
type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id" validate:"required"`
	Track_number string `json:"track_number" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	Rid          string `json:"rid" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Sale         int    `json:"sale" validate:"required"`
	Size         string `json:"size" `
	Total_price  int    `json:"total_price" validate:"required"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

type Payment struct {
	Transaction   string `json:"transaction" validate:"required"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency" validate:"required"`
	Provider      string `json:"provider" validate:"required"`
	Amount        int    `json:"amount" validate:"required"`
	Paymrnt_dt    int    `json:"paymrnt_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}
