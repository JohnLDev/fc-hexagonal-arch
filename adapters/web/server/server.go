package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/johnldev/go-hexagonal/adapters/web/handler"
	"github.com/johnldev/go-hexagonal/app"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service app.ProductServiceInterface
}

func NewWebServer(service app.ProductServiceInterface) *WebServer {
	return &WebServer{Service: service}
}

func (w *WebServer) Serve() {

	router := mux.NewRouter()
	loggerMiddleware := negroni.New(negroni.NewLogger())
	handler.MakeProductHandlers(router, loggerMiddleware, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		Addr:              ":8080",
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
