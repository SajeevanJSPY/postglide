package meta_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"postglide.io/postglide/internal/pg/connection"
	"postglide.io/postglide/internal/pg/connection/pgx"
	"postglide.io/postglide/internal/query/meta"
	"postglide.io/postglide/internal/testutils"
)

var pgContainer *testutils.PgContainer
var conn connection.Connection

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer = testutils.StartPgContainer(ctx)
	connInfo := pgx.NewConnectionInfo(
		pgx.WithHost(pgContainer.Host),
		pgx.WithPort(pgContainer.Port),
		pgx.WithUser(testutils.PgUserName),
		pgx.WithDatabase(testutils.PgDbName),
	)

	var err error
	conn, err = pgx.NewPgxConnection(connInfo)
	if err != nil {
		panic(err)
	}

	err = conn.Establish(context.Background(), testutils.PgPassword)
	if err != nil {
		panic(err)
	}

	m.Run()

	conn.Close(context.Background())
	pgContainer.Close(ctx)
}

func TestGetAllDatabase(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
	ping := conn.Ping(context.Background())
	assert.True(ping)
	dbs, err := meta.GetAllDatabases(context.Background(), conn)
	assert.NoError(err)
	assert.NotNil(dbs)
}

func TestCurrent(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
	ping := conn.Ping(context.Background())
	assert.True(ping)
	currentSession, err := meta.GetCurrentSession(context.Background(), conn)
	assert.NoError(err)
	assert.NotNil(currentSession)
}

func TestGetSchema(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
	ping := conn.Ping(context.Background())
	assert.True(ping)
	schemas, err := meta.GetSchema(context.Background(), conn)
	assert.NoError(err)
	assert.NotNil(schemas)
}

func TestGetExtensions(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
	ping := conn.Ping(context.Background())
	assert.True(ping)
	extensions, err := meta.GetExtensions(context.Background(), conn)
	assert.NoError(err)
	assert.NotNil(extensions)
}
