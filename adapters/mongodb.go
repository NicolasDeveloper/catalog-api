package adapters

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbContext database context
type DbContext struct {
	client *mongo.Client
}

//NewDbContext database context
func NewDbContext() *DbContext {
	return &DbContext{}
}

//Connect method to connect in database
func (c *DbContext) Connect() error {
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")

	if connectionString == "" {
		connectionString = "mongodb://nicolas:Ni684250102@ds042128.mlab.com:42128/7180?retryWrites=false"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	c.client = client
	fmt.Println("Connected to MongoDB Catalog App!")

	return err
}

func getCollectionName(myvar interface{}) string {
	t := reflect.TypeOf(myvar)

	if t.Kind() == reflect.Ptr {
		return strings.ToLower(t.Elem().Name() + "s")
	}

	return strings.ToLower(t.Name() + "s")
}

//GetCollection retrive collection from mongo database based in struct instance name
func (c *DbContext) GetCollection(structInstance interface{}) (*mongo.Collection, error) {
	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		databaseName = "7180"
	}

	collection := c.client.Database(databaseName).Collection(getCollectionName(structInstance))

	return collection, nil
}

//GetCtx retrive database context
func (c *DbContext) GetCtx(structInstance interface{}) (*mongo.Client, error) {
	return c.client, nil
}
