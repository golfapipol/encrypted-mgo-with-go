package main

import (
	"encryptedmgo/model"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	DBConnection, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	todoCollection := DBConnection.Clone().DB("test").C("todos")

	err = todoCollection.Insert(model.Todo{
		Name:          "hello",
		EncryptedName: "world",
	})

	if err != nil {
		log.Fatal("Cannot Insert", err.Error())
	}

	var expected model.Todo

	todoCollection.Find(bson.M{"name": "hello"}).One(&expected)
	log.Println("query name:", expected.Name, "encrypt name:", expected.EncryptedName)
}
