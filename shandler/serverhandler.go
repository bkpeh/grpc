package serverhandler

import (
	"context"
	"errors"
	"fmt"

	. "github.com/bkpeh/grpc/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Server ...
type Server struct {
	UnimplementedGetPhoneNumberServer
}

//GetNum ...
func (s *Server) GetNum(ctx context.Context, in *Pid) (*Person, error) {
	fmt.Println("ID :", in.Id)

	x := &Person{
		Name:  "Wester",
		Id:    1020,
		Email: "wester@gmail.com",
		Phones: []*Person_PhoneNumber{
			{
				Number: "1234567",
				Type:   Person_HOME,
			},
		},
		LastUpdated: timestamppb.Now(),
	}

	if in.Id == 0 {
		return nil, errors.New("INVALID_ID")
	}

	if in.Id == x.Id {
		return x, nil
	}

	return nil, nil
}
