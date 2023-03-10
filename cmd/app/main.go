package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chakravarthy7102/serverless/config"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
	"github.com/Chakravarthy7102/serverless/internal/repository/instance"
	"github.com/Chakravarthy7102/serverless/internal/routes"
	"github.com/Chakravarthy7102/serverless/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)
	logger.INFO("Waiting for the service to start...", nil)

	errors := Migrate(connection)

	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on Migration", err)
		}
	}

	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)

	router := routes.NewRouter().SetRouters(repository)

	logger.INFO("service is running on port :%v", port)

	server := http.ListenAndServe(port, router)

	log.Fatal(server)

}

func Migrate(connection *dynamodb.DynamoDB) []error {

	var errors []error

	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})

	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)

	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})

	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables not found!", nil)
		}
	}

	for _, tableName := range response.TableNames {
		logger.INFO("Table found", *tableName)
	}

	return err
}
