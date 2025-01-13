package main

import (
	"Go_Practic/src/aplication"
	"Go_Practic/src/infraestructure"
	
)

func main(){
	mysql := infraestructure.NewMySql()
	CreateProductUseCase :=aplication.NewCreateProductUseCase(mysql)

	CreateProductUseCase.Run()
}


