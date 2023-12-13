package initializers

import (
	"categori/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/glebarez/sqlite"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var collection *mongo.Collection

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("iplocs.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	Database = DbInstance{
		Db: db,
	}
}

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blacklist").Collection("categories")
}

func GetCollection() *mongo.Collection {
	return collection
}

func GetLocations() ([]model.Location, error) {
	var locations []model.Location
	result := Database.Db.Table("iplocations").Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}

	return locations, nil
}

func PrintLocations(locations []model.Location) {
	jsonEncoded, _ := json.Marshal(&locations)
	fmt.Println(string(jsonEncoded))
}
