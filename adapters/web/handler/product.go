package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/moreirak14/fullcycle-arquitetura-hexagonal/application"
	"net/http"
)

func MakeProductHandlers(
	router *mux.Router,
	n *negroni.Negroni,
	service application.ProductServiceInterface) {

	router.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
