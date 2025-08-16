package proxy_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"postglide.io/postglide/internal/proxy"
)

const (
	DefaultTcpPort = ":5432"
)

func TestProxyConn(t *testing.T) {
	assert := assert.New(t)

	t.Log("Running the pg proxy on default port")

	listener, err := net.Listen("tcp", DefaultTcpPort)
	assert.NoError(err)

	for {
		conn, err := listener.Accept()
		assert.NoError(err)

		proxy := proxy.NewProxy(conn)
		proxy.Run()
	}
}
