package handler

import (
	"errors"
	"log"
	"net/http"
	"stockflow/internal/apperror"

	"github.com/gin-gonic/gin"
)

var statusByCode = map[apperror.Code]int{
	apperror.CodeValidation: http.StatusBadRequest,
	apperror.CodeNotFound:   http.StatusNotFound,
	apperror.CodeInternal:   http.StatusInternalServerError,
}

const genericInternalMessage = "internal server error"

func RespondData(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data})
}

func RespondDataWithPagination(c *gin.Context, status int, data any, page int, limit int) {
	c.JSON(status, gin.H{"data": data, "page": page, "limit": limit})
}

func RespondError(c *gin.Context, status int, code apperror.Code, message string, fields map[string]string) {
	body := gin.H{"code": code, "message": message}
	if len(fields) > 0 {
		body["fields"] = fields
	}
	c.JSON(status, gin.H{"error": body})
}

func RespondUsecaseError(c *gin.Context, err error) {
	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		appErr = apperror.Internal(err)
	}

	status, ok := statusByCode[appErr.Code]
	if !ok {
		status = http.StatusInternalServerError
	}

	code, message := appErr.Code, appErr.Message
	if status == http.StatusInternalServerError {
		log.Printf("internal error: %v", err)
		message = genericInternalMessage
		code = apperror.CodeInternal
	}

	if status == http.StatusBadGateway {
		log.Printf("upstream error: %v", err)
	}

	RespondError(c, status, code, message, appErr.Fields)
}
