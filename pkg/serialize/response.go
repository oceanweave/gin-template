package serialize

import (
	"gin-template/pkg/statuscode"
	"net/http"
	"sync"
)

type resResult struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

type Response struct {
	HttpStatus int
	R          resResult
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Response{}
	},
}

func NewResponse(status int, code statuscode.Code, data interface{}) *Response {
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.R.Status = code.Code()
	response.R.Message = code.Message()
	response.R.Data = data

	return response
}

func PutResponse(res *Response) {
	if res != nil {
		res.R.Data = nil
		pool.Put(res)
	}

}

func ResponseOK(code statuscode.Code, data interface{}) *Response {
	return NewResponse(http.StatusOK, code, data)
}

func ResponseSuccess(data interface{}) *Response {
	return NewResponse(http.StatusOK, statuscode.Success, data)
}

func ResponseError(code statuscode.Code) *Response {
	return NewResponse(http.StatusOK, code, nil)
}

func ResponseUnknownError() *Response {
	return NewResponse(http.StatusOK, statuscode.UnknownError, nil)
}
