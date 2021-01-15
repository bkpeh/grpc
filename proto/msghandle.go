package messages

import (
	"context"
	"fmt"
)

//Server ...
type Server struct {
	UnimplementedGetPhoneNumberServer
}

var ab AddressBook

//GetNum ...
func (s *Server) GetNum(ctx context.Context, in *Pid) (*Person, error) {
	fmt.Println("GetNum:", ab.GetPeople()[0].Name)

	return ab.GetPeople()[0], nil
}

//SetBook ...
func (s *Server) SetBook(a *AddressBook) {
	ab = *a
}
