package database

import (
	// "context"
	"fmt"
	"log"

	"github.com/wkirk01/AlgoeDB"

)

type UsersDB []map[string]interface{}
type User map[string]interface{}

type ProfilesDB []map[string]interface{}
type Profile map[string]interface{}

func GetDB() *AlgoeDB.Database {

	config := AlgoeDB.DatabaseConfig{Path: "./db/users.json"}
	db, err := AlgoeDB.NewDatabase(&config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetProfileDB() *AlgoeDB.Database {

	config := AlgoeDB.DatabaseConfig{Path: "./db/profiles.json"}	
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


func InsertProfile(db *AlgoeDB.Database, profile Profile) error {

	err := db.InsertOne(profile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


func FindUserByEmail(db *AlgoeDB.Database, email string) (User, error) {
	

	query := User{"email": email}

	result := db.FindOne(query)

	if result == nil {
		return User{}, fmt.Errorf("user not found")
	}

	return result, nil

}

func FindUserByUsername(db *AlgoeDB.Database, username string) (User, error) {

	
	query := User{"username": username}
	
	
	result := db.FindOne(query)
	
	if result == nil {
		return User{}, fmt.Errorf("user not found")
	}
	
	fmt.Println(result)
	return result, nil

}


func InsertUser(db *AlgoeDB.Database, user User) error {

	err := db.InsertOne(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


