package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.PrefLight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Backend server is running on: 3000")

	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		fmt.Println("Error Starting the Server: ", err)
	}
}
