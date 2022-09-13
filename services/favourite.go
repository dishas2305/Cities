package services

import (
	"context"
	"fmt"

	// "time"

	"cities/models"
	"cities/types"

	"cities/storage"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FavouritesReceiver struct {
	MDB *mongo.Database
	CId int

	CityPayload types.CityPayload
}

func (cr *FavouritesReceiver) AddFavouriteCity(cityID string) error {
	// var city models.CityModel
	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityID)
	if err != nil {
		logger.Error("func_AddFavouriteCity", err)

	}
	filter := bson.M{
		"_id": cId,
	}
	favourite := mdb.Collection(models.CitiesCollection).FindOne(context.TODO(), filter)
	_, err = mdb.Collection(models.FavouritesCollection).InsertOne(context.TODO(), favourite)
	// err = result.Decode(&city)
	// if err != nil {
	// 	logger.Error("func_S_GetGrant: Error in ", err)
	// 	return city, err
	// }
	return nil
}
func (cr *FavouritesReceiver) GetFavouriteCities() ([]models.FavouriteCityModel, error) {
	mdb := storage.MONGO_DB
	filter := bson.M{}
	var cities []models.FavouriteCityModel
	result, err := mdb.Collection(models.FavouritesCollection).Find(context.TODO(), filter)

	fmt.Println("result***", result, err)

	if err := result.All(context.Background(), &cities); err != nil {
		logger.Error("func_GetFavouriteCities: error cur.All() step ", err)
		return nil, err
	}

	fmt.Println("cities", cities, err)
	return cities, nil
}

func (cr *FavouritesReceiver) RemoveFavouriteCity(cityId string) error {

	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityId)
	if err != nil {
		logger.Error("func_RemoveFavouriteCity", err)

	}
	filter := bson.M{
		"_id": cId,
	}

	_, err = mdb.Collection(models.FavouritesCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
