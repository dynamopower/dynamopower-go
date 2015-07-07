package connections

import (
	"errors"

	"github.com/dynamopower/dynamopower-go/constants"
)

type ConnectionPool struct {
	connections map[string]Connector
}

// Disconnect all connections
func (this *ConnectionPool) DisconnectAll() {
	for alias, _ := range this.connections {
		this.Deregister(alias)
	}
}

// Deregister a connection
func (this *ConnectionPool) Deregister(alias string) {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	if _, ok := this.connections[alias]; ok {
		delete(this.connections, alias)
	}
}

// Get a connection. nil is returned if the connection does not exist
func (this *ConnectionPool) Get(alias string) Connector {
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}

	return this.connections[alias]
}

// List all connection names
func (this *ConnectionPool) List() map[string]Connector {
	return this.connections
}

// Register DynamoDB configuration. Set region to "local" to connect
// to DynamoDB Local on port 8000
func (this *ConnectionPool) Register(alias, accessKey, secretKey, region string) (Connector, error) {
	var err error

	// Set default alias and region
	if alias == "" {
		alias = constants.DEFAULTCONNECTION
	}
	if region == "" {
		region = constants.DEFAULTREGION
	}

	// Ensure that the connection name is unique
	if _, exists := this.connections[alias]; exists {
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
	this.connections[alias] = &connection

	return &connection, err
}
