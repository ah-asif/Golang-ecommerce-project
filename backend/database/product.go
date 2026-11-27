package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func init() {

	prd1 := Product{
		ID: 101,
		Title: "Orange",
		Description: "Orange is Red. I love Orange",
		Price: 100,
		ImgURL: "https://media.istockphoto.com/id/185284489/photo/orange.jpg?s=612x612&w=0&k=20&c=m4EXknC74i2aYWCbjxbzZ6EtRaJkdSJNtekh4m1PspE=",
	}

	prd2 := Product{
		ID: 102,
		Title: "Apple",
		Description: "Apple is Green. I hate Orange",
		Price: 120,
		ImgURL: "https://cdn.pixabay.com/photo/2016/09/29/08/33/apple-1702316_1280.jpg",
	}

	prd3 := Product{
		ID: 103,
		Title: "Banana",
		Description: "Banana is boring. I don't like Banana at all!",
		Price: 80,
		ImgURL: "https://cdn.pixabay.com/photo/2016/09/03/20/48/bananas-1642706_1280.jpg",
	}

	prd4 := Product{
		ID: 104,
		Title: "Mango",
		Description: "Mango is interesting. I  like Mango very much!",
		Price: 50,
		ImgURL: "https://img.freepik.com/free-psd/delicious-mango-studio_23-2151843107.jpg?semt=ais_hybrid&w=740&q=80",
	}

	prd5 := Product{
		ID: 105,
		Title: "Grapes",
		Description: "Grapes is interesting. It is very helpful for health.",
		Price: 140,
		ImgURL: "https://media.istockphoto.com/id/184099853/photo/fresh-grapes.jpg?s=612x612&w=0&k=20&c=IxIcZ4M8c4-ksK4bvmtKUp03jaZwRMBcxplyZB-iZ0A=",
	}

	prd6 := Product{
		ID: 106,
		Title: "Nut",
		Description: "Nut is nutritious. It is very helpful for health.",
		Price: 70,
		ImgURL: "https://media.istockphoto.com/id/496689738/photo/assorted-nuts.jpg?s=612x612&w=0&k=20&c=lJhqPaHqwvXiDFNni5nB9EKgvYlqMEljI-0JzaB-ZNA=",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)
}
