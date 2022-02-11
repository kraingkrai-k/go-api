package response

import (
	"errors"
	"github.com/kraingkrai-k/go-api/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Error        string      `json:"error,omitempty"`
	Total        int32       `json:"total,omitempty"`
	TotalRecords int32       `json:"totalRecords,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (r Response) Paginate(paginate model.Paginate) *Response {
	return &Response{
		Status:       r.Status,
		Data:         r.Data,
		Message:      r.Message,
		Error:        r.Error,
		TotalRecords: paginate.TotalRecords,
		Total:        paginate.Total,
	}
}

func New(statusCode int, output interface{}, paginate model.Paginate) *Response {
	var response = Response{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	}

	if err, ok := output.(error); ok {
		response.Error = err.Error()
		return &response
	}

	response.Data = output
	return response.Paginate(paginate)
}

func Output(ctx *gin.Context, statusCode int, output interface{}, paginate model.Paginate) {
	response := New(statusCode, output, paginate)
	ctx.JSON(statusCode, response)
	ctx.Abort()
}

func Success(ctx *gin.Context, output interface{}, paginates ...model.Paginate) {
	if len(paginates) > 0 {
		Output(ctx, http.StatusOK, output, paginates[0])
		return
	}
	Output(ctx, http.StatusOK, output, model.Paginate{})
}

func Error(ctx *gin.Context, output interface{}) {
	statusCode := http.StatusBadRequest
	if data, ok := output.(Response); ok {
		if data.Status != 0 {
			statusCode = data.Status
		}
		output = errors.New(data.Error)
	}

	Output(ctx, statusCode, output, model.Paginate{})
}
