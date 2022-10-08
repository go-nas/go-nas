package controller

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "go.buf.build/bufbuild/connect-go/go-nas/go-nas/gonas/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ctl *Controller) GetHostInfo(ctx context.Context, _ *connect_go.Request[v1alpha1.GetHostInfoRequest]) (*connect_go.Response[v1alpha1.GetHostInfoResponse], error) {
	info, err := ctl.srv.HostInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return connect_go.NewResponse(&v1alpha1.GetHostInfoResponse{
		HostName:        info.Hostname,
		UpTime:          int64(info.Uptime),
		Os:              info.OS,
		Platform:        info.Platform,
		PlatformFamily:  info.PlatformFamily,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}), nil
}
