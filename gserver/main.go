package main

import (
	"fmt"
	"net"

	messages "github.com/bkpeh/grpc/proto"
	serverhandler "github.com/bkpeh/grpc/shandler"
	"google.golang.org/grpc"
)

func main() {

	lsvr, err := net.Listen("tcp", ":9005")

	if err != nil {
		fmt.Println("Fail to start listening.", err)
	}

	svr := serverhandler.Server{}

	//svr.SetBook(p)

	gsvr := grpc.NewServer()
	messages.RegisterGetPhoneNumberServer(gsvr, &svr)

	if err := gsvr.Serve(lsvr); err != nil {
		fmt.Println("Fail to start grpc.", err)
	}
}
