package connections

import (
	"github.com/aws/aws-sdk-go/aws/crendentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Connection struct {
	connection *dynamodb.DynamoDB
	accessKey  string
	secretKey  string
	region     string
}

// Connect to DynamoDB
func (c *Connection) Connect() {
	creds := crendentials.NewStaticCredentials(c.accessKey, c.secretKey, "")
	c.connection = dynamodb.New(creds, c.region, nil)
}
