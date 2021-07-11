package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pecolynx/casbin-test/pkg/domain"
	"github.com/pecolynx/casbin-test/pkg/gateway"
)

type petEntity struct {
	Name string `json:"name"`
}

func NewPetEntity(pet domain.Pet) petEntity {
	return petEntity{
		Name: pet.GetName(),
	}
}

type PetHandler interface {
	FindPets(c *gin.Context)
}

type petHandler struct {
	petRepository gateway.PetRepository
}

func NewPetHandler(petRepository gateway.PetRepository) PetHandler {
	return &petHandler{
		petRepository: petRepository,
	}
}

func (h *petHandler) FindPets(c *gin.Context) {
	ctx := c.Request.Context()
	name := c.Query("name")
	petModels, err := h.petRepository.FindPets(ctx, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	petEntities := make([]petEntity, len(petModels))
	for i, m := range petModels {
		petEntities[i] = NewPetEntity(m)
	}

	c.JSON(http.StatusOK, petEntities)
}
