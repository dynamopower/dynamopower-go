package connections

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/dynamodb"
)

type Connection struct {
	connection *dynamodb.DynamoDB
	accessKey  string
	secretKey  string
	region     string
}

// Connect to DynamoDB
func (c *Connection) Connect() {
	creds := aws.Creds(c.accessKey, c.secretKey, "")
	c.connection = dynamodb.New(creds, c.region, nil)
}
