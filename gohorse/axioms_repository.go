package gohorse

import (
	"fmt"
	"math/rand"
)

type AxiomsRepository struct {
	axiomsList map[int]Axiom
}

func NewAxiomsRepository(language string) AxiomsRepository {
	if language == "" {
		language = DEFAULT_LANGUAGE
	}

	return AxiomsRepository{
		axiomsList: AXIOMS[language],
	}
}

func (repo *AxiomsRepository) GetByNumber(number int) (*Axiom, error) {
	axiom, exists := repo.axiomsList[number]
	if !exists {
		return nil, NewAxiomNotFoundError(fmt.Sprintf("Axiom by number %d not found", number))
	}

	return &axiom, nil
}

func (repo *AxiomsRepository) GetRandom() Axiom {
	listSize := len(repo.axiomsList)
	randomAxiomNumber := rand.Intn(listSize)
	for randomAxiomNumber == 0 {
		randomAxiomNumber = rand.Intn(listSize)
	}
	return repo.axiomsList[randomAxiomNumber]
}

func (repo *AxiomsRepository) GetAll() map[int]Axiom {
	return repo.axiomsList
}
