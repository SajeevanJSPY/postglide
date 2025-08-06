package connection

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PgConnection struct {
	IpAddr      string
	Port        uint16
	User        string
	isConnected bool
	conn        *pgx.Conn
}

func New(ipaddr string, port uint16, user string) *PgConnection {
	return &PgConnection{
		IpAddr:      ipaddr,
		Port:        port,
		User:        user,
		isConnected: false,
	}
}

func (p *PgConnection) DbUrl(dbname string, secret string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", p.User, secret, p.IpAddr, p.Port, dbname)
}

func (p *PgConnection) IsConnected() bool {
	return p.isConnected
}

func (p *PgConnection) PgxConnection() *pgx.Conn {
	return p.conn
}

func (p *PgConnection) Connect(dbname string, secret string) error {
	if p.IsConnected() {
		return errors.New("the connection already been established")
	}

	conn, err := pgx.Connect(context.Background(), p.DbUrl(dbname, secret))
	if err != nil {
		return err
	}

	p.isConnected = true
	p.conn = conn

	return nil
}

func (p *PgConnection) Close() {
	p.conn.Close(context.Background())
}
