package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jwjhuang/blog/service/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func responseWithJSON(c *gin.Context, data interface{}) {
	res := model.Response{
		Status: model.ResponseStatusOK,
		Data:   data,
	}
	c.JSON(http.StatusOK, res)
}

func abortWithError(c *gin.Context, code int, err error) {
	res := model.Response{
		Status:  model.ResponseStatusFail,
		Message: err.Error(),
	}
	c.AbortWithStatusJSON(code, res)
}

func firestoreResponse(c *gin.Context, err error) {
	switch {
	case grpc.Code(err) == codes.Canceled:
		abortWithError(c, http.StatusBadRequest, err)
	case grpc.Code(err) == codes.Unknown:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.InvalidArgument:
		abortWithError(c, http.StatusBadRequest, err)
	case grpc.Code(err) == codes.DeadlineExceeded:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.NotFound:
		abortWithError(c, http.StatusNotFound, err)
	case grpc.Code(err) == codes.AlreadyExists:
		abortWithError(c, http.StatusConflict, err)
	case grpc.Code(err) == codes.PermissionDenied:
		abortWithError(c, http.StatusForbidden, err)
	case grpc.Code(err) == codes.ResourceExhausted:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.FailedPrecondition:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Aborted:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.OutOfRange:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unimplemented:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Internal:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unavailable:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.DataLoss:
		abortWithError(c, http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unauthenticated:
		abortWithError(c, http.StatusForbidden, err)
	default:
		abortWithError(c, http.StatusInternalServerError, err)
	}
}
