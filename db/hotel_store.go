package db

import (
	"context"

	"github.com/PratikVarute/BookMyStay_Go_Fiber/literals"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsterHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
}

type MongoHotelStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoHotelStore(Client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		Client: Client,
		Coll:   Client.Database(literals.DbName).Collection(literals.HotelColl),
	}
}

func (m *MongoHotelStore) InsterHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := m.Coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (m *MongoHotelStore) Update(ctx context.Context, filter, update bson.M) error {
	_, err := m.Coll.UpdateOne(ctx, filter, update)
	return err
}
