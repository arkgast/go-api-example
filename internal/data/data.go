package data

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type MovieInput struct {
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type MovieStore struct {
	movies map[string]Movie
	mu     sync.Mutex
}

func NewMovieStore() *MovieStore {
	return &MovieStore{
		movies: make(map[string]Movie),
	}
}

func (ms *MovieStore) Lock() {
	ms.mu.Lock()
}

func (ms *MovieStore) Unlock() {
	ms.mu.Unlock()
}

func (ms *MovieStore) GetMovies() map[string]Movie {
	return ms.movies
}

func (ms *MovieStore) GetMovie(id string) (*Movie, error) {
	for _, item := range ms.movies {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("Movie with ID: %s not found", id)
}

func (ms *MovieStore) CreateMovie(newMovie MovieInput) *Movie {
	movie := Movie{
		ID:    randomNumber(100, 10),
		Isbn:  newMovie.Isbn,
		Title: newMovie.Title,
		Director: &Director{
			Firstname: newMovie.Director.Firstname,
			Lastname:  newMovie.Director.Lastname,
		},
	}

	ms.movies[movie.ID] = movie

	return &movie
}

func (ms *MovieStore) UpdateMovie(id string, movieInput MovieInput) (*Movie, error) {
	movie, exists := ms.movies[id]
	if !exists {
		return nil, fmt.Errorf("Movie with ID: %s not found", id)
	}

	if movieInput.Isbn != "" {
		movie.Isbn = movieInput.Isbn
	}

	if movieInput.Title != "" {
		movie.Title = movieInput.Title
	}

	if movieInput.Director != nil {
		if movieInput.Director.Firstname != "" {
			movie.Director.Firstname = movieInput.Director.Firstname
		}
		if movieInput.Director.Lastname != "" {
			movie.Director.Lastname = movieInput.Director.Lastname
		}
	}

	ms.movies[id] = movie

	return &movie, nil
}

func (ms *MovieStore) DeleteMovie(id string) error {
	_, exists := ms.movies[id]
	if !exists {
		return fmt.Errorf("Movie with ID: %s not found", id)
	}

	delete(ms.movies, id)
	return nil
}

func randomNumber(max int, min int) string {
	randNumber := rand.Intn(max-min) + min
	return strconv.Itoa(randNumber)
}
