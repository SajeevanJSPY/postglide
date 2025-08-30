package meta

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"postglide.io/postglide/internal/pg/connection"
)

type Current struct {
	Schema   pgtype.Text `json:"schema"`
	Database pgtype.Text `json:"database"`
	User     pgtype.Text `json:"user"`
	Role     pgtype.Text `json:"role"`
}

const currentQuery string = `
	SELECT
	current_schema() AS schema,
	current_database() AS database,
	current_user AS user,
	current_role AS role`

func GetCurrentSession(ctx context.Context, conn connection.Connection) (*Current, error) {
	var current Current

	err := conn.QueryRow(ctx, currentQuery).Scan(
		&current.Schema,
		&current.Database,
		&current.User,
		&current.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query server info: %w", err)
	}

	return &current, nil
}
