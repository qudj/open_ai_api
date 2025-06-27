package responses

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	// Success http status 200, grpc status code OK(0)
	Success = Response{http.StatusOK, ResponseData{0, "", "success", nil}}
)

var (
	// InvalidArgument http status 400, grpc status code InvalidArgument(3)
	InvalidArgument = Response{http.StatusBadRequest, ResponseData{40000300, ReasonBy("INVALID_ARGUMENT"), "invalid argument", nil}}
	// Unauthenticated http status 401, grpc status code OutOfRange(16)
	Unauthenticated = Response{http.StatusUnauthorized, ResponseData{40101600, ReasonBy("UNAUTHORIZED"), "unauthorized", nil}}
	// NotFound http status 404, grpc status code OutOfRange(5)
	NotFound = Response{http.StatusNotFound, ResponseData{40400500, ReasonBy("RESOURCE_NOT_FOUND"), "resource not found", nil}}

	InternalError = Response{http.StatusInternalServerError, ResponseData{50001300, ReasonBy("INTERNAL_ERROR"), "internal error", nil}}
	TimeoutError  = Response{http.StatusGatewayTimeout, ResponseData{50401800, ReasonBy("TIMEOUT"), "timeout", nil}}
)

func ReasonBy(s string) string {
	return fmt.Sprintf("P_OPEN_AI_%s", strings.ToUpper(s))
}
