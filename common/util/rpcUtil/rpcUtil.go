package rpcUtil

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.uber.org/zap"
)

type Client struct {
	resty  *resty.Client
	logger *zap.Logger
	nacos  *naming_client.INamingClient
}

// avoid import cycle
func New(logger *zap.Logger, nacos *naming_client.INamingClient) *Client {
	return &Client{
		resty:  resty.New(),
		logger: logger,
		nacos:  nacos,
	}
}

func (i *Client) Get(name string, path string, querys map[string]string, result any) error {
	req := i.resty.R()
	if querys != nil {
		req.SetQueryParams(querys)
	}
	req.SetResult(result)
	_, err := req.Get(fmt.Sprintf("http://%s%s", i.getInstanceHost(name), path))
	return err
}

func (i *Client) GetWithPathVariable(name string, path string, pathVar string, result any) error {
	req := i.resty.R()
	req.SetResult(result)
	_, err := req.Get(fmt.Sprintf("http://%s%s/%s", i.getInstanceHost(name), path, pathVar))
	return err
}

func (i *Client) Post(name string, path string, body any, result any) error {
	req := i.resty.R()
	if body != nil {
		req.SetBody(body)
	}
	req.SetResult(result)
	_, err := req.Post(fmt.Sprintf("http://%s%s", i.getInstanceHost(name), path))
	return err
}

func (i *Client) getInstanceHost(name string) string {
	instance, err := (*i.nacos).SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
	})
	if err != nil {
		i.logger.Panic(fmt.Sprintf("failed to get %s instance: %s", name, err.Error()))
	}
	return fmt.Sprintf("%s:%d", instance.Ip, instance.Port)
}
