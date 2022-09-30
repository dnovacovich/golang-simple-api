package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	logger *log.Logger
}

func NewGoodBye(logger *log.Logger) *GoodBye {
	return &GoodBye{logger}
}

func (goodBye *GoodBye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Byeeee"))
}
