package main

import (
	"google.golang.org/protobuf/proto"
	userpb "grpc-with-go/gen/sender/go/user/v1"
	"log"
	"os"
)

func main() {
	var addresses []*userpb.Address
	addresses = append(addresses, &userpb.Address{
		Street: "Armutlu",
		City:   "Istanbul",
	})
	// This is a comment
	user := userpb.User{
		Uuid:      "e0e3f1c0-1f1b-4b3b-8e1b-3b1f0e0e3f1c",
		FullName:  "John Doe",
		BirthYear: 1980,
		Salary:    nil,
		Addresses: addresses,
	}

	//Write user to the user.bin file
	data, err := proto.Marshal(&user)
	if err != nil {
		log.Fatal("Error marshalling user: ", err)
	}

	// Write data to the file
	err = os.WriteFile("user.bin", data, 0644)
	if err != nil {
		log.Fatal("Error writing to file: ", err)
	}

}
