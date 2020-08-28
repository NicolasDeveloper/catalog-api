package controllers

import (
	"net/http"

	"github.com.br/NicolasDeveloper/catalog-api/commands"
	"github.com.br/NicolasDeveloper/catalog-api/usecases"

	"github.com.br/NicolasDeveloper/catalog-api/api/common"
	"github.com/golobby/container"
)

//ProductController controller
type ProductController struct {
	common.Controller
}

//NewProductController contructor
func NewProductController() ProductController {
	return ProductController{}
}

//Index get products in database
func (c *ProductController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		nil,
		http.StatusOK,
	)
}

//Post insert product
func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	var usecase *usecases.RegisterProduct
	container.Make(&usecase)

	command := commands.ProductInput{}
	c.GetContent(command, r)

	output, err := usecase.Execute(command)

	if err != nil {
		c.SendJSON(
			w,
			nil,
			http.StatusBadRequest,
		)
	}

	c.SendJSON(
		w,
		output,
		http.StatusCreated,
	)
}
