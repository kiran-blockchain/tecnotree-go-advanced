package interfaces

import "github.com/kiran-blockchain/ekart/entities"

type IProduct interface{
	Insert(product *entities.Product) (string,error)
	GetProducts() ([]*entities.Product,error)
	UpdateProduct(product *entities.Product) (string,error)

}