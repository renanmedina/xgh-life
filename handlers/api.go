package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
)

var repository gohorse.AxiomsRepository

func init() {
	repository = gohorse.NewAxiomsRepository()
}

func AxiomsListHandler(c *gin.Context) {
	c.JSON(200, repository.GetAll())
}

func fetchAxiomHandler(id string) (*gohorse.Axiom, int, error) {
	use_case := gohorse.NewGetAxiomUseCase()

	if id == "roulette" {
		id = gohorse.RANDOM_OPTION
	}

	axiom, err := use_case.Execute(id)

	if err != nil {
		statusCode := http.StatusBadRequest
		if errors.As(err, &gohorse.AxiomNotFoundError{}) {
			statusCode = http.StatusNotFound
		}

		return nil, statusCode, err
	}

	return &axiom, http.StatusOK, nil
}

func AxiomDetailsHandlerJson(c *gin.Context) {
	axiom, statusCode, err := fetchAxiomHandler(c.Param("id"))

	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, axiom)
}

func AxiomDetailsHandlerHtml(c *gin.Context) {
	axiom, statusCode, err := fetchAxiomHandler(c.Param("id"))

	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "horse_axiom.tmpl", gin.H{
		"axiom": axiom,
	})
}
