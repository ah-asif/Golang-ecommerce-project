package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.LoggerHudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Backend server is running on: 3000")

	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error Starting the Server: ", err)
	}
}
