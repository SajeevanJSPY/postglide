package meta

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"postglide.io/postglide/internal/pg/connection"
)

type Database struct {
	Name pgtype.Text
}

const databasesQuery = `
	SELECT datname
	FROM pg_database
	WHERE datistemplate = FALSE
`

func GetAllDatabases(ctx context.Context, conn connection.Connection) ([]Database, error) {
	rows, err := conn.Query(ctx, databasesQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query server info: %w", err)
	}
	defer rows.Close()

	var databases []Database
	for rows.Next() {
		var db Database
		if err := rows.Scan(&db.Name); err != nil {
			return nil, err
		}
		databases = append(databases, db)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return databases, nil
}
