package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guzmanmo/go-imdb/internal/app"
	"github.com/guzmanmo/go-imdb/internal/app/evaluation"
	"github.com/guzmanmo/go-imdb/internal/app/movie"
	"github.com/guzmanmo/go-imdb/internal/app/repository"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

//var router *mux.Router

var router *gin.Engine

var DbConnection *gorm.DB

func main() {
	router = gin.Default()
	DbConnection = app.SetUpConnection()
	setUpRoutes()
	StartServer()
	router.Run()
}

func setUpRoutes() {

	movieRepository := repository.NewMovieRepository(DbConnection)
	movieService := movie.NewMovieService(movieRepository)
	movieController := movie.NewMovieController(movieService)

	router.POST("/api/movie", movieController.CreateMovie)
	router.GET("/api/movie/:id", movieController.GetOneMovie)
	router.GET("/api/movie", movieController.GetAllMovies)

	evaluationRepository := repository.NewEvaluationRepository(DbConnection)
	evaluationService := evaluation.NewEvaluationService(evaluationRepository, movieRepository)
	evaluationController := evaluation.NewEvaluationController(evaluationService)

	router.GET("/api/movie/:id/evaluation", evaluationController.GetAllEvaluation)
	router.POST("/api/movie/:id/evaluation", evaluationController.CreateEvaluation)
}

func StartServer() {

	fmt.Printf("Server initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
