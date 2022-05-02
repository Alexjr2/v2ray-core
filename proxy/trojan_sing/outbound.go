package trojan_sing

import (
	"context"
	"crypto/tls"
	"io"
	"time"

	C "github.com/sagernet/sing/common"
	B "github.com/sagernet/sing/common/buf"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/rw"
	"github.com/sagernet/sing/protocol/socks"
	"github.com/sagernet/sing/protocol/trojan"
	"github.com/v2fly/v2ray-core/v5/common"
	"github.com/v2fly/v2ray-core/v5/common/buf"
	"github.com/v2fly/v2ray-core/v5/common/net"
	"github.com/v2fly/v2ray-core/v5/common/retry"
	"github.com/v2fly/v2ray-core/v5/common/session"
	"github.com/v2fly/v2ray-core/v5/proxy/shadowsocks_sing"
	"github.com/v2fly/v2ray-core/v5/transport"
	"github.com/v2fly/v2ray-core/v5/transport/internet"
	"github.com/v2fly/v2ray-core/v5/transport/pipe"
)

func init() {
	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewClient(ctx, config.(*ClientConfig))
	}))
}

type Client struct {
	ctx        context.Context
	server     net.Destination
	key        [trojan.KeyLength]byte
	serverName string
	nextProtos []string
	insecure   bool
}

func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	c := &Client{
		ctx: ctx,
		server: net.Destination{
			Address: config.Address.AsAddress(),
			Port:    net.Port(config.Port),
			Network: net.Network_TCP,
		},
		key:        trojan.Key(config.Password),
		serverName: config.ServerName,
		nextProtos: config.NextProtos,
		insecure:   config.Insecure,
	}
	if c.serverName == "" {
		c.serverName = c.server.Address.String()
	}
	return c, nil
}

func (c *Client) ProcessConn(ctx context.Context, conn net.Conn, dialer internet.Dialer) error {
	outbound := session.OutboundFromContext(ctx)
	if outbound == nil || !outbound.Target.IsValid() {
		return newError("target not specified")
	}
	destination := outbound.Target
	network := destination.Network

	var outboundConn internet.Connection
	err := retry.ExponentialBackoff(5, 100).On(func() error {
		rawConn, err := dialer.Dial(ctx, c.server)
		if err != nil {
			return err
		}

		outboundConn = rawConn
		return nil
	})
	if err != nil {
		return newError("failed to find an available destination").AtWarning().Base(err)
	}
	newError("tunneling request to ", destination, " via ", c.server).WriteToLog(session.ExportIDToError(ctx))

	connElem := net.AddConnection(outboundConn)
	defer net.RemoveConnection(connElem)

	tlsConfig := &tls.Config{
		ServerName:         c.serverName,
		InsecureSkipVerify: c.insecure,
	}

	if len(c.nextProtos) > 0 {
		tlsConfig.NextProtos = c.nextProtos
	}

	tlsConn := tls.Client(outboundConn, tlsConfig)

	if network == net.Network_TCP {
		serverConn := trojan.NewClientConn(tlsConn, c.key, shadowsocks_sing.SingDestination(destination))

		if cr, ok := conn.(rw.CachedReader); ok {
			cached := cr.ReadCached()
			if cached != nil && !cached.IsEmpty() {
				_, err = serverConn.Write(cached.Bytes())
				cached.Release()
				if err != nil {
					return newError("client handshake").Base(err)
				}
				goto direct
			}
		}

		{
			err = conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			if err != nil {
				return err
			}

			_request := B.StackNew()
			request := C.Dup(_request)

			_, err = request.ReadFrom(conn)
			if err != nil && !E.IsTimeout(err) {
				return err
			}

			err = conn.SetReadDeadline(time.Time{})
			if err != nil {
				return err
			}

			_, err = serverConn.Write(request.Bytes())
			if err != nil {
				return newError("client handshake").Base(err)
			}
		}

	direct:
		return rw.CopyConn(ctx, conn, serverConn)
	} else {
		var packetConn socks.PacketConn
		if sc, isPacketConn := conn.(socks.PacketConn); isPacketConn {
			packetConn = sc
		} else if nc, isNetPacket := conn.(net.PacketConn); isNetPacket {
			packetConn = &socks.PacketConnWrapper{PacketConn: nc}
		} else {
			packetConn = &shadowsocks_sing.PacketConnWrapper{
				Reader: buf.NewReader(conn),
				Writer: buf.NewWriter(conn),
				Conn:   conn,
				Dest:   destination,
			}
		}
		return socks.CopyPacketConn(ctx, trojan.NewClientPacketConn(outboundConn, c.key), packetConn)
	}
}

