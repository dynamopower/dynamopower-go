package connections

import (
	"errors"

	"github.com/dynamopower/dynamopower-go/constants"
)

var connections = make(map[string]Connector)

// Disconnect all connections
func DisconnectAll() {
	for alias, _ := range connections {
		Deregister(alias)
	}
}

// Deregister a connection
func Deregister(alias string) {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	if _, ok := connections[alias]; ok {
		delete(connections, alias)
	}
}

// Get a connection. nil is returned if the connection does not exist
func Get(alias string) (connection Connector) {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	connection = connections[alias]
	return
}

// List all connection names
func List() map[string]Connector {
	return connections
}

// Register DynamoDB configuration. Set region to "local" to connect
// to DynamoDB Local on port 8000
func Register(alias, accessKey, secretKey, region string) (Connector, error) {
	var err error

	// Set default alias and region
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}
	if region == "" {
		region = constants.DEFAULTREGION
	}

	// Ensure that the connection name is unique
	if _, exists := connections[alias]; exists {
		err = errors.New("Connection already exists")
		return new(Connection), err
	}

	// Create the connection object
	connection := Connection{
		accessKey: accessKey,
		secretKey: secretKey,
		region:    region}

	// Connect to DynamoDB
	connection.Connect()

	// Add the connection to the list
	connections[alias] = &connection

	return &connection, err
}
