package evaluation

import (
	"github.com/guzmanmo/go-imdb/internal/app/repository"
)

type EvaluationService struct {
	evaluationRepository *repository.EvaluationRepository
	movieRepository      *repository.MovieRepository
}

type EvaluationServiceInterface interface {
	CreateEvaluation(idMovie uint64, newEvaluation repository.Evaluation) (repository.Evaluation, error)
	GetAllEvaluation(idMovie uint64) ([]repository.Evaluation, error)
}

func NewEvaluationService(evaluationRepository *repository.EvaluationRepository, movieRepository *repository.MovieRepository) *EvaluationService {
	return &EvaluationService{
		evaluationRepository: evaluationRepository,
		movieRepository:      movieRepository,
	}
}

func (repo *EvaluationService) GetAllEvaluation(idMovie uint64) ([]repository.Evaluation, error) {
	movie, error := repo.movieRepository.FindOne(idMovie)
	if error != nil {
		return nil, error
	}
	resMovies, err := repo.evaluationRepository.GetAll(movie.ID)
	if err != nil {
		return []repository.Evaluation{}, err
	}
	return resMovies, nil
}

func (repo *EvaluationService) CreateEvaluation(idMovie uint64, newEvaluation repository.Evaluation) (repository.Evaluation, error) {
	_, error := repo.movieRepository.FindOne(idMovie)
	if error != nil {
		return repository.Evaluation{}, error
	}
	newEvaluation.MovieFk = int64(idMovie)
	newEvaluation.ID = 0

	resEvaluation, err := repo.evaluationRepository.CreateEvaluation(newEvaluation)
	if err != nil {
		return repository.Evaluation{}, err
	}
	return resEvaluation, nil
}
