package proxy

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/jackc/pgx/v5/pgproto3"
)

type Proxy struct {
	backend     *pgproto3.Backend
	backendConn net.Conn

	msgch  chan pgproto3.FrontendMessage
	errch  chan error
	nextch chan struct{}
}

func NewProxy(backendConn net.Conn) *Proxy {
	backend := pgproto3.NewBackend(backendConn, backendConn)

	proxy := &Proxy{
		backend:     backend,
		backendConn: backendConn,
		msgch:       make(chan pgproto3.FrontendMessage),
		errch:       make(chan error, 1),
		nextch:      make(chan struct{}),
	}

	return proxy
}

func (p *Proxy) Run() error {
	defer p.Close()

	go p.ReadClientConn()

	for {
		select {
		case msg := <-p.msgch:
			buf, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			fmt.Println("F", string(buf))

			p.backend.Send(&pgproto3.AuthenticationOk{})
			p.backend.Flush()

			p.nextch <- struct{}{}
		case err := <-p.errch:
			return err
		}
	}
}

func (p *Proxy) Close() error {
	return p.backendConn.Close()
}

func (p *Proxy) ReadClientConn() {
	startupMessage, err := p.backend.ReceiveStartupMessage()
	if err != nil {
		p.errch <- err
		return
	}

	p.msgch <- startupMessage
	<-p.nextch

	for {
		msg, err := p.backend.Receive()
		if err != nil {
			p.errch <- err
			return
		}

		p.msgch <- msg
		<-p.nextch
	}
}
