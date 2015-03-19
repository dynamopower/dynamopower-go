package connections

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/dynamodb"
	"github.com/dynamopower/dynamopower-go/constants"
)

var connections = make(map[string]*dynamodb.DynamoDB)

// Connect to DynamoDB. Set region to "local" to connect
// to DynamoDB Local on port 8000
func connect(alias, accessKey, secretKey, region string) bool {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}
	if region == "" {
		region = constants.DEFAULTREGION
	}

	creds := aws.Creds(accessKey, secretKey, "")
	connections[alias] = dynamodb.New(creds, region, nil)

	return true
}

// Remove a connection from the list. Returns true if the connection was removed
func disconnect(alias string) bool {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	if _, ok := connections[alias]; ok {
		delete(connections, alias)
		return true
	}
	return false
}

// Get a connection. nil is returned if the connection does not exist
func get(alias string) *dynamodb.DynamoDB {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	if connection, ok := connections[alias]; ok {
		return connection
	}
	return nil
}

// List all connections
func list() []string {
	keys := make([]string, len(connections))

	i := 0
	for key := range connections {
		keys[i] = key
		i += 1
	}

	return keys
}
