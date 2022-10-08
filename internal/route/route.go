package route

import (
	"net/http"

	"github.com/go-nas/go-nas/internal/controller"
	"go.buf.build/bufbuild/connect-go/go-nas/go-nas/gonas/v1alpha1/gonasv1alpha1connect"
)

func Register(mux *http.ServeMux, ctl *controller.Controller) {
	mux.Handle(gonasv1alpha1connect.NewDeviceServiceHandler(ctl))
}
