package repo

import (
	"context"
	"stockflow/internal/apperror"
	"stockflow/internal/entity"
	"stockflow/internal/shared"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

var filterWhere = map[string]string{
	"active":    " WHERE p.is_active = TRUE",
	"inactive":  " WHERE p.is_active <> TRUE",
	"low_stock": " WHERE p.stock <= 10",
}

func (r *ProductRepository) ListProducts(ctx context.Context, queryParams shared.ProductQueryParams) ([]entity.Product, error) {
	args := []any{}

	query := `
		SELECT p.id, p.sku, p.name, p.price, p.stock, p.is_active, p.created_at, COALESCE(s.created_at, p.created_at) as last_movement_at
		FROM products p
		LEFT JOIN stock_movements s ON s.product_id = p.id
	`

	filter, filterOk := filterWhere[queryParams.Filter];
	if filterOk {
		query += filter
	}

	if queryParams.Search != "" {
		search := "%" + queryParams.Search + "%"

		if filterOk {
			query += " AND (p.sku ILIKE ? OR p.name ILIKE ?)"
		} else {
			query += " WHERE (p.sku ILIKE ? OR p.name ILIKE ?)"
		}
		
		args = append(args, search, search)
	}

	query += " ORDER BY p.sku"

	query += " LIMIT ? OFFSET ?"
	args = append(args, queryParams.Limit, queryParams.Offset)

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, apperror.ErrDatabaseQueryFailed.Wrap(err)
	}

	query = r.db.Rebind(query)

	result := []entity.Product{}
	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return nil, apperror.ErrDatabaseQueryFailed.Wrap(err)
	}

	return result, nil
}