func (c *Client) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	var inboundConn net.Conn
	inbound := session.InboundFromContext(ctx)
	if inbound != nil {
		inboundConn = inbound.Conn
	}

	outbound := session.OutboundFromContext(ctx)
	if outbound == nil || !outbound.Target.IsValid() {
		return newError("target not specified")
	}
	destination := outbound.Target
	network := destination.Network

	var outboundConn internet.Connection
	err := retry.ExponentialBackoff(5, 100).On(func() error {
		rawConn, err := dialer.Dial(ctx, c.server)
		if err != nil {
			return err
		}

		outboundConn = rawConn
		return nil
	})
	if err != nil {
		return newError("failed to find an available destination").AtWarning().Base(err)
	}
	newError("tunneling request to ", destination, " via ", c.server).WriteToLog(session.ExportIDToError(ctx))

	connElem := net.AddConnection(outboundConn)
	defer net.RemoveConnection(connElem)

	tlsConfig := &tls.Config{
		ServerName:         c.serverName,
		InsecureSkipVerify: c.insecure,
	}

	if len(c.nextProtos) > 0 {
		tlsConfig.NextProtos = c.nextProtos
	}

	tlsConn := tls.Client(outboundConn, tlsConfig)

	if network == net.Network_TCP {
		serverConn := trojan.NewClientConn(tlsConn, c.key, shadowsocks_sing.SingDestination(destination))

		var handshake bool
		if cachedReader, isCached := link.Reader.(pipe.CachedReader); isCached {
			cached, _ := cachedReader.ReadMultiBufferCached()
			if cached != nil && !cached.IsEmpty() {
				_payload := B.StackNew()
				payload := C.Dup(_payload)
				for {
					payload.FullReset()
					nb, n := buf.SplitBytes(cached, payload.FreeBytes())
					if n > 0 {
						payload.Truncate(n)
						_, err = serverConn.Write(payload.Bytes())
						if err != nil {
							return newError("write payload").Base(err)
						}
						handshake = true
					}
					if nb.IsEmpty() {
						break
					} else {
						cached = nb
					}
				}
			}
		}
		if !handshake {
			if timeoutReader, isTimeoutReader := link.Reader.(buf.TimeoutReader); isTimeoutReader {
				mb, err := timeoutReader.ReadMultiBufferTimeout(time.Millisecond * 100)
				if err != nil && err != buf.ErrNotTimeoutReader && err != buf.ErrReadTimeout {
					return newError("read payload").Base(err)
				}
				_payload := B.StackNew()
				payload := C.Dup(_payload)
				for {
					payload.FullReset()
					nb, n := buf.SplitBytes(mb, payload.FreeBytes())
					if n > 0 {
						payload.Truncate(n)
						_, err = serverConn.Write(payload.Bytes())
						if err != nil {
							return newError("write payload").Base(err)
						}
						handshake = true
					}
					if nb.IsEmpty() {
						break
					} else {
						mb = nb
					}
				}
			}
		}
		if !handshake {
			_, err = serverConn.Write(nil)
			if err != nil {
				return newError("client handshake").Base(err)
			}
		}

		pipeIn := pipe.IsPipe(link.Reader)
		pipeOut := pipe.IsPipe(link.Writer)

		if inboundConn != nil && !pipeIn && !pipeOut {
			return rw.CopyConn(ctx, inboundConn, serverConn)
		}

		conn := &shadowsocks_sing.PipeConnWrapper{
			W:       link.Writer,
			PipeOut: pipeOut,
			Conn:    inboundConn,
		}
		if ir, ok := link.Reader.(io.Reader); ok {
			conn.R = ir
		} else {
			conn.R = &buf.BufferedReader{Reader: link.Reader}
		}

		return rw.CopyConn(ctx, conn, serverConn)
	} else {
		var packetConn socks.PacketConn
		if pc, isPacketConn := inboundConn.(socks.PacketConn); isPacketConn {
			packetConn = pc
		} else {
			packetConn = &shadowsocks_sing.PacketConnWrapper{
				Reader: link.Reader,
				Writer: link.Writer,
				Conn:   inboundConn,
				Dest:   destination,
			}
		}

		return socks.CopyPacketConn(ctx, packetConn, trojan.NewClientPacketConn(outboundConn, c.key))
	}
}
