/*
Package metadata provides utilities to retrieve structural and runtime
information from a PostgreSQL database using:

  - information_schema views
  - PostgreSQL system views (e.g., pg_stat_*)
  - PostgreSQL system catalogs (e.g., pg_class, pg_attribute)

This package is useful for dynamically discovering or introspecting the
underlying database — such as during query planning, routing decisions,
schema syncing, or health checking.
*/

package pgmeta
