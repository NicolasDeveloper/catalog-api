package ioc

import (
	"github.com.br/NicolasDeveloper/catalog-api/adapters"
	"github.com.br/NicolasDeveloper/catalog-api/ports"
	"github.com.br/NicolasDeveloper/catalog-api/usecases"
	"github.com/golobby/container"
)

//Container typing
type Container struct {
	IContainer
}

//NewContainer retrive a dig container resolution
func NewContainer() IContainer {
	return &Container{}
}

//RegisterDependencies register dependency injection
func (i *Container) RegisterDependencies() error {
	container.Singleton(func() *adapters.DbContext {
		return adapters.NewDbContext()
	})

	container.Transient(func(dbcontext *adapters.DbContext) ports.IProductRepository {
		adapter, _ := adapters.NewProductsRepository(dbcontext)
		return adapter
	})

	container.Transient(func(port ports.IProductRepository) *usecases.RegisterProduct {
		return usecases.NewRegisterProduct(port)
	})

	return nil
}
