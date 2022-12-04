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
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to LMDB!")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			// delete the existing entry
			movies = append(movies[:i], movies[i+1:]...)

			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				log.Fatal(err)
				return
			}

			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Status", "200")

	params := mux.Vars(r)

	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
}

func main() {
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "1234",
		Title: "Shawshank Redemption",
		Director: &Director{
			FirstName: "Frank",
			Lastname:  "Darabont",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "1235",
		Title: "The Godfather 1",
		Director: &Director{
			FirstName: "Francis",
			Lastname:  "Ford Coppola",
		},
	})

	r := mux.NewRouter()

	r.HandleFunc("/", homePageHandler).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Print("Listening on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
