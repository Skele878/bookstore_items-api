package app

import (
	"net/http"
	"time"

	"github.com/Skele878/bookstore_items-api/src/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := (srv.ListenAndServe()); err != nil {
		panic(err)
	}
}
