package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	userpb "grpc-with-go/gen/sender/go/user/v1"
	"log"
	"os"
)

func main() {

	// Read user from the user.bin file
	data, err := os.ReadFile("user.bin")
	if err != nil {
		log.Fatal("Error reading from file: ", err)
	}

	// Decode the data
	var user userpb.User
	err = proto.Unmarshal(data, &user)
	if err != nil {
		log.Fatal("Error unmarshalling user: ", err)
	}

	fmt.Println("User UUID: ", user.Uuid)
	fmt.Println("User Full Name: ", user.FullName)
	fmt.Println("User Birth Year: ", user.BirthYear)
	fmt.Println("User Salary: ", user.Salary)
	fmt.Println("User Addresses: ", user.Addresses)

}
