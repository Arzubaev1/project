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

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {

	var (
		id                        = uuid.New().String()
		query                     string
		first_client_total_price  int64
		second_client_total_price int64
	)
	if req.Discount == "fixed" {
		first_client_total_price = req.FirstClientMillage*req.PriceForMillage - req.DiscountPrice
	} else if req.Discount == "percent" {
		first_client_total_price = req.FirstClientMillage*req.PriceForMillage - (req.DiscountPrice * req.DiscountPrice / 100)
	}
	query = `
		INSERT INTO order(
			id, 
			driver_id, 
			date,
			discount, 
			discount_price, 
			branch_id,
			status,
			first_client_id,
			second_client_id,
			first_client_location,
			second_client_location,
			first_client_destination,
			second_client_destination,
			first_client_millage,
			second_client_millage,
			price_for_millage,
			first_client_total_price,
			second_client_total_price,
			payment_type,
			updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, 'searching', $8,$9,$10, $11, $12, $13, $14, $15, $17, $18, $19, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.DriverId,
		req.Date,
		req.Discount,
		req.DiscountPrice,
		req.BranchId,
		req.FirstClientId,
		req.SecondClientId,
		req.FirstClientLocation,
		req.SecondClientLocation,
		req.FirstClientDestination,
		req.SecondClientDestination,
		req.FirstClientMillage,
		req.SecondClientMillage,
		req.PriceForMillage,
		first_client_total_price,
		second_client_total_price,
		req.PaymentType,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	var whereField = "id"
	var (
		query string

		id                        sql.NullString
		driver_id                 sql.NullString
		date                      sql.NullString
		discount                  sql.NullString
		discount_price            sql.NullInt64
		branch_id                 sql.NullString
		status                    sql.NullString
		first_client_id           sql.NullString
		second_client_id          sql.NullString
		first_client_location     sql.NullString
		second_client_location    sql.NullString
		first_client_destination  sql.NullString
		second_client_destination sql.NullString
		first_client_millage      sql.NullInt64
		second_client_millage     sql.NullInt64
		price_for_millage         sql.NullInt64
		first_client_total_price  sql.NullInt64
		second_client_total_price sql.NullInt64
		payment_type              sql.NullString
		createdAt                 sql.NullString
		updatedAt                 sql.NullString
	)

	query = `
		SELECT
			id, 
			driver_id, 
			date,
			discount, 
			discount_price, 
			branch_id,
			status,
			first_client_id,
			second_client_id,
			first_client_location,
			second_client_location,
			first_client_destination,
			second_client_destination,
			first_client_millage,
			second_client_millage,
			price_for_millage,
			first_client_total_price,
			second_client_total_price,
			payment_type,
			created_at,
			updated_at
		FROM order
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&driver_id,
		&date,
		&discount,
		&discount_price,
		&branch_id,
		&status,
		&first_client_location,
		&second_client_location,
		&first_client_destination,
		&second_client_destination,
		&first_client_millage,
		&second_client_millage,
		&price_for_millage,
		&first_client_total_price,
		&second_client_total_price,
		&payment_type,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:                      id.String,
		DriverId:                driver_id.String,
		Date:                    date.String,
		Discount:                discount.String,
		DiscountPrice:           discount_price.Int64,
		BranchId:                branch_id.String,
		Status:                  status.String,
		FirstClientId:           first_client_id.String,
		SecondClientId:          second_client_id.String,
		FirstClientLocation:     first_client_location.String,
		SecondClientLocation:    second_client_location.String,
		FirstClientDestination:  first_client_destination.String,
		SecondClientDestination: second_client_destination.String,
		FirstClientMillage:      first_client_millage.Int64,
		SecondClientMillage:     second_client_millage.Int64,
		PriceForMillage:         price_for_millage.Int64,
		FirstClientTotalPrice:   first_client_total_price.Int64,
		SecondClientTotalPrice:  second_client_total_price.Int64,
		PaymentType:             payment_type.String,
		CreatedAt:               createdAt.String,
		UpdatedAt:               updatedAt.String,
	}, nil
}

