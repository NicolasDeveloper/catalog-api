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
func (u *RegisterProduct) Execute(input commands.ProductInput) (commands.ProductOutput, error) {
	entity, error := entities.NewProduct(
		input.Name,
		input.Price,
		input.Description,
	)

	u.repository.Save(entity)

	return commands.ProductOutput{
		ID: entity.ID,
	}, error
}
