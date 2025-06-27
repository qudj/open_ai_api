package responses

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qudj/open_ai_api/utils"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Status int
	ResponseData
}

type APIError interface {
	Response() *Response
	Error() string
}

func (s *Response) Error() string {
	return s.Message
}

func (s *Response) Response() *Response {
	return s
}

func (s Response) WithData(data interface{}) *Response {
	s.Data = data
	return &s
}

func (s Response) WithMessage(message string) *Response {
	s.Message = message
	return &s
}

func (s Response) WithCodeMessage(code int, message string) *Response {
	s.Code = s.Code + code
	s.Message = message
	return &s
}

func (s Response) WithCodeMessageData(code int, message string, data interface{}) *Response {
	s.Code = s.Code + code
	s.Message = message
	s.Data = data
	return &s
}

func (s Response) WithCodeError(code int, error error) *Response {
	s.Code = s.Code + code
	s.Message = error.Error()
	return &s
}

func (s Response) WithError(error error) *Response {
	s.Message = error.Error()
	return &s
}

func (s Response) Resp(ctx *gin.Context) {
	ctx.JSON(s.Status, &s.ResponseData)
}

func (s Response) RespData(ctx *gin.Context, data interface{}) {
	s.Data = data
	ctx.JSON(s.Status, &s.ResponseData)
}

func (s Response) RespCodeMessage(ctx *gin.Context, code int, message string) {
	s.Code = s.Code + code
	s.Message = message
	ctx.JSON(s.Status, &s.ResponseData)
}

func (s Response) RespCodeMessageData(ctx *gin.Context, message string, data interface{}) {
	s.Message = message
	s.Data = data
	ctx.JSON(s.Status, &s.ResponseData)
}

func (s Response) RespErrorMessage(ctx *gin.Context, err error) {
	s.Message = err.Error()
	ctx.JSON(s.Status, &s.ResponseData)
}

func RespFromError(ctx *gin.Context, error error) {
	if apiError, ok := error.(APIError); ok {
		utils.SugarLogger().Errorf("api error: %v", error)
		apiError.Response().Resp(ctx)
	} else if error == context.DeadlineExceeded {
		TimeoutError.Resp(ctx)
	} else {
		utils.SugarLogger().Errorf("invalid error type: %v", error)
		InternalError.RespCodeMessage(ctx, -99, error.Error())
	}
}
