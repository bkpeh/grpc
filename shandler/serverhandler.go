package serverhandler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

	getfromDB(svc)

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

func getfromDB(svc *dynamodb.Client) bool {

	output, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("People"),
		Key:       map[string]types.AttributeValue{"Id": &types.AttributeValueMemberN{"1020"}},
	})

	p := &Person{}

	y := &Person{
		Name:  "Wester",
		Id:    1080,
		Email: "wester@gmail.com",
		Phones: []*Person_PhoneNumber{
			{
				Number: "1234567",
				Type:   Person_HOME,
			},
		},
		LastUpdated: timestamppb.Now(),
	}

	item, err := attributevalue.MarshalMap(y)

	if err != nil {
		fmt.Println("Error in MarshalMap", err)
		return false
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("People"),
		Item:      item,
	})

	if err != nil {
		fmt.Println("Error in PutItem", err)
		return false
	}

	err = attributevalue.UnmarshalMap(output.Item, p)

	if err != nil {
		fmt.Println("Error in UnmarshalMap", err)
		return false
	}

	fmt.Println("p:", p)

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5)})

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
		return false
	}

	for _, tableName := range resp.TableNames {
		fmt.Println("Table Name:", tableName)
	}

	return true
}
