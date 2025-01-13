package infraestructure

import "fmt"

type MySql struct {
}

func NewMySql() *MySql {
	return &MySql{}
}

func (mysql *MySql) Save() {
    fmt.Println("Metodo Save")
}

func (mysql *MySql) GetAll() {
	fmt.Println("Metodo Get")
}