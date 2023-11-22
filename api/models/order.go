package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}
type Order struct {
	Id                      string `json:"id"`
	DriverId                string `json:"dircer_id"`
	Date                    string `json:"date"`
	Discount                string `json:"discount"`
	DiscountPrice           int64  `json:"discount_price"`
	BranchId                string `json:"branch_id"`
	Status                  string `json:"status"`
	FirstClientId           string `json:"first_client_id"`
	SecondClientId          string `json:"second_client_id"`
	FirstClientLocation     string `json:"first_client_location"`
	SecondClientLocation    string `json:"second_client_location"`
	FirstClientDestination  string `json:"first_client_destination"`
	SecondClientDestination string `json:"second_client_destination"`
	FirstClientMillage      int64  `json:"first_client_millage"`
	SecondClientMillage     int64  `json:"second_client_millage"`
	PriceForMillage         int64  `json:"price_for_millage"`
	FirstClientTotalPrice   int64  `json:"first_client_total_price"`
	SecondClientTotalPrice  int64  `json:"second_client_total_price"`
	PaymentType             string `json:"payment_type"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
}
type CreateOrder struct {
	DriverId                string `json:"dircer_id"`
	Date                    string `json:"date"`
	Discount                string `json:"discount"`
	DiscountPrice           int64  `json:"discount_price"`
	BranchId                string `json:"branch_id"`
	FirstClientId           string `json:"first_client_id"`
	SecondClientId          string `json:"second_client_id"`
	FirstClientLocation     string `json:"first_client_location"`
	SecondClientLocation    string `json:"second_client_location"`
	FirstClientDestination  string `json:"first_client_destination"`
	SecondClientDestination string `json:"second_client_destination"`
	FirstClientMillage      int64  `json:"first_client_millage"`
	SecondClientMillage     int64  `json:"second_client_millage"`
	PriceForMillage         int64  `json:"price_for_millage"`
	PaymentType             string `json:"payment_type"`
}
type UpdateOrder struct {
	Id                      string `json:"id"`
	DriverId                string `json:"dircer_id"`
	Date                    string `json:"date"`
	Discount                string `json:"discount"`
	DiscountPrice           int64  `json:"discount_price"`
	BranchId                string `json:"branch_id"`
	Status                  string `json:"status"`
	FirstClientId           string `json:"first_client_id"`
	SecondClientId          string `json:"second_client_id"`
	FirstClientLocation     string `json:"first_client_location"`
	SecondClientLocation    string `json:"second_client_location"`
	FirstClientDestination  string `json:"first_client_destination"`
	SecondClientDestination string `json:"second_client_destination"`
	FirstClientMillage      int64  `json:"first_client_millage"`
	SecondClientMillage     int64  `json:"second_client_millage"`
	PriceForMillage         int64  `json:"price_for_millage"`
	FirstClientTotalPrice   int64  `json:"first_client_total_price"`
	SecondClientTotalPrice  int64  `json:"second_client_total_price"`
	PaymentType             string `json:"payment_type"`
}
type OrderGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type OrderGetListResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}
