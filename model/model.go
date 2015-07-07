package model

type Modeller interface {
	// Returns the name of the DynamoDB table
	GetTableName() string

	// Returns the following
	// - Hash key
	// - Hash key type
	// - Range key
	// - Range key type
	GetKeys() (string, string, string, string)
}