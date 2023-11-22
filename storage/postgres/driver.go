package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type driverRepo struct {
	db *pgxpool.Pool
}

func NewDriverRepo(db *pgxpool.Pool) *driverRepo {
	return &driverRepo{
		db: db,
	}
}

func (r *driverRepo) Create(ctx context.Context, req *models.CreateDriver) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO driver(id, first_name, last_name, car_id, phone_number, email, password, branch_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, ,$7, $8, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirsName,
		req.LastName,
		req.CarId,
		req.PhoneNumber,
		req.Email,
		req.Password,
		req.BranchId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *driverRepo) GetByID(ctx context.Context, req *models.DriverPrimaryKey) (*models.Driver, error) {

	var whereField = "id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	} else if len(req.PhoneNumber) > 0 {
		whereField = "phone_number"
		req.Id = req.PhoneNumber
	}

	var (
		query string

		id           sql.NullString
		first_name   sql.NullString
		last_name    sql.NullString
		car_id       sql.NullString
		phone_number sql.NullString
		email        sql.NullString
		password     sql.NullString
		branch_id    sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query = `
		SELECT
			id,
			first_name,
			last_name,
			car_id,
			phone_number,
			email,
			password,
			branch_id,
			created_at,
			updated_at
		FROM driver
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&last_name,
		&car_id,
		&phone_number,
		&email,
		&password,
		&branch_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Driver{
		Id:          id.String,
		FirsName:    first_name.String,
		LastName:    last_name.String,
		CarId:       car_id.String,
		PhoneNumber: phone_number.String,
		Email:       email.String,
		Password:    password.String,
		BranchId:    branch_id.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *driverRepo) GetList(ctx context.Context, req *models.DriverGetListRequest) (*models.DriverGetListResponse, error) {

	var (
		resp   = &models.DriverGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			car_id,
			phone_number,
			email,
			password,
			branch_id,
			created_at,
			updated_at
		FROM Drivers
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND first_name ILIKE '%' || '` + req.Search + `' || '%'`
	}
	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id           sql.NullString
			first_name   sql.NullString
			last_name    sql.NullString
			car_id       sql.NullString
			phone_number sql.NullString
			email        sql.NullString
			password     sql.NullString
			branch_id    sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&car_id,
			&last_name,
			&phone_number,
			&email,
			&password,
			&branch_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Drivers = append(resp.Drivers, &models.Driver{
			Id:          id.String,
			FirsName:    first_name.String,
			LastName:    last_name.String,
			CarId:       car_id.String,
			PhoneNumber: phone_number.String,
			Email:       email.String,
			Password:    password.String,
			BranchId:    branch_id.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *driverRepo) Update(ctx context.Context, req *models.UpdateDriver) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)
	query = `
		UPDATE
			Drivers
		SET
			first_name = :first_name,
			last_name = :last_name,
			car_id = :car_id,
			phone_number = :phone_number,
			email = :email,
			password = :password,
			branch_id = :branch_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"first_name":   req.FirsName,
		"last_name":    req.LastName,
		"car_id":       req.CarId,
		"phone_number": req.PhoneNumber,
		"email":        req.Email,
		"password":     req.Password,
		"branch_id":    req.BranchId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *driverRepo) Delete(ctx context.Context, req *models.DriverPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM Drivers WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
