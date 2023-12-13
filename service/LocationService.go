package service

import (
	"categori/initializers"
	"categori/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetLocationByIPFromSQLLite(ipAddress int64) (*model.LocationResponse, error) {
	var location model.Location

	result := initializers.Database.Db.Table("iplocations").
		Where("ip_to <= ? AND ip_from >= ?", ipAddress, ipAddress).
		First(&location)

	if result.Error != nil {
		return nil, result.Error
	}
	response := &model.LocationResponse{
		IpTo:        location.IpTo,
		IpFrom:      location.IpFrom,
		CountryCode: location.CountryCode,
		CountryName: location.CountryName,
	}

	return response, nil
}

func GetLocationByDomainFromMongoDb(domain string) (*model.DomainCategoryResponse, error) {
	filter := bson.D{{"domains", domain}}

	var result model.DomainCategory
	err := initializers.GetCollection().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Fatal(err)
	}
	response := &model.DomainCategoryResponse{
		Domain:       domain,
		CountryName:  "",
		IpAddress:    "",
		CategoryName: result.CategoryName,
	}
	return response, nil

}
