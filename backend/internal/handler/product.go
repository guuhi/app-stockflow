package handler

import (
	"context"
	"net/http"
	"stockflow/internal/apperror"
	"stockflow/internal/dto"
	"stockflow/internal/shared"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductUsecase interface {
	ListProducts(ctx context.Context, queryParams shared.ProductQueryParams) ([]dto.ProductResponse, error)
}

type ProductHandler struct {
	uc ProductUsecase
}

func NewProductHandler(uc ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	search := strings.TrimSpace(c.DefaultQuery("search", ""))

	filter := strings.TrimSpace(c.DefaultQuery("filter", "all"))
	if _, ok := shared.ValidFilters[filter]; !ok {
		_ = c.Error(apperror.New(apperror.CodeValidation, "invalid query parameter"))
		return
	}

	page, err := strconv.Atoi(strings.TrimSpace(c.DefaultQuery("page", "1")))
	if err != nil {
		_ = c.Error(apperror.New(apperror.CodeValidation, "invalid query parameter"))
		return
	}
	if page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(strings.TrimSpace(c.DefaultQuery("limit", "5")))
	if err != nil {
		_ = c.Error(apperror.New(apperror.CodeValidation, "invalid query parameter"))
		return
	}
	if limit < 1 || limit > 100 {
		limit = 5
	}

	offset := (page - 1) * limit

	queryParams := shared.ProductQueryParams{
		Search: search,
		Filter: filter,
		Page: page,
		Limit: limit,
		Offset: offset,
	}

	products, err := h.uc.ListProducts(c.Request.Context(), queryParams)
	if err != nil {
		_ = c.Error(err)
		return
	}

	RespondDataWithPagination(c, http.StatusOK, products, page, limit)
}
