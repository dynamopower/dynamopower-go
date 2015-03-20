package connections_test

import (
	"testing"

	"github.com/dynamopower/dynamopower-go/connections"
)

func TestConnect(t *testing.T) {
	if len(connections.List()) != 0 {
		t.Errorf("Unexpected number of connections: %n", len(connections.List()))
	}
	connections.Connect("", "abc", "123", "local")
	if len(connections.List()) != 1 {
		t.Errorf("Unexpected number of connections: %n", len(connections.List()))
	}
}
