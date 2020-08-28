package ports

import "github.com.br/NicolasDeveloper/catalog-api/entities"

//IProductRepository database port
type IProductRepository interface {
	Save(product entities.Product) error
}
