package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)

	return encoder.Encode(p)
}

func (product *Product) FromJSON(reader io.Reader) error {
	encoder := json.NewDecoder(reader)

	return encoder.Decode(product)
}

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) {
	product.ID = getNextID()

	productList = append(productList, product)
}

func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	product.ID = id
	productList[pos] = product

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]

	return lastProduct.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Description",
		Price:       2.25,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Description",
		Price:       1.99,
		SKU:         "abc124",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
