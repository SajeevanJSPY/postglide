package meta

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/pgtype"

	"postglide.io/postglide/internal/pg/connection"
)

type VersionInfo struct {
	Version           pgtype.Text `json:"version"`
	VersionNumber     pgtype.Int8 `json:"verionNumber"`
	ActiveConnections pgtype.Int8 `json:"activeConnections"`
	MaxConnections    pgtype.Int8 `json:"maxConnections"`
}

const versionQuery = `
	SELECT
		version(),
		current_setting('server_version_num')::int8 AS version_number,
		(
		SELECT
		  count(*) AS active_connections
		FROM
		  pg_stat_activity
		) AS active_connections,
		current_setting('max_connections')::int8 AS max_connections
`

func GetVersionInfo(ctx context.Context, conn connection.Connection) (*VersionInfo, error) {
	var info VersionInfo
	err := conn.QueryRow(ctx, versionQuery).Scan(
		&info.Version,
		&info.VersionNumber,
		&info.ActiveConnections,
		&info.MaxConnections,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query server info: %w", err)
	}

	return &info, nil
}
