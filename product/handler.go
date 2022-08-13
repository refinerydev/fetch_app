package product

import (
	"github.com/labstack/echo/v4"
	"github.com/refinerydev/fetch_app/helper"
)

func GetProducts(c echo.Context) error {
	repo := NewRepository()
	pservice := NewService(repo)
	pservice.FetchProducts()
	return c.JSON(200, helper.M{"messange": "products"})
}
