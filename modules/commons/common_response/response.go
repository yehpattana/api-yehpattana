package commonresponse

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natersland/b2b-e-commerce-api/utils/logger"
)

type ResponseInterface interface {
	Success(code int, data any) ResponseInterface
	Error(code int, treceId string, errorMessage string) ResponseInterface
	Res() error
}

type CommonResponse struct {
	Status       bool   `json:"status"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	StatusCode    int
	Data          any
	ErrorResponse *ErrorResponse
	Context       *fiber.Ctx
	IsError       bool
}

type ErrorResponse struct {
	TraceId      string `json:"trace_id"`
	ErrorMessage string `json:"message"`
}

type PaginateResponse struct {
	Data      any `json:"data"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}

func NewResponse(c *fiber.Ctx) ResponseInterface {
	return &Response{
		Context: c,
	}
}

func (r *Response) Success(code int, data any) ResponseInterface {
	r.StatusCode = code
	r.Data = data
	logger.InitLogger(r.Context, &r.Data, code).Print()
	return r
}

func (r *Response) Error(code int, treceId string, errorMessage string) ResponseInterface {
	r.StatusCode = code
	r.ErrorResponse = &ErrorResponse{
		TraceId:      treceId,
		ErrorMessage: errorMessage,
	}
	r.IsError = true
	logger.InitLogger(r.Context, &r.ErrorResponse, code).Print()
	return r
}

func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			return &r.ErrorResponse
		}
		return &r.Data
	}())
}

type SingularResponse[T any] struct {
	Data *T `json:"data"`
}
