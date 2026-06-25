package dynamodb

import (
	"context"
	localConfig "github.com/imabg/dynamodb-go/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)


func setupAWSConfig(env localConfig.Config) aws.Config {
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

func NewDynamoDBClient(env localConfig.Config) *dynamodb.Client {
	cfg := setupAWSConfig(env)
	return dynamodb.NewFromConfig(cfg)
}