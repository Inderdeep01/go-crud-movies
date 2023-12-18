package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, mov := range movies {
		if mov.ID == params["id"] {
			json.NewEncoder(w).Encode(mov)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, mov := range movies {
		if mov.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
	return
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)
	for i, mov := range movies {
		if mov.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			json.NewEncoder(w).Encode("Movie deleted")
			break
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "2345", Title: "Movie 1", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "2346", Title: "Movie 2", Director: &Director{Firstname: "Jassa", Lastname: "Jatt"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	http.ListenAndServe(":8000", r)
}
