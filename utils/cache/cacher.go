package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "johndoe@example.com",
	}
	//id := string(user.ID)
	//data, err := user.MarshalBinary()
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Set(context.Background(), "john_doe", data, time.Hour).Err()

	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Del(context.Background(), "john_doe").Result()
	if err != nil {
		log.Fatal(err)
	}
	exist, err := client.Exists(context.Background(), "john_doe").Result()
	if err != nil {
		log.Fatal(err)
	}
	if exist == 0 {
		fmt.Println("Data sudah terhapus")
		return
	}
	val, err := client.Get(context.Background(), "john_doe").Result()
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal([]byte(val), &m)
	if err != nil {
		// Handle the error
	}

	fmt.Println("ID:")
	fmt.Println("Username:", m["username"])
	fmt.Println("Email:", m["email"])

	//menggunakan struct
	//var m User
	//err = json.Unmarshal([]byte(val), &m)
	//if err != nil {
	//	// Handle the error
	//}
	//
	//fmt.Println("ID:", m.ID)
	//fmt.Println("Username:", m.Username)
	//fmt.Println("Email:", m.Email)
}
