package meta

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"postglide.io/postglide/internal/pg/connection"
)

type Schema struct {
	Id    pgtype.Int4 `json:"id"`
	Name  pgtype.Text `json:"name"`
	Owner pgtype.Text `json:"owner"`
}

const schemaQuery = `
	SELECT
		n.oid AS id,
		n.nspname AS name,
		u.rolname AS owner
	FROM
		pg_namespace n,
		pg_roles u
	WHERE
		n.nspowner = u.oid
		AND (
			pg_has_role(n.nspowner, 'USAGE')
			OR has_schema_privilege(n.oid, 'CREATE, USAGE')
		)
		AND NOT pg_catalog.starts_with(n.nspname, 'pg_temp_')
		AND NOT pg_catalog.starts_with(n.nspname, 'pg_toast_temp_')
`

func GetSchema(ctx context.Context, conn connection.Connection) ([]Schema, error) {
	rows, err := conn.Query(ctx, schemaQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query server info: %w", err)
	}
	defer rows.Close()

	var schemas []Schema
	for rows.Next() {
		var schema Schema
		if err := rows.Scan(&schema.Id, &schema.Name, &schema.Owner); err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return schemas, nil
}
