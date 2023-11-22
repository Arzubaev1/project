package models

type Car struct {
	Id          string `json:"id"`
	Model       string `json:"model"`
	Brand       string `json:"brand"`
	StateNumber string `json:"state_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CarPrimaryKey struct {
	Id string `json:"id"`
}
type CreateCar struct {
	Model       string `json:"model"`
	Brand       string `json:"brand"`
	StateNumber string `json:"state_number"`
}
type UpdateCar struct {
	Id          string `json:"id"`
	Model       string `json:"model"`
	Brand       string `json:"brand"`
	StateNumber string `json:"state_number"`
}
type CarGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type CarGetListResponse struct {
	Count int    `json:"count"`
	Cars  []*Car `json:"cars"`
}
