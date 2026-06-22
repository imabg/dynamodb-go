package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	localConfig "github.com/imabg/dynamodb-go/config"
	"github.com/spf13/viper"
)

type DynamoDB struct {
	ctx 		context.Context
	dbClient *dynamodb.Client
	tableName string
}

func (d *DynamoDB) CreateTable() (*types.TableDescription, error) {
	table, err := d.dbClient.CreateTable(d.ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(d.tableName),
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("year"),
			AttributeType: types.ScalarAttributeTypeN,
		}, {
			AttributeName: aws.String("title"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("year"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("title"),
			KeyType:       types.KeyTypeRange,
		}},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return nil, err
	}
	conn := dynamodb.NewTableExistsWaiter(d.dbClient)
	err = conn.Wait(d.ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(d.tableName),
	}, 0)	
	tableDesc := table.TableDescription
	return tableDesc, nil
}

var err error
var env localConfig.Config

func init() {
	// setup config
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&env)	
	if err != nil {
		panic(err)
	}
}

func setupAWSConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		env.AWS_ACCESS_KEY_ID,
		env.AWS_SECRET_ACCESS_KEY,
		"",
	)), config.WithRegion(env.AWS_REGION))
	if err != nil {
		panic(err)
	}
	return cfg
}

func main() {
	//  Create table
	cfg := setupAWSConfig()
	db := &DynamoDB{
		ctx: context.Background(),
		dbClient: dynamodb.NewFromConfig(cfg),
		tableName: "Movies",
	}
	tableDesc, err := db.CreateTable()
	if err != nil {
		panic(err)
	}
	fmt.Println(tableDesc.TableId)
}
