package connection_test

import (
	"context"
	"testing"
	"time"

	"postglide.io/postglide/internal/connection"

	"github.com/stretchr/testify/assert"
)

const (
	META_IP_ADDR         = "localhost"
	META_USER            = "metaadmin"
	META_PASSWORD        = "admin"
	META_PORT            = 5432
	META_SCHEMA_DATABASE = "vschema"
)

func TestProxyConnection(t *testing.T) {
	pgConn := connection.New(META_IP_ADDR, META_PORT, META_USER)
	err := pgConn.Connect(META_SCHEMA_DATABASE, META_PASSWORD)
	if err != nil {
		panic(err)
	}
	defer pgConn.Close()

	var timestamp time.Time
	err = pgConn.PgxConnection().QueryRow(context.Background(), "SELECT current_timestamp;").Scan(&timestamp)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, pgConn.IsConnected(), true)
}
