package sales_and_distribution_repo

import "go.mongodb.org/mongo-driver/mongo"

type userRepository struct {
	database       *mongo.Database
	collectionUser string
}
