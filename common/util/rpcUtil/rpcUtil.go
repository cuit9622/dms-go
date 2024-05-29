package rpcUtil

import (
	"fmt"

	"github.com/cuit9622/dms/common/util"

	"github.com/go-resty/resty/v2"
)

type client struct {
	resty *resty.Client
}

func New() *client {
	return &client{
		resty: resty.New(),
	}
}

func (i client) Get(name string, path string, querys map[string]string, result any) error {
	req := i.resty.R()
	if querys != nil {
		req.SetQueryParams(querys)
	}
	req.SetResult(result)
	_, err := req.Get(fmt.Sprintf("http://%s%s", i.getInstanceHost(name), path))
	return err
}

func (i client) GetWithPathVariable(name string, path string, pathVar string, result any) error {
	req := i.resty.R()
	req.SetResult(result)
	_, err := req.Get(fmt.Sprintf("http://%s%s/%s", i.getInstanceHost(name), path, pathVar))
	return err
}

func (i client) Post(name string, path string, body any, result any) error {
	req := i.resty.R()
	if body != nil {
		req.SetBody(body)
	}
	req.SetResult(result)
	_, err := req.Post(fmt.Sprintf("http://%s%s", i.getInstanceHost(name), path))
	return err
}

func (i client) getInstanceHost(name string) string {
	instance := util.GetInstance(name)
	return fmt.Sprintf("%s:%d", instance.Ip, instance.Port)
}
