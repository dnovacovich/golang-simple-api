package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	hello.logger.Println("Hello world")

	d, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error detected"))

		return
		// Can be replaced with:
		// http.Error(writer, "Error detected", http.StatusBadRequest)
		// return
	}

	fmt.Fprintf(writer, "Hello %s", d)
}
