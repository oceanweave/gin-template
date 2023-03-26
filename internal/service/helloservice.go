package service

import "gin-template/pkg/statuscode"

type HelloService struct {
}

func (s HelloService) GetMessage() (string, statuscode.Code) {
	return "hello,world", statuscode.Success
}
