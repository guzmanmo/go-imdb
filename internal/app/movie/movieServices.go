package movie

import (
	"github.com/guzmanmo/go-imdb/internal/app/repository"
)

type MovieService struct {
	movieRepository *repository.MovieRepository
}

type MovieServiceInterface interface {
	CreateMovie(newMovie repository.Movie) (repository.Movie, error)
	GetOneMovie(id uint64) (repository.Movie, error)
	GetAllMovies() ([]repository.Movie, error)
}

func NewMovieService(repository *repository.MovieRepository) *MovieService {
	return &MovieService{
		movieRepository: repository,
	}
}

func (repo *MovieService) CreateMovie(newMovie repository.Movie) (repository.Movie, error) {
	resMovie, err := repo.movieRepository.Create(newMovie)
	if err != nil {
		return repository.Movie{}, err
	} else {
		return resMovie, nil
	}
}

func (repo *MovieService) GetOneMovie(id uint64) (repository.Movie, error) {

	resMovie, err := repo.movieRepository.FindOne(id)
	if err != nil {
		return repository.Movie{}, err
	}
	return resMovie, nil

}

func (repo *MovieService) GetAllMovies() ([]repository.Movie, error) {

	resMovies, err := repo.movieRepository.GetAll()
	if err != nil {
		return []repository.Movie{}, err
	}
	return resMovies, nil
}
