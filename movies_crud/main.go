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

// Movie struct represents a movie with an ID, ISBN, Title, and Director
type Movie struct {
	ID       string   `json:"id"`
	ISBN     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents a director with a Firstname and Lastname
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Initialize a slice to store movies
var movies []Movie

// GetMovies responds with the list of all movies as JSON
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// DeleteMovie deletes a movie by its ID and responds with the updated list of movies
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// GetMovie responds with a single movie by its ID as JSON
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateMovie creates a new movie and responds with the created movie as JSON
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000)) // Generate a random ID for the new movie
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// UpdateMovie updates an existing movie by its ID and responds with the updated movie as JSON
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies[index] = movie
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Seed the movies slice with some initial data
	movies = append(movies, Movie{ID: "1", ISBN: "1234567890", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Joe", Lastname: "Schmoe"}})
	movies = append(movies, Movie{ID: "2", ISBN: "9876543210", Title: "The Godfather", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})

	// Define the endpoints and their corresponding handler functions
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	// Start the server and listen on port 8080
	fmt.Printf("Server is running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
