package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type movie struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
} 

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type Movies []movie
var movies []movie
func init() {
	movies = Movies {
		movie{
			ID: "1",
			ISBN: "3456778", 
			Title: "A good day to die hard",
			Director: &Director{Firstname: "John", Lastname:"Doe"},
		},

		movie{
			ID: "2",
			ISBN: "2345687980", 
			Title: "Die hard",
			Director: &Director{Firstname: "John", Lastname:"Snow"},
		},
	}
}

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Movie movie 
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, Movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var Movie movie
			_= json.NewDecoder(r.Body).Decode(&Movie)
			Movie.ID = params["id"]
			movies = append(movies, Movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func deleteMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}	
	json.NewEncoder(w).Encode(movies)
}

func main() {
 r := mux.NewRouter()
 r.HandleFunc("/getmovies", getMovies).Methods("GET")	
 r.HandleFunc("/getmovie/{id}", getMovie).Methods("GET")	
 r.HandleFunc("/getmovies", createMovies).Methods("POST")	
 r.HandleFunc("/getmovies/{id}", updateMovies).Methods("PUT")	
 r.HandleFunc("/getmovies/{id}", deleteMovies).Methods("DELETE")
 
 fmt.Println("Starting Server at port 8080")
 err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, r)
 if err != nil{
	 log.Fatal("Error starting server at port 8080: ", err)
	 return
 }
}