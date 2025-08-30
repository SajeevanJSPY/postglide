package pgx

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"postglide.io/postglide/internal/pg/connection"
)

var (
	ErrConnectionNotEstablished     = errors.New("connection is not established")
	ErrConnectionAlreadyEstablished = errors.New("connection is already established")
)

type pgxConnection struct {
	ConnInfo

	pool        *pgxpool.Pool
	isConnected bool
}

func NewPgxConnection(connInfo ConnInfo) (*pgxConnection, error) {
	// TODO: Improve validation errors:
	//   - Return clear, user-friendly messages
	//   - Decide if we should auto-retry/reconnect or let caller handle
	if err := connInfo.Validate(); err != nil {
		return nil, err
	}

	return &pgxConnection{
		ConnInfo:    connInfo,
		isConnected: false,
	}, nil
}

var _ connection.Connection = (*pgxConnection)(nil)

func (c *pgxConnection) Establish(ctx context.Context, pass string) error {
	if c.IsConnected() || c.pool != nil {
		return ErrConnectionAlreadyEstablished
	}

	pool, err := pgxpool.New(ctx, c.Dsn(pass))
	if err != nil {
		return err
	}

	c.pool = pool
	c.isConnected = true

	return nil
}

func (c *pgxConnection) Ping(ctx context.Context) bool {
	if !c.IsConnected() {
		return false
	}

	err := c.pool.Ping(ctx)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (c *pgxConnection) IsConnected() bool {
	return c.isConnected
}

func (c *pgxConnection) Query(ctx context.Context, sql string, args ...any) (connection.Rows, error) {
	if !c.IsConnected() {
		return nil, ErrConnectionNotEstablished
	}

	r, err := c.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &pgxRows{inner: r}, nil
}

func (c *pgxConnection) QueryRow(ctx context.Context, sql string, args ...any) connection.Row {
	if !c.IsConnected() {
		return nil
	}

	r := c.pool.QueryRow(ctx, sql, args...)
	return &pgxRow{inner: r}
}

func (c *pgxConnection) Close(ctx context.Context) error {
	c.pool.Close()

	return nil
}
