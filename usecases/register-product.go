package usecases

import (
	"github.com.br/NicolasDeveloper/catalog-api/commands"
	"github.com.br/NicolasDeveloper/catalog-api/entities"
	"github.com.br/NicolasDeveloper/catalog-api/ports"
)

//RegisterProduct use case
type RegisterProduct struct {
	repository ports.IProductRepository
}

//NewRegisterProduct constructor
func NewRegisterProduct(repository ports.IProductRepository) *RegisterProduct {
	return &RegisterProduct{
		repository: repository,
	}
}

//Execute execute use case
func (u *RegisterProduct) Execute(request *commands.Product) error {
	entity, error := entities.NewProduct(
		request.Name,
		request.Price,
		request.Description,
	)

	u.repository.Save(entity)

	request.ID = entity.ID

	return error
}
