package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PratikVarute/BookMyStay_Go_Fiber/db"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/literals"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(literals.DbUrl))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Database(literals.DbName).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	roomStroe := db.NewMongoRoomStore(client, hotelStore)

	hotel := types.Hotel{
		Name:     "Agonda Holidays",
		Location: "Goa",
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []types.Room{
		{
			Type:  types.SingleRoom,
			Price: 1000.00,
		},
		{
			Type:  types.SeeSiteRoom,
			Price: 4000.00,
		},
		{
			Type:  types.DeluxRoom,
			Price: 10000.00,
		},
	}

	insertedHotel, err := hotelStore.InsterHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insteredRoom, err := roomStroe.InsterRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("New resgistered room details :", insteredRoom)
	}
	fmt.Println("New resgistered hotel details :", insertedHotel)

}
