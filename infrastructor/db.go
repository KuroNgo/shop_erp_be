package infrastructor

import (
	"context"
	"fmt"
	mongo_driven "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"shop_erp_mono/bootstrap"
	role_repo "shop_erp_mono/repository/human_resource_management/role/data_seeder"
	user_repo "shop_erp_mono/repository/human_resource_management/user/data_seeder"
	"time"
)

//func NewMongoDatabase(env *bootstrap.Database) *mongo_driven.Client {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	//dbHost := env.DBHost
//	//dbPort := env.DBPort
//	dbUser := env.DBUser
//	dbPass := env.DBPassword
//
//	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ykpyhgp.mongodb.net/?authMechanism=SCRAM-SHA-1", dbUser, dbPass)
//
//	//if dbUser == "" || dbPass == "" {
//	//	mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
//	//}
//
//	mongoCon := options.Client().ApplyURI(mongodbURI)
//	client, err := mongo_driven.Connect(ctx, mongoCon)
//	if err != nil {
//		log.Fatal("error while connecting with mongo", err)
//	}
//
//	err = client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal("error while trying to ping mongo", err)
//	}
//
//	session, err := client.StartSession()
//	if err != nil {
//		log.Fatal("err to start sessions", err)
//	}
//	defer session.EndSession(context.Background())
//
//	err = session.StartTransaction()
//	if err != nil {
//		log.Fatal("err to start sessions", err)
//	}
//	defer func() {
//		if err != nil {
//			// Rollback giao dịch
//			err := session.AbortTransaction(context.Background())
//			if err != nil {
//				return
//			}
//			return
//		}
//		// Commit giao dịch
//		err := session.CommitTransaction(context.Background())
//		if err != nil {
//			return
//		}
//	}()
//	return client
//}

// unit_test
func NewMongoDatabase(env *bootstrap.Database) *mongo_driven.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf("mongodb://localhost:27017/")

	mongoCon := options.Client().ApplyURI(mongodbURI)
	client, err := mongo_driven.Connect(ctx, mongoCon)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	// migration
	err = Migrations(ctx, client)
	if err != nil {
		return nil
	}

	return client
}

func CloseMongoDBConnection(client *mongo_driven.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func Migrations(ctx context.Context, client *mongo_driven.Client) error {
	// migration
	err := user_repo.SeedUser(ctx, client)
	if err != nil {
		return nil
	}

	err = role_repo.SeedRole(ctx, client)
	if err != nil {
		return err
	}

	return nil
}
