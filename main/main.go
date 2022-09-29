package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello world")

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
	})

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":8085", nil)
}
