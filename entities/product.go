package entities

import (
	"github.com/beevik/guid"
)

//Product entity
type Product struct {
	ID          string  `bson:"_id" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Price       float64 `bson:"price" json:"price"`
	Description string  `bson:"description" json:"description"`
}

//NewProduct constructor
func NewProduct(
	name string,
	price float64,
	descripion string,
) (Product, error) {
	newguid := guid.New()

	return Product{
		ID:          newguid.String(),
		Name:        name,
		Price:       price,
		Description: descripion,
	}, nil
}
