package controller

import (
	"github.com/go-nas/go-nas/internal/service"
	"go.buf.build/bufbuild/connect-go/go-nas/go-nas/gonas/v1alpha1/gonasv1alpha1connect"
)

var _ gonasv1alpha1connect.DeviceServiceClient = new(Controller)

type Controller struct {
	gonasv1alpha1connect.UnimplementedDeviceServiceHandler
	srv *service.Service
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		srv: srv,
	}
}
