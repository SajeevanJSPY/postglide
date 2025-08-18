package connection

import (
	"context"
)

// Connection defines the interface for managing and interacting with a PostgreSQL connection
type Connection interface {
	// Establish opens a new connection to the database
	Establish(ctx context.Context, pass string) error

	// Ping verifies that the database is reachable and responsive
	Ping(context.Context) bool

	// IsConnected reports whether the connection has already been established
	IsConnected() bool

	// Close gracefully terminates the database connection
	Close(context.Context) error
}
