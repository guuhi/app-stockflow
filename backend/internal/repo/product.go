package repo

import (
	"context"
	"stockflow/internal/apperror"
	"stockflow/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) ListProducts(ctx context.Context) ([]entity.Product, error) {
	result := []entity.Product{}

	err := r.db.SelectContext(ctx, &result, `
		SELECT p.id, p.sku, p.name, p.price, p.stock, p.is_active, p.created_at, COALESCE(s.created_at, p.created_at) as last_movement_at
		FROM products p
		LEFT JOIN stock_movements s ON s.product_id = p.id
		ORDER BY p.sku
	`)
	if err != nil {
		return nil, apperror.ErrDatabaseQueryFailed.Wrap(err)
	}
	return result, nil
}
