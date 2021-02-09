package main

import (
	"context"
	"fmt"

	messages "github.com/bkpeh/grpc/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9005", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	client := messages.NewGetPhoneNumberClient(conn)

	p := messages.Pid{
		Id: 1020,
	}

	respond, err := client.GetNum(context.Background(), &p)

	if err != nil {
		fmt.Println("Respond error.", err)
	}

	fmt.Println("Respond:", respond.Name)
}
