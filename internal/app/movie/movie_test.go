package movie

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guzmanmo/go-imdb/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
}

func (mock *mockService) CreateMovie(newMovie repository.Movie) (repository.Movie, error) {
	return repository.Movie{
		ID: 1, Title: "El conjuro 1", Description: "A scary movie!",
	}, nil
}

func (mock *mockService) GetOneMovie(id uint64) (repository.Movie, error) {
	return repository.Movie{ID: 1, Title: "El conjuro 1", Description: "A scary movie!"}, nil
}
func (mock *mockService) GetAllMovies() ([]repository.Movie, error) {
	return []repository.Movie{
		{ID: 1, Title: "El conjuro 1", Description: "A scary movie!"},
	}, nil
}

func getTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	return ctx, w
}

func TestGetAllMovieController(t *testing.T) {
	c, res := getTestGinContext()
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	movieController := NewMovieController(&mockService{})
	movieController.GetAllMovies(c)
	var movieUnmarshall []repository.Movie
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &movieUnmarshall)
	assert.Equal(t, movieUnmarshall, []repository.Movie{
		{ID: 1, Title: "El conjuro 1", Description: "A scary movie!"},
	})
	assert.Equal(t, c.Writer.Status(), 200)
}

func TestGetOneMovieController(t *testing.T) {
	c, res := getTestGinContext()
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}
	movieController := NewMovieController(&mockService{})
	movieController.GetOneMovie(c)
	var movieUnmarshall repository.Movie
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &movieUnmarshall)
	assert.Equal(t, movieUnmarshall, repository.Movie{
		ID: 1, Title: "El conjuro 1", Description: "A scary movie!",
	})
	assert.Equal(t, c.Writer.Status(), 200)
}

func TestCreateMovieController(t *testing.T) {
	newMovie := repository.Movie{ID: 2, Title: "Scream 5", Description: "A scary movie!"}
	newMovieJson, _ := json.Marshal(newMovie)
	c, res := getTestGinContext()
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(bytes.NewBuffer(newMovieJson))
	movieController := NewMovieController(&mockService{})
	movieController.CreateMovie(c)
	var movieUnmarshall repository.Movie
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &movieUnmarshall)
	assert.Equal(t, movieUnmarshall, repository.Movie{
		ID: 1, Title: "El conjuro 1", Description: "A scary movie!",
	})
	assert.Equal(t, c.Writer.Status(), 200)
}
