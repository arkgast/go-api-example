package data

import (
	"testing"

	"github.com/bxcodec/faker/v3"
)

func movieInputFactory() *MovieInput {
	movieInput := MovieInput{}
	err := faker.FakeData(&movieInput)
	if err != nil {
		panic(err)
	}

	return &movieInput
}

func TestCreateMovie(t *testing.T) {
	movieStore := NewMovieStore()

	movie := movieInputFactory()
	movieStore.CreateMovie(*movie)

	if len(movieStore.GetMovies()) != 1 {
		t.Errorf("Expected 1 movie, got %d", len(movieStore.GetMovies()))
	}
}

func TestGetMovie(t *testing.T) {
	movieStore := NewMovieStore()

	movie := movieInputFactory()
	createdMovie := movieStore.CreateMovie(*movie)

	movieResult, error := movieStore.GetMovie(createdMovie.ID)
	if error != nil {
		t.Errorf("Error getting movie: %s", error)
	}

	if movieResult.Title != movie.Title {
		t.Errorf("Expected %s, got %s", movie.Title, movieResult.Title)
	}
}

func TestGetMovies(t *testing.T) {
	movieStore := NewMovieStore()

	movie := movieInputFactory()
	movieStore.CreateMovie(*movie)
	movieStore.CreateMovie(*movie)

	movies := movieStore.GetMovies()

	if len(movies) != 2 {
		t.Errorf("Expected 1 movie, got %d", len(movies))
	}
}

func TestUpdateMovie(t *testing.T) {
	movieStore := NewMovieStore()

	movie := movieInputFactory()
	createdMovie := movieStore.CreateMovie(*movie)

	movie.Title = "Updated Title"
	updatedMovie, error := movieStore.UpdateMovie(createdMovie.ID, *movie)
	if error != nil {
		t.Errorf("Error updating movie: %s", error)
	}

	if updatedMovie.Title != movie.Title {
		t.Errorf("Expected %s, got %s", createdMovie.Title, updatedMovie.Title)
	}
}

func TestDeleteMovie(t *testing.T) {
	movieStore := NewMovieStore()

	movie := movieInputFactory()
	createdMovie := movieStore.CreateMovie(*movie)

	error := movieStore.DeleteMovie(createdMovie.ID)
	if error != nil {
		t.Errorf("Error deleting movie: %s", error)
	}

	if len(movieStore.GetMovies()) != 0 {
		t.Errorf("Expected 0 movies, got %d", len(movieStore.GetMovies()))
	}
}
