package movie

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guzmanmo/go-imdb/internal/app/repository"
)

type MovieController struct {
	movieService MovieServiceInterface
}

func NewMovieController(service MovieServiceInterface) *MovieController {
	return &MovieController{
		movieService: service,
	}
}

func (movieServices *MovieController) GetAllMovies(context *gin.Context) {
	movies, err := movieServices.movieService.GetAllMovies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, movies)
	}

}

func (movieServices *MovieController) GetOneMovie(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	resMovie, err := movieServices.movieService.GetOneMovie(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, resMovie)
	}

}

func (movieServices *MovieController) CreateMovie(context *gin.Context) {
	var movieRes repository.Movie
	var newMovie repository.Movie
	err := context.ShouldBindJSON(&newMovie)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	movieRes, err = movieServices.movieService.CreateMovie(newMovie)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, movieRes)
	}
}
