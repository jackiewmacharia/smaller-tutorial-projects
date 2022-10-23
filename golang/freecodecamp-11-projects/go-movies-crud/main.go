package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // * is a pointer
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // slice - a dynamic array

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // replaced target with the rest of the array
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie) // from user input

	movie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// delete ID then append
	w.Header().Set("COntent-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"] // set same ID as deleted
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	muxRouter := mux.NewRouter()

	// &Director - gives the reference object/address
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie One", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "54678", Title: "Movie Two", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	muxRouter.HandleFunc("/movies", getMovies).Methods("GET")
	muxRouter.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	muxRouter.HandleFunc("/movies", createMovie).Methods("POST")
	muxRouter.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	muxRouter.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}
