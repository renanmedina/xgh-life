package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
)

var service gohorse.AxiomsService

func init() {
	service = gohorse.NewAxiomsService()
}

func AxiomsListHandler(c *gin.Context) {
	c.JSON(200, service.GetAxiomsList())
}

func AxiomDetailsHandler(c *gin.Context) {
	axiomNumber, err := strconv.Atoi(c.Param("number"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid axiom number parameter"})
		return
	}

	axiom, err := service.GetByNumber(axiomNumber)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, axiom)
}
