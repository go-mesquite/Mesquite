package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
)

/*
This file is just Jackson messing around with Go and what a database might look like.
I was looking at building an ORM or using GOORM initially. But The following are reasons why building my own database is advantageous:
- SQLite can only be used with Cgo
- Building something like litesail is easier if I don't use SQLite
- I can also better address the problem of saving uploaded files to the database. (Images, sound, ect.)
- We can have native type support. One example is that SQLite does not support arrays
- I care less about this one but it's fun to think that this approach might be faster than the convention
- Building a database in this way also gives us the freedom to implement things like ques. All of the state can be contained here for ease of use

I tried building this in binary first but I ran into problems with vars needing to be fixed length (So I had to use something like [50]byte instead of a string. I don't like it)
So I'm thinking that the best route will be storing data in json and other files in a folder.
This has the side benefit of data being recoverable. And uploading to s3 is like a regular backup

Also, it would be nice to build this in a way that it could be used with Python in the future. I could have used this with my previous Flask projects

Overview:
Define "tables" with structs (or Classes in Python's case)
Field constraints/details. unique, default values, many to many
Common dynamic types for fields like String, Integer, Float, list, Image/Bin (All mutable)
An interface for creating, reading, updating, and deleting records
*/

// What the database file might look like
// jack is used as a placeholder for whatever the db will be called. https://stackoverflow.com/questions/10858787/what-are-the-uses-for-struct-tags-in-go#30889373
// Use it similar to: https://gorm.io/docs/indexes.html#Index-Tag

type Blog struct {
	ID      int    `jack:"primary_key"`
	Author  Author `jack:"not_required"`
	Upvotes int32  `jack:"unique,nil,default:0"`
}

type Author struct {
	ID    int `jack:"primary_key"`
	Name  string
	Email string
}

// User represents a simple user struct.
type User struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

// JSONDB is a minimal in-memory database for storing structs in a JSON file.
type JSONDB struct {
	data     []User
	mu       sync.Mutex
	filename string
}

// NewJSONDB creates a new instance of the JSONDB.
func NewJSONDB(filename string) *JSONDB {
	return &JSONDB{
		data:     make([]User, 0),
		filename: filename,
	}
}

// Insert adds a struct to the in-memory database and writes the data to a JSON file.
func (db *JSONDB) Insert(user User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Add the user to the in-memory data
	db.data = append(db.data, user)

	// Write the data to the JSON file
	return db.writeToFile()
}

// GetAll retrieves all structs from the in-memory database.
func (db *JSONDB) GetAll() []User {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.data
}

// writeToFile writes the in-memory data to the JSON file.
func (db *JSONDB) writeToFile() error {
	jsonData, err := json.MarshalIndent(db.data, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(db.filename, jsonData, 0644)
}

// readFromFile reads the data from the JSON file into the in-memory data.
func (db *JSONDB) readFromFile() error {
	data, err := ioutil.ReadFile(db.filename)
	if err != nil {
		// Handle the case where the file does not exist or cannot be read
		return err
	}

	// Check if the file is empty
	if len(data) == 0 {
		// Initialize the data with an empty slice
		db.data = make([]User, 0)
		return nil
	}

	// Unmarshal JSON data into the in-memory data
	return json.Unmarshal(data, &db.data)
}

func main() {
	blog := Blog{
		ID:      1,
		Author:  Author{Name: "John Doe", Email: "john@example.com"},
		Upvotes: 42,
	}

	// Accessing struct tags using reflection
	blogType := reflect.TypeOf(blog)
	authorField, _ := blogType.FieldByName("Author")
	upvotesField, _ := blogType.FieldByName("Upvotes")

	fmt.Printf("Author struct tag: %s\n", authorField.Tag.Get("jack"))
	fmt.Printf("Upvotes struct tag: %s\n", upvotesField.Tag.Get("jack"))

	// Saving
	db_path := "db.jackalope/db.json"

	// Create a new JSONDB instance
	db := NewJSONDB(db_path)

	// Read data from the JSON file into the in-memory data
	err := db.readFromFile()
	if err != nil {
		fmt.Println("Error reading data from JSON file:", err)
		return
	}

	// Insert a user struct
	user := User{
		ID:   1, // Should things have an ID?
		Name: "Jackson Lohman",
		Age:  22,
	}
	err = db.Insert(user)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return
	}

	// Retrieve all users
	allUsers := db.GetAll()
	fmt.Println("All Users:")
	fmt.Println(allUsers)
}
