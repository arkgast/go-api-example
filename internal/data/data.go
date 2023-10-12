package data

import (
	"fmt"
	"math/rand"
	"strconv"
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

func GetMovies() []Movie {
	return movies
}

func GetMovie(id string) (*Movie, error) {
	for _, item := range movies {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("Movie with %s not found", id)
}

func CreateMovie() *Movie {
	newMovie := Movie{
		ID:    strconv.Itoa(rand.Intn(10)),
		Isbn:  strconv.Itoa(rand.Intn(1000)),
		Title: "One",
		Director: &Director{
			Firstname: "Name",
			Lastname:  "Lastname",
		},
	}

	movies = append(movies, newMovie)

	return &newMovie
}

func DeleteMovie(id string) error {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Movie with id %s not found", id)
}

func UpdateMovie(id string) (*Movie, error) {
	idx := -1
	for index, movie := range movies {
		if movie.ID == id {
			idx = index
			break
		}
	}

	if idx == -1 {
		return nil, fmt.Errorf("Movie with id %s not found", id)
	}

	movie := movies[idx]
	movie.Isbn = fmt.Sprintf("%s - updated", movie.Isbn)

	movies[idx] = movie

	return &movies[idx], nil
}
