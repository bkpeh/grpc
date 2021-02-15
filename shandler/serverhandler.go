package serverhandler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

	result, err := getfromDB(svc, in)
	/*
		conf := aws.Config{
			Region:   aws.String("us-west-2"),
			Endpoint: aws.String("http://localhost:8000"),
		}

		sess := session.New(&conf)
		oldsvc := dynamodb.New(sess)

		result, err := getfromDB2(oldsvc, in)
	*/
	if err != nil {
		fmt.Println("Error.", err)
	}

	if in.Id == 0 {
		return nil, errors.New("INVALID_ID")
	}

	if in.Id == result.Id {
		return result, nil
	}

	return nil, nil
}

func getfromDB(svc *dynamodb.Client, pid *Pid) (*Person, error) {

	output, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("People"),
		Key:       map[string]types.AttributeValue{"Id": &types.AttributeValueMemberN{strconv.Itoa(int(pid.Id))}},
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
		return nil, errors.New("Error in MarshalMap" + err.Error())
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("People"),
		Item:      item,
	})

	if err != nil {
		fmt.Println("Error in PutItem", err)
		return nil, errors.New("Error in PutItem" + err.Error())
	}

	err = attributevalue.UnmarshalMap(output.Item, p)

	if err != nil {
		fmt.Println("Error in UnmarshalMap", err)
		return nil, errors.New("Error in UnmarshalMap" + err.Error())
	}

	fmt.Println("p:", p)

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5)})

	if err != nil {
		fmt.Println("Error in list table", err)
		return nil, errors.New("Error in list table" + err.Error())
	}

	for _, tableName := range resp.TableNames {
		fmt.Println("Table Name:", tableName)
	}

	return p, nil
}

/*
func getfromDB2(svc dynamodbiface.DynamoDBAPI, pid *Pid) (*Person, error) {

	output, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("People"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": &dynamodb.AttributeValue{N: aws.String("1020")},
		},
	})

	if output.Item == nil {
		fmt.Println("Error in GetItem", err)
	}

	p := &Person{}

	y := &Person{
		Name:  "Wester",
		Id:    1090,
		Email: "wester@gmail.com",
		Phones: []*Person_PhoneNumber{
			{
				Number: "1234567",
				Type:   Person_HOME,
			},
		},
		LastUpdated: timestamppb.Now(),
	}

	item, err := dynamodbattribute.MarshalMap(y)

	if err != nil {
		fmt.Println("Error in MarshalMap", err)
		return nil, errors.New("Error in MarshalMap" + err.Error())
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("People2"),
		Item:      item,
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Error in PutItem", err)
		return nil, errors.New("Error in PutItem" + err.Error())
	}
	/*
		err = dynamodbattribute.UnmarshalMap(output.Item, p)

		if err != nil {
			fmt.Println("Error in UnmarshalMap", err)
			return nil, errors.New("Error in UnmarshalMap" + err.Error())
		}

			newitem := item{}

			err = dynamodbattribute.UnmarshalMap(output.Item, &newitem)

			if err != nil {
				fmt.Println("Error in UnmarshalMap", err)
			}
*/
//fmt.Println("p:", p)
/*
		// Build the request with its input parameters
		resp, err := svc.ListTables(&dynamodb.ListTablesInput{})

		if err != nil {
			fmt.Println("Error in list table", err)
			return nil, errors.New("Error in list table" + err.Error())
		}

		for _, tableName := range resp.TableNames {
			fmt.Println("Table Name:", tableName)
		}

	return p, nil
}
*/
