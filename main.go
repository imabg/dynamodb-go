package main

import (
	localConfig "github.com/imabg/dynamodb-go/config"
	dynamodb "github.com/imabg/dynamodb-go/dynamodb"	
	"github.com/imabg/dynamodb-go/logger"
	"github.com/imabg/dynamodb-go/server"
	"github.com/spf13/viper"
)

var env localConfig.Config

func init() {
	// setup config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&env)	
	if err != nil {
		panic(err)
	}
}


func main() {
	// setup-server
	logger := logger.NewLogger()
	srv := server.NewServer(env, logger)
	srv.Start()
	defer srv.Stop()
	// setup dynamodb
	 _ = dynamodb.NewDynamoDBClient(env)
}
