package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Driver() DriverRepoI
	Car() CarRepoI
	Order() OrderRepoI
	Branch() BranchRepoI
}
type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type DriverRepoI interface {
	Create(context.Context, *models.CreateDriver) (string, error)
	GetByID(context.Context, *models.DriverPrimaryKey) (*models.Driver, error)
	GetList(context.Context, *models.DriverGetListRequest) (*models.DriverGetListResponse, error)
	Update(context.Context, *models.UpdateDriver) (int64, error)
	Delete(context.Context, *models.DriverPrimaryKey) error
}
type CarRepoI interface {
	Create(context.Context, *models.CreateCar) (string, error)
	GetByID(context.Context, *models.CarPrimaryKey) (*models.Car, error)
	GetList(context.Context, *models.CarGetListRequest) (*models.CarGetListResponse, error)
	Update(context.Context, *models.UpdateCar) (int64, error)
	Delete(context.Context, *models.CarPrimaryKey) error
}
type OrderRepoI interface {
	Create(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimaryKey) (*models.Order, error)
	GetList(context.Context, *models.OrderGetListRequest) (*models.OrderGetListResponse, error)
	Update(context.Context, *models.UpdateOrder) (int64, error)
	Delete(context.Context, *models.OrderPrimaryKey) error
}
type BranchRepoI interface {
	Create(context.Context, *models.CreateBranch) (string, error)
	GetByID(context.Context, *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(context.Context, *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(context.Context, *models.UpdateBranch) (int64, error)
	Delete(context.Context, *models.BranchPrimaryKey) error
}
