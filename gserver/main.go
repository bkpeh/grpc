package main

import (
	"fmt"
	"net"

	messages "github.com/bkpeh/grpc/proto"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	x := []*messages.Person{
		{
			Name:  "Wester",
			Id:    1020,
			Email: "wester@gmail.com",
			Phones: []*messages.Person_PhoneNumber{
				{
					Number: "1234567",
					Type:   messages.Person_HOME,
				},
			},
			LastUpdated: timestamppb.Now(),
		},
	}
	fmt.Println(x)
	p := messages.AddressBook{
		People: []*messages.Person{
			{
				Name:  "Wester",
				Id:    1020,
				Email: "wester@gmail.com",
				Phones: []*messages.Person_PhoneNumber{
					{
						Number: "1234567",
						Type:   messages.Person_HOME,
					},
				},
				LastUpdated: timestamppb.Now(),
			},
		},
	}
	fmt.Println(p)
	lsvr, err := net.Listen("tcp", ":9005")

	if err != nil {
		fmt.Println("Fail to start listening.", err)
	}

	svr := messages.Server{}

	svr.SetBook(&p)

	gsvr := grpc.NewServer()
	messages.RegisterGetPhoneNumberServer(gsvr, &svr)

	if err := gsvr.Serve(lsvr); err != nil {
		fmt.Println("Fail to start grpc.", err)
	}
}
