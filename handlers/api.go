package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
)

// @Summary List Axioms
// @Schemes
// @Description returns a list of XGH methodology axioms
// @Tags List Axioms
// @Accept json
// @Produce json
// @Success 200 {array} gohorse.Axiom
// @Router /axioms [get]
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

// @Summary Get Axiom
// @Schemes
// @Description returns the XGH methodology axiom by its number
// @Tags Get Axiom
// @Accept json
// @Produce json
// @Success 200 {object} gohorse.Axiom
// @Failure 404 {object} ResponseError
// @Param axiomNumber path int true "Axiom Number"
// @Router /axioms/{axiomNumber} [get]
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

// @Summary Get Random Axiom
// @Schemes
// @Description returns the XGH methodology axiom randomly
// @Tags Get Axiom
// @Accept json
// @Produce json
// @Success 200 {object} gohorse.Axiom
// @Router /random [get]
func RandomAxiomHandlerJson(c *gin.Context) {
	c.AddParam("id", "random")
	AxiomDetailsHandlerJson(c)
}

type ResponseError struct {
	Error string
}
