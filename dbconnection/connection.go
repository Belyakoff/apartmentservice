package dbconnection

import (
	"context"
	"log"
	"time"
	"fmt"
	//"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)


func connectMongoDB(dbname string)(*mongo.Collection){
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 	defer cancel()


 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://darius:"+pswd+"@cluster0.aghsw.mongodb.net/"+dbname+"?retryWrites=true&w=majority",
  	))
	if err != nil { log.Fatal(err) }

	err = client.Ping(context.TODO(), nil)
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database(dbname).Collection("apartments")

	return collection
}
	