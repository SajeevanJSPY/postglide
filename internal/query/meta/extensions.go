package meta

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/pgtype"

	"postglide.io/postglide/internal/pg/connection"
)

type Extension struct {
	Name             pgtype.Text `json:"name"`
	Schema           pgtype.Text `json:"schema"`
	DefaultVersion   pgtype.Text `json:"defaultVersion"`
	InstalledVersion pgtype.Text `json:"installedVersion"`
	Comment          pgtype.Text `json:"comment"`
}

const extensionsQuery = `
	SELECT
		e.name,
		n.nspname as schema,
		e.default_version,
		x.extversion AS installed_version,
		e.comment
	FROM
		pg_available_extensions() e(name, default_version, comment)
		LEFT JOIN pg_extension x ON e.name = x.extname
		LEFT JOIN pg_namespace n ON x.extnamespace = n.oid
`

func GetExtensions(ctx context.Context, conn connection.Connection) ([]Extension, error) {
	rows, err := conn.Query(ctx, extensionsQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query server info: %w", err)
	}
	defer rows.Close()

	var extensions []Extension
	for rows.Next() {
		var extension Extension
		if err := rows.Scan(&extension.Name, &extension.Schema, &extension.DefaultVersion, &extension.InstalledVersion, &extension.Comment); err != nil {
			return nil, err
		}
		extensions = append(extensions, extension)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return extensions, nil
}
