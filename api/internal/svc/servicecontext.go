package svc

import (
	"shorturl/api/internal/config"
	transformer "shorturl/rpc/transform/transformclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transform
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransform(zrpc.MustNewClient(c.Transform)),
	}
}
