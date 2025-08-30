package connection

import (
	"context"
)

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}

// Connection defines the interface for managing and interacting with a PostgreSQL connection
type Connection interface {
	// Establish opens a new connection to the database
	Establish(ctx context.Context, pass string) error

	// Ping verifies that the database is reachable and responsive
	Ping(context.Context) bool

	// IsConnected reports whether the connection has already been established
	IsConnected() bool

	QueryRow(ctx context.Context, sql string, args ...any) Row

	Query(ctx context.Context, sql string, args ...any) (Rows, error)

	// Close gracefully terminates the database connection
	Close(context.Context) error
}
