package repository 

import (
	"context"
	"log"
	"fmt"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Belyakoff/apartmentservice/data"

)



func GetItems(collection *mongo.Collection)(data.Apartments){

	filter := bson.M{}

	return FindRequestToMongoDB(collection, filter)
}


func GetItemsByPrice(collection *mongo.Collection, price int, op string)(data.Apartments){

	filter := bson.M{"price":bson.M{op:price}}

	return FindRequestToMongoDB(collection, filter)
}


func FindRequestToMongoDB(collection *mongo.Collection, filter map[string]interface{})(data.Apartments){

	var results []* data.Apartment

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
    	log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

    
    	var elem data.Apartment
    	err := cur.Decode(&elem)
    	if err != nil {
        	log.Fatal(err)
    	}

    	results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
    	log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results
}


func InsertItem(collection *mongo.Collection, item data.Apartment){

	

	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
        	log.Fatal(err)
    }

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	

}
