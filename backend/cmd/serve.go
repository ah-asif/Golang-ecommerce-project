package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Backend server is running on: 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error Starting the Server: ", err)
	}
}
