package pgx

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"postglide.io/postglide/internal/pg/connection"
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
		return errors.New("the connection already been established")
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
	if err == nil {
		return true
	}

	return false
}

func (c *pgxConnection) IsConnected() bool {
	return c.isConnected
}

func (c *pgxConnection) Close(ctx context.Context) error {
	c.pool.Close()

	return nil
}
