package models

type Branch struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type BranchPrimaryKey struct {
	Id string `json:"id"`
}
type CreateBranch struct {
	Name string `json:"name"`
}
type UpdateBranch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type BranchGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type BranchGetListResponse struct {
	Count   int       `json:"count"`
	Branchs []*Branch `json:"branchs"`
}
