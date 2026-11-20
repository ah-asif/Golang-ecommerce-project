package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am a Software Engineer.")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Please give me get request", 400)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(productList)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server is running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          101,
		Title:       "Orange",
		Description: "Orange is yeollow",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/RrTrNH4X94RH0dfkpEwvmpoZ4aPB-kqanUb_riqk1nw/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wNzAv/MTMwLzIxNS9zbWFs/bC9jbG9zZS11cC1v/Zi1hbi1vcmFuZ2Ut/b24tdHJhbnNwYXJl/bnQtYmFja2dyb3Vu/ZC1oaWdoLXF1YWxp/dHktaW1hZ2Utd2hv/bGUtb3JhbmdlLXBl/cmZlY3QtZm9yLXBu/Zy5wbmc",
	}

	prd2 := Product{
		ID:          102,
		Title:       "Apple",
		Description: "Apple is very sweet",
		Price:       40,
		ImgUrl:      "https://imgs.search.brave.com/wrenyXJ9ChdHD6Ay3or0aiSirERdYCH4CMFV_Q-Va6U/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tcGhv/dG8vcmVkLWFwcGxl/LXdpdGgtbGVhZi1p/c29sYXRlZF8xMTAz/NTEzLTU0NC5qcGc_/c2VtdD1haXNfaHli/cmlkJnc9NzQwJnE9/ODA",
	}

	prd3 := Product{
		ID:          103,
		Title:       "Banana",
		Description: "Banana is Good",
		Price:       45,
		ImgUrl:      "https://imgs.search.brave.com/S2UptOhuywJw8B-NvyJnttRT3dU9jrEAWQZJmN2axI4/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvMTI5/MTI2MjExMi9waG90/by9iYW5hbmEuanBn/P2I9MSZzPTYxMng2/MTImdz0wJms9MjAm/Yz14WjBXQk11U2Iy/a0xHYWFJUVZsS25y/XzQwb0NiSHZpbTB2/YzkxUURScXhNPQ",
	}

	prd4 := Product{
		ID:          104,
		Title:       "Mango",
		Description: "Mango is well known to everyone.",
		Price:       80,
		ImgUrl:      "https://imgs.search.brave.com/xS54zCpqVDXbUIKE0qq_NOV7lVj4L7c62BVXhHijwzk/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tcGhv/dG8vc2luZ2xlLWZy/ZXNoLXJpcGUtbWFu/Z28tZnJ1aXQtd2l0/aC1ncmVlbi1sZWFm/LWlzb2xhdGVkLXdo/aXRlLWJhY2tncm91/bmQtc3R1ZGlvLWNh/cHR1cmVfMTI5MjMz/My0yNjAuanBnP3Nl/bXQ9YWlzX2h5YnJp/ZCZ3PTc0MCZxPTgw",
	}

	prd5 := Product{
		ID:          105,
		Title:       "Lichi",
		Description: "Lichi is very populer in bangladesh.",
		Price:       300,
		ImgUrl:      "https://imgs.search.brave.com/I4CUKx5Rx6isHFoyVwIVcBtJePtbUcvHNqHlImGxqZg/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90My5m/dGNkbi5uZXQvanBn/LzA2LzA5Lzg0LzMw/LzM2MF9GXzYwOTg0/MzAyOF9CWlcyT0p1/Zmc5S3FhOHdoODdq/N2REZElkZGREV1VG/Vy5qcGc",
	}

	prd6 := Product{
		ID:          106,
		Title:       "Guava",
		Description: "Guava is very helthy.",
		Price:       50,
		ImgUrl:      "https://imgs.search.brave.com/W8agWEDqnc5cgpmZThakOfOG19ZF9h4K0aLU7QV5vK0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wODgv/NzcwLzY4OS9zbWFs/bC9mcmVzaC13aG9s/ZS1hbmQtc2xpY2Vk/LXBpbmstZ3VhdmEt/ZnJ1aXRzLXdpdGgt/Z3JlZW4tbGVhZi1p/c29sYXRlZC1vbi13/aGl0ZS1iYWNrZ3Jv/dW5kLXBob3RvLmpw/Zw",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
