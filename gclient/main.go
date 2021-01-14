package main

import (
	"fmt"

	messages "github.com/bkpeh/grpc/proto"
)

func main() {
	id := messages.Pid{
		Id: 1020,
	}

	fmt.Println(id)
}
