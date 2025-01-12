package main

import (
	"Go_Practic/src/domain"
	"fmt"
)

func main(){
	soda := domain.NewProduct("coca-cola",17.50)
	fmt.Println(soda.GetName())

	soda.SetName("sabritas")
	fmt.Println(soda.GetName())
}


