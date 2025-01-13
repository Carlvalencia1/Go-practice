package aplication

import "Go_Practic/src/domain"

type CreateProductUseCase struct {
	db domain.Iproduct
}

func NewCreateProductUseCase(db domain.Iproduct) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (uc CreateProductUseCase) Run ()  {
    uc.db.Save()
}