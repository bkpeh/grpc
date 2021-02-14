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

/*
func TestGetfromdb(t *testing.T) {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "us-west-2" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:8000",
				SigningRegion: "us-west-2",
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"), config.WithEndpointResolver(customResolver))
	svc := dynamodb.NewFromConfig(cfg)

	result := getfromDB(svc)

	if result != true {
		fmt.Println("Unexpected result.")
	}
}
*/
