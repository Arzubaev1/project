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

type carRepo struct {
	db *pgxpool.Pool
}

func NewCarRepo(db *pgxpool.Pool) *carRepo {
	return &carRepo{
		db: db,
	}
}

func (r *carRepo) Create(ctx context.Context, req *models.CreateCar) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO car(id, model, brand, state_number, updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Model,
		req.Brand,
		req.StateNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *carRepo) GetByID(ctx context.Context, req *models.CarPrimaryKey) (*models.Car, error) {
	var whereField = "id"
	var (
		query string

		id           sql.NullString
		model        sql.NullString
		brand        sql.NullString
		state_number sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query = `
		SELECT
			id,
			model,
			brand,
			state_number,
			created_at,
			updated_at
		FROM car
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&model,
		&brand,
		&state_number,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Car{
		Id:          id.String,
		Model:       model.String,
		Brand:       brand.String,
		StateNumber: state_number.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *carRepo) GetList(ctx context.Context, req *models.CarGetListRequest) (*models.CarGetListResponse, error) {

	var (
		resp   = &models.CarGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			model,
			brand,
			state_number,
			created_at,
			updated_at
		FROM car
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
			model        sql.NullString
			brand        sql.NullString
			state_number sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&model,
			&brand,
			&state_number,

			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Cars = append(resp.Cars, &models.Car{
			Id:          id.String,
			Model:       model.String,
			Brand:       brand.String,
			StateNumber: state_number.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *carRepo) Update(ctx context.Context, req *models.UpdateCar) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)
	query = `
		UPDATE
			car
		SET
			model = :model,
			brand = :brand,
			state_number = :state_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"model":        req.Model,
		"brand":        req.Brand,
		"state_number": req.StateNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *carRepo) Delete(ctx context.Context, req *models.CarPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM car WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
