package nextID

import (
	"ecommerce/database"
)

func GetNextID() int {
	if len(database.ProductList) == 0 {
		return 1
	}
	return database.ProductList[len(database.ProductList)-1].ID + 1
}