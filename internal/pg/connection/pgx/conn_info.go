package pgx

import (
	"errors"
	"fmt"
	"strings"
)

const (
	LoopBackAddr           = "127.0.0.1"
	DefaultPort     uint16 = 5432
	DefaultDatabase        = "postgres"

	SSLModeEnable  = "enable"
	SSLModeDisable = "disable"
)

type ConnInfo struct {
	host     string
	port     uint16
	user     string
	database string
	sslmode  string
}

type ConnOptFunc func(*ConnInfo)

func defaultPgConnInfo() ConnInfo {
	return ConnInfo{
		host:     LoopBackAddr,
		port:     DefaultPort,
		database: DefaultDatabase,
		sslmode:  SSLModeDisable,
	}
}

func WithHost(host string) ConnOptFunc {
	return func(pci *ConnInfo) {
		pci.host = host
	}
}

func WithPort(port uint16) ConnOptFunc {
	return func(pci *ConnInfo) {
		pci.port = port
	}
}

func WithUser(username string) ConnOptFunc {
	return func(pci *ConnInfo) {
		pci.user = username
	}
}

func WithDatabase(database string) ConnOptFunc {
	return func(pci *ConnInfo) {
		pci.database = database
	}
}

func WithSSLMode(isEnabled bool) ConnOptFunc {
	return func(ci *ConnInfo) {
		if isEnabled {
			ci.sslmode = SSLModeEnable
		} else {
			ci.sslmode = SSLModeDisable
		}
	}
}

func NewConnectionInfo(opts ...ConnOptFunc) ConnInfo {
	connInfo := defaultPgConnInfo()

	for _, fn := range opts {
		fn(&connInfo)
	}

	return connInfo
}

func (c *ConnInfo) Dsn(password string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.user, password, c.host, c.port, c.database, c.sslmode,
	)
}

func (c *ConnInfo) Validate() error {
	if strings.TrimSpace(c.host) == "" {
		return errors.New("hostname is empty")
	}

	if strings.TrimSpace(c.user) == "" {
		return errors.New("username is empty")
	}

	if strings.TrimSpace(c.database) == "" {
		return errors.New("database name is empty")
	}

	return nil
}
