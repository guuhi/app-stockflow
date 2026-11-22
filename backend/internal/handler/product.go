package handler

import (
	"context"
	"log"
	"net/http"
	"stockflow/internal/dto"

	"github.com/gin-gonic/gin"
)

type ProductUsecase interface {
	ListProducts(ctx context.Context) ([]dto.ProductResponse, error)
}

type ProductHandler struct {
	uc ProductUsecase
}

func NewProductHandler(uc ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.uc.ListProducts(c.Request.Context())
	if err != nil {
		log.Println(err)
		_ = c.Error(err)
		return
	}

	RespondData(c, http.StatusOK, products)
}
