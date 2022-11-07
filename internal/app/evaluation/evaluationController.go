package evaluation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guzmanmo/go-imdb/internal/app/repository"
)

type EvaluationController struct {
	evaluationService EvaluationServiceInterface
}

func NewEvaluationController(service EvaluationServiceInterface) *EvaluationController {
	return &EvaluationController{
		evaluationService: service,
	}
}

func (evaluationServices *EvaluationController) GetAllEvaluation(context *gin.Context) {

	idParam := context.Param("id")
	idMovie, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	evaluations, err := evaluationServices.evaluationService.GetAllEvaluation(idMovie)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, evaluations)
	}

}

func (evaluationServices *EvaluationController) CreateEvaluation(context *gin.Context) {

	idParam := context.Param("id")

	var newEvaluation repository.Evaluation
	err := context.ShouldBindJSON(&newEvaluation) // here should be good use a DTO instead but ... you know ..

	idMovie, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	evaluations, err := evaluationServices.evaluationService.CreateEvaluation(idMovie, newEvaluation)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, evaluations)
	}

}
