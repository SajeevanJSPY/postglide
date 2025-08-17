/*
Package connection is responsible for managing PostgreSQL backend connections
from the proxy layer. It provides a connection pool for efficient reuse of
PostgreSQL sessions, reducing the overhead of establishing new connections
for each query.

The package is typically used by the proxy to borrow and release backend
connections when routing client queries to PostgreSQL shards.

Features:
  - Connection pooling with limits on max open/idle connections.
  - Health checking and lazy reconnection on failure.
  - Connection lifecycle management (acquire, release, close).
  - Safe concurrent access.

This package does not implement PostgreSQL wire protocol directly — it
assumes the use of a protocol layer (e.g., pgproto3) to send/receive messages.
*/

package connection
