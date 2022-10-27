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

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies (w http.ResponseWriter, r *http.Request) {

	//set json Content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params :=mux.Vars(r)
	//loop over the movies , rang
	//delete the movie with the id you sent
	//add a new movie that we send in the body of postman 
	for index , item := range movies {
		if item.Id==params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id=params["id"]
			movies=append(movies , movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "438227", Title: "movie One", Director: &Director{Firstname: "hadis", Lastname: "rastegar"}})
	movies = append(movies, Movie{Id: "2", Isbn: "45455", Title: "movie Two", Director: &Director{Firstname: "kave", Lastname: "haj"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("Get")
	r.HandleFunc("/movies", createMovies).Methods("Post")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("Put")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("Delete")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
