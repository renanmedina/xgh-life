package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
)

type CommandRequest struct {
	CommandType   string `json:"command" binding:"required"`
	RequestedType string `json:"text" binding:"required"`
}

func (c CommandRequest) isXGH() bool {
	return c.CommandType == "/xgh"
}

func (c CommandRequest) isRandomAxiomType() bool {
	return c.RequestedType == "random"
}

func (c CommandRequest) axiomNumber() (int, error) {
	number, err := strconv.Atoi(c.RequestedType)
	if err != nil {
		return 0, errors.New("Invalid axiom number parameter")
	}

	return number, nil
}

func SlackBotHandler(c *gin.Context) {
	var request CommandRequest
	c.BindJSON(&request)

	if !request.isXGH() {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Command %s not available", request.CommandType)})
		return
	}

	service := gohorse.NewAxiomsService()

	if request.isRandomAxiomType() {
		axiom := service.PickOneRandom()
		c.JSON(http.StatusOK, axiom)
		return
	}

	axiom_number, err := request.axiomNumber()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	axiom, err := service.GetByNumber(axiom_number)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, axiom)
}
