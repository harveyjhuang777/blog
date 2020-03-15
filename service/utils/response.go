package response

import (
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/utils/errs"
)

type RequestBody map[string]interface{}

type HttpResponse struct {
	Code     int         `json:"-"`
	CodeName string      `json:"code"`
	Error    error       `json:"-"`
	Body     interface{} `json:"message"`
}

// RespondError makes the error response with payload as json format
func RespondError(c gin.Context, defaultCode int, err error) {
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, HttpResponse{CodeName: errs.CodeNotFound.Name(), Body: errs.ErrRecordNotFound.Error()})
		return
	}

	if IsDuplicateKeyError(err) {
		c.JSON(http.StatusConflict, HttpResponse{CodeName: errs.CodeConflict.Name(), Body: errs.ErrRecordExist.Error()})
		return
	}

	if IsValueTooLongError(err) {
		c.JSON(http.StatusBadRequest, HttpResponse{CodeName: errs.CodeFieldInvalid.Name(), Body: errs.ErrValueTooLong.Error()})
		return
	}

	if IsViolateFKConstraint(err) {
		c.JSON(http.StatusBadRequest, HttpResponse{CodeName: errs.CodeFieldInvalid.Name(), Body: errs.ErrInvalidRequest.Error()})
		return
	}

	if v := math.Floor(float64(defaultCode / 100)); v == 4.0 {
		c.JSON(defaultCode, HttpResponse{CodeName: errs.CodeBadRequest.Name(), Body: err.Error()})
		return
	}

	if v := math.Floor(float64(defaultCode / 100)); v == 5.0 {
		c.JSON(defaultCode, HttpResponse{CodeName: errs.CodeServerError.Name(), Body: err.Error()})
		return
	}
}

func IsDuplicateKeyError(err error) bool {
	if strings.Contains(err.Error(), "duplicate key") {
		return true
	}
	return false
}

func IsValueTooLongError(err error) bool {
	if strings.Contains(err.Error(), "too long") {
		return true
	}
	return false
}

func IsViolateFKConstraint(err error) bool {
	if strings.Contains(err.Error(), "violates foreign key constraint") {
		return true
	}
	return false
}
