package connections_test

import (
	"testing"

	"github.com/dynamopower/dynamopower-go/connections"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConnectionsTestSuite struct {
	suite.Suite
}

// Prepare tests
func (suite *ConnectionsTestSuite) SetupTest() {
	connection, _ := connections.Register("", "access", "secret", "local")
	connection.Connect()
}

// Clean up after tests
func (suite *ConnectionsTestSuite) TearDownTest() {
	connections.DisconnectAll()
}

func (suite *ConnectionsTestSuite) TestRegister() {
	assert.Equal(suite.T(), len(connections.List()), 1)

	// Register a new connection
	_, err := connections.Register("test", "abc", "123", "local")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(connections.List()), 2)

	// It should not be possible to register a new connection
	// with the same alias
	_, err = connections.Register("test", "abc", "123", "local")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), len(connections.List()), 2)
}

// Make the test suite available to go test
func TestConnectionsTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionsTestSuite))
}
