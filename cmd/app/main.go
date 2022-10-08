package main

import (
	"log"
	"net/http"

	"github.com/go-nas/go-nas/internal/controller"
	"github.com/go-nas/go-nas/internal/route"
	"github.com/go-nas/go-nas/internal/service"
	"go.uber.org/dig"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	c := dig.New()

	c.Provide(func() *http.ServeMux {
		return http.NewServeMux()
	})

	c.Provide(service.NewService)

	c.Provide(controller.NewController)

	c.Invoke(route.Register)

	c.Invoke(func(mux *http.ServeMux) {
		addr := ":8080"
		log.Println("service listen on: ", addr)
		err := http.ListenAndServe(
			addr,
			h2c.NewHandler(mux, &http2.Server{}),
		)
		log.Fatalf("listen failed: %v", err)
	})
}
