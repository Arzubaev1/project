package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type store struct {
	db     *pgxpool.Pool
	user   *userRepo
	driver *driverRepo
	car    *carRepo
	order  *orderRepo
	branch *branchRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}
func (s *store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
func (s *store) Driver() storage.DriverRepoI {
	if s.driver == nil {
		s.driver = NewDriverRepo(s.db)
	}
	return s.driver
}
func (s *store) Car() storage.CarRepoI {
	if s.car == nil {
		s.car = NewCarRepo(s.db)
	}
	return s.car
}
func (s *store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}
	return s.order
}
func (s *store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}
