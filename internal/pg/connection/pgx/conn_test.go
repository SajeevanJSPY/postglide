package pgx_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"postglide.io/postglide/internal/pg/connection/pgx"
	"postglide.io/postglide/internal/testutils"
)

var pgContainer *testutils.PgContainer

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer = testutils.StartPgContainer(ctx)

	m.Run()

	pgContainer.Close(ctx)
}

func TestPgxConnection(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()

	// connection context
	connCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// query context with timeout
	queryCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	connInfo := pgx.NewConnectionInfo(
		pgx.WithHost(pgContainer.Host),
		pgx.WithPort(pgContainer.Port),
		pgx.WithUser(testutils.PgUserName),
		pgx.WithDatabase(testutils.PgDbName),
	)

	pgxConn, err := pgx.NewPgxConnection(connInfo)
	assert.NoError(err)
	assert.NotNil(pgxConn)
	defer func() {
		assert.NoError(pgxConn.Close(connCtx))
	}()

	err = pgxConn.Establish(connCtx, testutils.PgPassword)
	assert.NoError(err)

	assert.True(pgxConn.Ping(queryCtx))
}

func TestConnInfo(t *testing.T) {
	assert := assert.New(t)

	t.Run("only with default values", func(t *testing.T) {
		connInfo := pgx.NewConnectionInfo()
		assert.Error(connInfo.Validate(), "ConnInfo without user/database should not validate")
	})

	t.Run("matching the docker Dsn", func(t *testing.T) {
		connInfo := pgx.NewConnectionInfo(
			pgx.WithHost(pgContainer.Host),
			pgx.WithPort(pgContainer.Port),
			pgx.WithUser(testutils.PgUserName),
			pgx.WithDatabase(testutils.PgDbName),
		)
		assert.NoError(connInfo.Validate(), "ConnInfo with all fields should validate")

		got := connInfo.Dsn(testutils.PgPassword)
		assert.Equal(pgContainer.Dsn, got, "ConnInfo DSN should match container DSN")
	})
}
