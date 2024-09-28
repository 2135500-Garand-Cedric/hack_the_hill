package database

import (
	// "context"
	"fmt"
	"log"

	"github.com/wkirk01/AlgoeDB"

)

type UsersDB []map[string]interface{}
type User map[string]interface{}
func GetDB() *AlgoeDB.Database {

	config := AlgoeDB.DatabaseConfig{Path: "./db/users.json"}
	db, err := AlgoeDB.NewDatabase(&config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
	

// people := People{}
// people = append(people, Person{"name": "Billy", "age": 27})
// people = append(people, Person{"name": "Carisa", "age": 26})

// err = db.InsertMany(people)
// if err != nil {
// 	log.Fatal(err)
// }	


func FindUserByEmail(db *AlgoeDB.Database, email string) (User, error) {
	

	query := User{"email": email}

	result := db.FindOne(query)

	if result == nil {
		return User{}, fmt.Errorf("user not found")
	}

	return result, nil

}

// func findEntryByID(client *mongo.Client, id string) (bson.M, error) {
// 	// Access the collection
// 	collection := client.Database(DatabaseName).Collection(DatabaseCollection)

// 	// Perform the find query
// 	var result bson.M
// 	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func DoesUserExist(client *mongo.Client, email string) (bool, error) {
// 	// Access the collection
// 	collection := client.Database(DatabaseName).Collection(DatabaseCollection)

// 	// Perform the find query
// 	var result bson.M
// 	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
// 	if err != nil {
// 		return false, nil
// 	}

// 	return true, nil
// }

func InsertUser(db *AlgoeDB.Database, user User) error {

	err := db.InsertOne(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}