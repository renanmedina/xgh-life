package gohorse

import (
	"errors"
	"strconv"
)

type GetAxiomUseCase struct {
	repository AxiomsRepository
}

func NewGetAxiomUseCase(language string) GetAxiomUseCase {
	return GetAxiomUseCase{
		repository: NewAxiomsRepository(language),
	}
}

const RANDOM_OPTION = "random"

func (use_case *GetAxiomUseCase) Execute(option string) (Axiom, error) {
	if option == RANDOM_OPTION {
		return use_case.repository.GetRandom(), nil
	}

	number, err := strconv.Atoi(option)
	if err != nil {
		return Axiom{}, errors.New("Invalid axiom number parameter")
	}

	axiom, err := use_case.repository.GetByNumber(number)

	if err != nil {
		return Axiom{}, err
	}

	return *axiom, nil
}
