package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	mongoURI = "mongodb://localhost:27017"
)

type Traning struct {
	Name  string
	Phone int32
}

func main() {
	//создание клиента
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	//создание подключения
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// проверка подключения
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("sibcoder").Collection("test") // обращение к коллекции
	fmt.Println(collection)

	//структуры для добавления
	dan := Traning{"Faruf", 321269}
	sin := Traning{"Islam", 321243}
	das := Traning{"Alan", 321256}
	//одиночная вставка
	insertResult, err := collection.InsertOne(context.TODO(), dan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//множественная вставка
	trainers := []interface{}{sin, das}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// обновление документов
	filter := bson.D{{"name", "Alan"}}

	update := bson.D{
		{"$inc", bson.D{
			{"phone", 1},
		}}, //увеличение значения по ключу phone на 1
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//поиск документов
	var result Traning
	filter = bson.D{{"name", "Alan"}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result) //Чтобы найти несколько документов, используйте collection.Find()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	//удаление документов collection.DeleteOne() или collection.DeleteMany() или collection.Drop() для удаления всей коллекции
	filter = bson.D{{"name", "Alan"}}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//закрытие подключения
	err = client.Disconnect(context.TODO()) //закрытие подключения
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
