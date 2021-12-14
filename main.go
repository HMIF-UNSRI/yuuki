package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"yuuki/app/mariadb"
	"yuuki/pkg/config"
)

func main() {
	configuration := config.NewConfiguration(`./.env`)
	mariadb.GetConnection(configuration)

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Yuuki!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
