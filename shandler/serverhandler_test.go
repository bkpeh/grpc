package serverhandler

import (
	"context"
	"fmt"
	"testing"

	messages "github.com/bkpeh/grpc/proto"
)

type test struct {
	testname string
	pid      messages.Pid
}

func TestGetNum(t *testing.T) {
	testtable := []test{
		{testname: "Test1", pid: messages.Pid{Id: 1020}},
		{testname: "Test2", pid: messages.Pid{Id: 1000}},
		{testname: "Test3", pid: messages.Pid{Id: 0}},
	}

	svr := &Server{}

	for _, v := range testtable {
		t.Run(v.testname, func(t *testing.T) {
			fmt.Println("Running", v.testname)

			respond, err := svr.GetNum(context.Background(), &v.pid)

			switch v.testname {
			case "Test1":
				if respond.Name != "Wester" {
					fmt.Println("Expect Wester, Got ", respond.Name)
				}
			case "Test2":
				if respond != nil {
					fmt.Println("Expect NIL, Got ", respond.Name)
				}
			case "Test 3":
				if err == nil {
					fmt.Println("Expect Error:INVALID_ID, Got NIL")
				}
			}
		})
	}

}
