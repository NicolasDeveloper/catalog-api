package controllers

import (
	"net/http"

	"github.com.br/NicolasDeveloper/catalog-api/api/common"
)

//ProductRouter handle device resources
type ProductRouter struct {
	routes []common.Route
}

//NewProductRouter list of routes
func NewProductRouter() common.Bundle {
	ctrl := NewProductController()

	r := []common.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: ctrl.Post,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: ctrl.Index,
		},
	}

	return &ProductRouter{
		routes: r,
	}
}

//GetRoutes implement interface common.Bundle
func (b *ProductRouter) GetRoutes() []common.Route {
	return b.routes
}
