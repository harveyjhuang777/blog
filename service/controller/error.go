package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func firestoreResponse(c *gin.Context, err error) {
	switch {
	case grpc.Code(err) == codes.Canceled:
		c.JSON(http.StatusBadRequest, err)
	case grpc.Code(err) == codes.Unknown:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, err)
	case grpc.Code(err) == codes.DeadlineExceeded:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.NotFound:
		c.JSON(http.StatusNotFound, err)
	case grpc.Code(err) == codes.AlreadyExists:
		c.JSON(http.StatusConflict, err)
	case grpc.Code(err) == codes.PermissionDenied:
		c.JSON(http.StatusForbidden, err)
	case grpc.Code(err) == codes.ResourceExhausted:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.FailedPrecondition:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Aborted:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.OutOfRange:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unimplemented:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Internal:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unavailable:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.DataLoss:
		c.JSON(http.StatusInternalServerError, err)
	case grpc.Code(err) == codes.Unauthenticated:
		c.JSON(http.StatusForbidden, err)
	default:
		c.JSON(http.StatusInternalServerError, err)
	}
}
