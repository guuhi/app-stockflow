package usecase

import (
	"context"
	"stockflow/internal/dto"
	"stockflow/internal/entity"
	"stockflow/internal/shared"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, queryParams shared.ProductQueryParams) ([]entity.Product, error)
	GetProductsCount(ctx context.Context) (int, error)
}

type ProductUsecase struct {
	repo ProductRepository
}

func NewProductUsecase(repo ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) ListProducts(ctx context.Context, queryParams shared.ProductQueryParams) ([]dto.ProductResponse, int, error) {
	products, err := uc.repo.ListProducts(ctx, queryParams)
	if err != nil {
		return nil, 0, err
	}

	totalProducts, err := uc.repo.GetProductsCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = productToDTO(product)
	}

	return responses, totalProducts, nil
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
