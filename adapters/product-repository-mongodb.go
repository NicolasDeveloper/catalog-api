package adapters

import (
	"context"

	"github.com.br/NicolasDeveloper/catalog-api/ports"

	"github.com.br/NicolasDeveloper/catalog-api/entities"
)

type repository struct {
	ctx *DbContext
	ports.IProductRepository
}

//NewProductsRepository constructor
func NewProductsRepository(ctx *DbContext) (ports.IProductRepository, error) {
	return &repository{
		ctx: ctx,
	}, nil
}

//Save search products in database
func (r *repository) Save(product entities.Product) error {
	collection, error := r.ctx.GetCollection(product)
	_, error = collection.InsertOne(context.TODO(), product)
	return error
}
