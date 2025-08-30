package pgx

import (
	"github.com/jackc/pgx/v5"

	"postglide.io/postglide/internal/pg/connection"
)

type pgxRows struct {
	inner pgx.Rows
}

var _ connection.Rows = (*pgxRows)(nil)

func (r *pgxRows) Next() bool {
	return r.inner.Next()
}

func (r *pgxRows) Scan(dest ...any) error {
	return r.inner.Scan(dest...)
}

func (r *pgxRows) Close() error {
	r.inner.Close()
	return nil
}

func (r *pgxRows) Err() error {
	return r.inner.Err()
}

type pgxRow struct {
	inner pgx.Row
}

var _ connection.Row = (*pgxRow)(nil)

func (r *pgxRow) Scan(dest ...any) error {
	return r.inner.Scan(dest...)
}
