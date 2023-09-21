package services

import (
	"context"

	"github.com/kiran-blockchain/ekart/entities"
	"github.com/kiran-blockchain/ekart/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct{
	Product *mongo.Collection
	
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct{
	
   return &ProductService{Product: collection}
}

func (p *ProductService)Insert(product *entities.Product)(string,error){
	product.ID= primitive.NewObjectID()
	_,err:=p.Product.InsertOne(context.Background(),product)
	 if err!=nil{
		return "",err
	 }else{
		return "Record Inserted Successfully",nil
	 }
}