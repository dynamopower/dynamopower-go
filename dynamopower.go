package dynamopower_go

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dynamopower/dynamopower-go/connections"
	"github.com/dynamopower/dynamopower-go/model"
)

type DynamoPower struct {
	connectionPool *connections.ConnectionPool
}

// Constructor
func NewDynamoPower() *DynamoPower {
	return &DynamoPower{}
}

// Create table
func (this *DynamoPower) CreateTable(model model.Modeller) (*dynamodb.CreateTableOutput, error) {
}

// Build CreateTableInput
func (this *DynamoPower) buildCreateTableInput(model model.Modeller) *dynamodb.CreateTableInput {
	hashKey, hashKeyType, rangeKey, rangeKeyType := model.GetKeys()

	var attrDefs []*dynamodb.AttributeDefinition
	attrDefs = append(attrDefs, &dynamodb.AttributeDefinition{
		AttributeName: hashKey,
		AttributeType: hashKeyType,
	})

	if rangeKey != "" {
		attrDefs = append(attrDefs, &dynamodb.AttributeDefinition{
			AttributeName: rangeKey,
			AttributeType: rangeKeyType,
		})
	}

	return &dynamodb.CreateTableInput{
		AttributeDefinitions: attrDefs,
	}
}
