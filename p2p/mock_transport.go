package p2p

import (
	"net"
	"time"

	na "github.com/cometbft/cometbft/p2p/netaddr"
)

type mockConnection struct {
	net.Conn
}

func (c *mockConnection) OpenStream(streamID byte) error {
	return nil
}
func (c *mockConnection) Read(streamID byte, b []byte) (n int, err error) {
	return 0, nil
}
func (c *mockConnection) Write(streamID byte, b []byte) (n int, err error) {
	return 0, nil
}
func (c *mockConnection) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}
func (c *mockConnection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *mockConnection) SetDeadline(t time.Time) error      { return c.Conn.SetReadDeadline(t) }
func (c *mockConnection) SetReadDeadline(t time.Time) error  { return c.Conn.SetReadDeadline(t) }
func (c *mockConnection) SetWriteDeadline(t time.Time) error { return c.Conn.SetWriteDeadline(t) }
func (c *mockConnection) Close(reason string) error          { return c.Conn.Close() }
func (c *mockConnection) FlushAndClose(reason string) error  { return c.Conn.Close() }
func (c *mockConnection) ConnectionState() any               { return nil }

var _ Transport = (*mockTransport)(nil)

type mockTransport struct {
	ln   net.Listener
	addr na.NetAddr
}

func (t *mockTransport) Listen(addr na.NetAddr) error {
	ln, err := net.Listen("tcp", addr.DialString())
	if err != nil {
		return err
	}
	t.addr = addr
	t.ln = ln
	return nil
}

func (t *mockTransport) NetAddr() na.NetAddr {
	return t.addr
}

func (t *mockTransport) Accept() (Connection, *na.NetAddr, error) {
	c, err := t.ln.Accept()
	return &mockConnection{Conn: c}, nil, err
}

func (*mockTransport) Dial(addr na.NetAddr) (Connection, error) {
	c, err := addr.DialTimeout(time.Second)
	return &mockConnection{Conn: c}, err
}

func (*mockTransport) Cleanup(Connection) error {
	return nil
}
