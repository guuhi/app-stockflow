package entity

import "time"

type Product struct {
	ID             int        `db:"id"`
	SKU            string     `db:"sku"`
	Name           string     `db:"name"`
	Price          int        `db:"price"`
	Stock          int        `db:"stock"`
	IsActive       bool       `db:"is_active"`
	CreatedAt      time.Time  `db:"created_at"`
	LastMovementAt *time.Time `db:"last_movement_at"`
	Total          int
}
