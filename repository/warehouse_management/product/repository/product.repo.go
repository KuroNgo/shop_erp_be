package product_repository

import "go.mongodb.org/mongo-driver/mongo"

type productRepository struct {
	database          *mongo.Database
	productCollection string
}

func NewProductRepository(database *mongo.Database, productCollection string) {

}
