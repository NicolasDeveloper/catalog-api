package usecases

import (
	"testing"

	"github.com.br/NicolasDeveloper/catalog-api/commands"
	"github.com.br/NicolasDeveloper/catalog-api/entities"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(product entities.Product) error {
	product.ID = ""
	args := r.Called(product)
	return args.Error(0)
}

func TestRegisterUseCase(t *testing.T) {
	t.Run("Should register a new product in catalog", func(t *testing.T) {
		repository := new(repositoryMock)

		product := entities.Product{
			Name:        "Teste",
			Price:       19.99,
			Description: "Teste de descrição",
		}

		repository.On("Save", product).Return(nil)

		productDTO := commands.Product{
			Name:        "Teste",
			Price:       19.99,
			Description: "Teste de descrição",
		}

		useCase := NewRegisterProduct(repository)
		useCase.Execute(&productDTO)

		repository.AssertExpectations(t)
	})
}
