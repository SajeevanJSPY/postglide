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
	URI       string
}

func StartPgContainer(ctx context.Context) *PgContainer {
	pgContainer, err := postgres.Run(
		ctx,
		postgresContainerImage,
		postgres.WithDatabase(PgDbName),
		postgres.WithUsername(PgUserName),
		postgres.WithPassword(PgPassword),
	)
	if err != nil {
		panic(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}

	return &PgContainer{
		container: *pgContainer,
		URI:       connStr,
	}
}

func (pc *PgContainer) Close(ctx context.Context) {
	err := pc.container.Terminate(ctx)
	if err != nil {
		panic(err)
	}
}
