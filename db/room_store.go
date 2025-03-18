package db

import (
	"context"

	"github.com/PratikVarute/BookMyStay_Go_Fiber/literals"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsterRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
	HotelStore
}

func NewMongoRoomStore(Client *mongo.Client, HotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		Client:     Client,
		Coll:       Client.Database(literals.DbName).Collection(literals.RoomsColl),
		HotelStore: HotelStore,
	}
}

func (m *MongoRoomStore) InsterRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := m.Coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = resp.InsertedID.(primitive.ObjectID)
	//update hotel with this room id
	// The lines `filter := bson.M{"_id": room.HotelID}` and `update := bson.M{"": bson.M{"rooms":
	// room.ID}}` are creating MongoDB update operations to associate a room with a hotel in a MongoDB
	// collection.
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}
	if err = m.Update(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}
