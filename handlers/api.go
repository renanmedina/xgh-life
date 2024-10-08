package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
)

func AxiomsListHandler(c *gin.Context) {
	repository := gohorse.NewAxiomsRepository(c.GetString("language"))
	c.JSON(200, repository.GetAll())
}

func fetchAxiomHandler(language string, id string) (*gohorse.Axiom, int, error) {
	use_case := gohorse.NewGetAxiomUseCase(language)

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
	axiom, statusCode, err := fetchAxiomHandler(c.GetString("language"), c.Param("id"))

	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, axiom)
}

func AxiomDetailsHandlerHtml(c *gin.Context) {
	axiom, statusCode, err := fetchAxiomHandler(c.GetString("language"), c.Param("id"))

	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	autoplayUrl := "/assets/static/audios/horse_sound.mp3"
	shouldPlayVoice, err := strconv.Atoi(c.Query("voice"))

	if err == nil && shouldPlayVoice == 1 {
		autoplayUrl = axiom.AudioUrl()
	}

	c.HTML(http.StatusOK, "horse_axiom.tmpl", gin.H{
		"axiom":       axiom,
		"autoplayUrl": autoplayUrl,
	})
}
