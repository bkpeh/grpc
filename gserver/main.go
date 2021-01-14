package main

import (
	"fmt"
	"net"

	messages "github.com/bkpeh/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	p := messages.Person{
		Name:  "Wester",
		Id:    1020,
		Email: "wester@gmail.com",
		Phones: []*messages.Person_PhoneNumber{
			{
				Number: "1234567",
				Type:   messages.Person_HOME,
			},
		},
	}

	lsvr, err := net.Listen("tcp", ":9005")

	if err != nil {
		fmt.Println("Fail to start listening.", err)
	}

	gsvr := grpc.NewServer()

	if err := gsvr.Serve(lsvr); err != nil {
		fmt.Println("Fail to start grpc.", err)
	}

	fmt.Println(p)
}
