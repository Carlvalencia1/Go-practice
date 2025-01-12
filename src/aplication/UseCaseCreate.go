package aplication

import "Go_Practic/src/domain"

type useCaseCreate struct {
}

func NewUseCaseCreate() *useCaseCreate {
	return &useCaseCreate{}
}

func (uc *useCaseCreate) Run () *(domain.Product) {
	return &domain.Product{}
}