func (r *orderRepo) GetList(ctx context.Context, req *models.OrderGetListRequest) (*models.OrderGetListResponse, error) {

	var (
		resp   = &models.OrderGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			driver_id, 
			date,
			discount, 
			discount_price, 
			branch_id,
			status,
			first_client_id,
			second_client_id,
			first_client_location,
			second_client_location,
			first_client_destination,
			second_client_destination,
			first_client_millage,
			second_client_millage,
			price_for_millage,
			first_client_total_price,
			second_client_total_price,
			payment_type,
			created_at,
			updated_at
		FROM order
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
			id                        sql.NullString
			driver_id                 sql.NullString
			date                      sql.NullString
			discount                  sql.NullString
			discount_price            sql.NullInt64
			branch_id                 sql.NullString
			status                    sql.NullString
			first_client_id           sql.NullString
			second_client_id          sql.NullString
			first_client_location     sql.NullString
			second_client_location    sql.NullString
			first_client_destination  sql.NullString
			second_client_destination sql.NullString
			first_client_millage      sql.NullInt64
			second_client_millage     sql.NullInt64
			price_for_millage         sql.NullInt64
			first_client_total_price  sql.NullInt64
			second_client_total_price sql.NullInt64
			payment_type              sql.NullString
			createdAt                 sql.NullString
			updatedAt                 sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&driver_id,
			&date,
			&discount,
			&discount_price,
			&branch_id,
			&status,
			&first_client_location,
			&second_client_location,
			&first_client_destination,
			&second_client_destination,
			&first_client_millage,
			&second_client_millage,
			&price_for_millage,
			&first_client_total_price,
			&second_client_total_price,
			&payment_type,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &models.Order{
			Id:                      id.String,
			DriverId:                driver_id.String,
			Date:                    date.String,
			Discount:                discount.String,
			DiscountPrice:           discount_price.Int64,
			BranchId:                branch_id.String,
			Status:                  status.String,
			FirstClientId:           first_client_id.String,
			SecondClientId:          second_client_id.String,
			FirstClientLocation:     first_client_location.String,
			SecondClientLocation:    second_client_location.String,
			FirstClientDestination:  first_client_destination.String,
			SecondClientDestination: second_client_destination.String,
			FirstClientMillage:      first_client_millage.Int64,
			SecondClientMillage:     second_client_millage.Int64,
			PriceForMillage:         price_for_millage.Int64,
			FirstClientTotalPrice:   first_client_total_price.Int64,
			SecondClientTotalPrice:  second_client_total_price.Int64,
			PaymentType:             payment_type.String,
			CreatedAt:               createdAt.String,
			UpdatedAt:               updatedAt.String,
		})
	}

	return resp, nil
}

func (r *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)
	query = `
		UPDATE
			order
		SET
			driver_id = :driver_id,                
			date = :date,                    
			discount = :discount,                 
			discount_price = :discount_price,           
			branch_id =  :branch_id,               
			status = :status,                  
			first_client_id = :first_client_id,         
			second_client_id = :second_client_id,        
			first_client_location = :first_client_location,
			second_client_location = :second_client_location,
			first_client_destination = :first_client_destination,
			second_client_destination = :second_client_destination,
			first_client_millage = :first_client_millage,
			second_client_millage = :second_client_millage,
			price_for_millage =  :price_for_millage,
			first_client_total_price = :first_client_total_price,
			second_client_total_price =  :second_client_total_price,
			payment_type = :payment_type,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                        req.Id,
		"driver_id":                 req.DriverId,
		"date":                      req.Date,
		"discount":                  req.Discount,
		"discount_price":            req.DiscountPrice,
		"branch_id":                 req.BranchId,
		"first_client_id":           req.FirstClientId,
		"second_client_id":          req.SecondClientId,
		"first_client_location":     req.FirstClientLocation,
		"second_client_location":    req.SecondClientLocation,
		"first_client_destination":  req.FirstClientDestination,
		"second_client_destination": req.SecondClientDestination,
		"first_client_millage":      req.FirstClientMillage,
		"second_client_millage":     req.SecondClientMillage,
		"price_for_millage":         req.PriceForMillage,
		"first_client_total_price":  req.FirstClientTotalPrice,
		"second_client_total_price": req.SecondClientTotalPrice,
		"payment_type":              req.PaymentType,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM order WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
