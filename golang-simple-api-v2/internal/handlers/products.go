package handlers

import (
	"golang-simple-api-v2/internal/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	}
	if request.Method == http.MethodPost {
		p.addProduct(writer, request)
		return
	}

	if request.Method == http.MethodPut {
		p.logger.Println("PUT", request.URL.Path)

		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(group) != 1 {
			p.logger.Println("Invalid URI more than one ID")
			http.Error(writer, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.logger.Println("Invalid URI more than one capture group")
			http.Error(writer, "Invalid url", http.StatusBadRequest)
			return
		}

		idString := group[0][1]

		id, err := strconv.Atoi(idString)

		if err != nil {
			p.logger.Println("Invalid URI unable to convert to number")
			http.Error(writer, "Invalid url", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, writer, request)

		p.logger.Println("Got id", id)
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(writer)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle POST Product")

	product := &data.Product{}
	err := product.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *Products) updateProducts(id int, writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle PUT Product")

	product := &data.Product{}
	err := product.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(writer, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, "Product not found", http.StatusBadRequest)
		return
	}
}
