package models

type Driver struct {
	Id          string `json:"id"`
	FirsName    string `json:"first_name"`
	LastName    string `json:"last_name"`
	CarId       string `json:"car_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BranchId    string `json:"branch_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type DriverPrimaryKey struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
type CreateDriver struct {
	FirsName    string `json:"first_name"`
	LastName    string `json:"last_name"`
	CarId       string `json:"car_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BranchId    string `json:"branch_id"`
}
type UpdateDriver struct {
	Id          string `json:"id"`
	FirsName    string `json:"first_name"`
	LastName    string `json:"last_name"`
	CarId       string `json:"car_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BranchId    string `json:"branch_id"`
}

type DriverGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type DriverGetListResponse struct {
	Count   int       `json:"count"`
	Drivers []*Driver `json:"drivers"`
}
