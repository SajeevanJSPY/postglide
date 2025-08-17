package testutils

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	postgresContainerTag = "17.5"

	PgDbName   = "pg-test-db"
	PgUserName = "admin"
	PgPassword = "admin"
)

var (
	postgresContainerImage = fmt.Sprintf("postgres:%s", postgresContainerTag)
)

type PgContainer struct {
	container postgres.PostgresContainer
	Dsn       string
	Host      string
	Port      uint16
}

func StartPgContainer(ctx context.Context) *PgContainer {
	pgContainer, err := postgres.Run(
		ctx,
		postgresContainerImage,
		postgres.WithDatabase(PgDbName),
		postgres.WithUsername(PgUserName),
		postgres.WithPassword(PgPassword),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		panic(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}

	host, _ := pgContainer.Host(ctx)
	port, _ := pgContainer.MappedPort(ctx, "5432")

	return &PgContainer{
		container: *pgContainer,
		Dsn:       connStr,
		Host:      host,
		Port:      uint16(port.Int()),
	}
}

func (pc *PgContainer) Close(ctx context.Context) {
	err := pc.container.Terminate(ctx)
	if err != nil {
		panic(err)
	}
}
