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

}

//Post insert product
func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	var usecase *usecases.RegisterProduct
	container.Make(&usecase)

	command := &commands.Product{}
	c.GetContent(command, r)

	usecase.Execute(command)

	c.SendJSON(
		w,
		command,
		http.StatusOK,
	)
}
