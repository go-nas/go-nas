package service

import (
	"context"

	"github.com/shirou/gopsutil/v3/host"
)

func (srv *Service) HostInfo(ctx context.Context) (*host.InfoStat, error) {
	return host.InfoWithContext(ctx)
}
