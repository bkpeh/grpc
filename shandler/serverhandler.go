package serverhandler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

	getfromDB()

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

func getfromDB() {

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

	//a := types.AttributeValueMemberS{Value: `"Email":"easter@gmail.com"`}
	//a := types.AttributeValueMemberS{S: "easter@gmail.com"}
	a := types.AttributeValueMemberN{"1020"}
	output, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("People"),
		Key:       map[string]types.AttributeValueMemberN{&a},
	})

	fmt.Println("output:", output)

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5)})

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}
