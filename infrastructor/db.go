package infrastructor

import (
	"context"
	"fmt"
	mongo_driven "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"shop_erp_mono/bootstrap"
	employeerepo "shop_erp_mono/repository/human_resource_management/employee/data_seeder"
	rolerepo "shop_erp_mono/repository/human_resource_management/role/data_seeder"
	salaryrepo "shop_erp_mono/repository/human_resource_management/salary/data_seeder"
	userrepo "shop_erp_mono/repository/human_resource_management/user/data_seeder"
	"time"
)

func NewMongoDatabase(env *bootstrap.Database) *mongo_driven.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPassword

	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@andrew.8ulkv.mongodb.net/?retryWrites=true&w=majority&appName=Andrew", dbUser, dbPass)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	mongoCon := options.Client().ApplyURI(mongodbURI)
	client, err := mongo_driven.Connect(ctx, mongoCon)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	session, err := client.StartSession()
	if err != nil {
		log.Fatal("err to start sessions", err)
	}
	defer session.EndSession(context.Background())

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
	err := userrepo.SeedUser(ctx, client)
	if err != nil {
		return nil
	}

	err = rolerepo.SeedRole(ctx, client)
	if err != nil {
		return err
	}

	err = salaryrepo.SeedSalary(ctx, client)
	if err != nil {
		return err
	}

	err = employeerepo.SeedEmployee(ctx, client)
	if err != nil {
		return err
	}
	return nil
}
