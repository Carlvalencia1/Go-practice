package domain

type product struct {
    
	id int
	name string
    price float32

}

func NewProduct(name string, price float32) *product  {
    obj := product{id: 1, name: name, price: price}
	return &obj
	//return &product{ id: 1, name: name, price: price}
}