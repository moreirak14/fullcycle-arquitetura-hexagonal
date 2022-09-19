package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/moreirak14/fullcycle-arquitetura-hexagonal/adapters/web/handler"
	"github.com/moreirak14/fullcycle-arquitetura-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	router := mux.NewRouter()
	logger := negroni.New(negroni.NewLogger())
	handler.MakeProductHandlers(router, logger, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
