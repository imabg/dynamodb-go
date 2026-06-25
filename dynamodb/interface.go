package dynamodb

type IDynamoDB interface {
	CreateTable() ( error)
}