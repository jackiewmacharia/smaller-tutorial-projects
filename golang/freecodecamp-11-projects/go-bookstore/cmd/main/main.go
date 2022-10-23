package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackiewmacharia/smaller-tutorial-projects/golang/freecodecamp-11-projects/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port 9010")
	log.Fatal((http.ListenAndServe(":9010", nil)))
}
