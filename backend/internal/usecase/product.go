package usecase

import (
	"context"
	"stockflow/internal/dto"
	"stockflow/internal/entity"
)

type ProductRepository interface {
	ListProducts(ctx context.Context) ([]entity.Product, error)
}

type ProductUsecase struct {
	repo ProductRepository
}

func NewProductUsecase(repo ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) ListProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := uc.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = productToDTO(product)
	}

	return responses, nil
}

func productToDTO(p entity.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:             p.ID,
		SKU:            p.SKU,
		Name:           p.Name,
		Price:          p.Price,
		Stock:          p.Stock,
		IsActive:       p.IsActive,
		CreatedAt:      p.CreatedAt,
		LastMovementAt: p.LastMovementAt,
	}
}
