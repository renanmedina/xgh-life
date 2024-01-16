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

func AxiomDetailsHandler(c *gin.Context) {
	use_case := gohorse.NewGetAxiomUseCase()

	axiom, err := use_case.Execute(c.Param("number"))

	if err != nil {
		statusCode := http.StatusBadRequest
		if errors.As(err, &gohorse.AxiomNotFoundError{}) {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, axiom)
}
