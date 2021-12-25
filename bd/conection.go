//nolint
package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://andremor:admin@cluster0.gw736.mongodb.net/twitter?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa")
	return client
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
