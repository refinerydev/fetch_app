package product

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var products = []ProductEntity{}

type Repository interface {
	FetchProduct()
}

type repository struct {
	products *[]ProductEntity
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) FetchProduct() {
	url := "https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	fmt.Println(r.products)
	fmt.Println(products)
}
