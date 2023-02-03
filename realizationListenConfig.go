package kcp

import (
	"net"
	"syscall"
	"time"
)

type ListenConfig struct {
	// If Control is not nil, it is called after creating the network
	// connection but before binding it to the operating system.
	//
	// Network and address parameters passed to Control method are not
	// necessarily the ones passed to Listen. For example, passing "tcp" to
	// Listen will cause the Control function to be called with "tcp4" or "tcp6".
	Control func(network, address string, c syscall.RawConn) error

	// KeepAlive specifies the keep-alive period for network
	// connections accepted by this listener.
	// If zero, keep-alives are enabled if supported by the protocol
	// and operating system. Network protocols or operating systems
	// that do not support keep-alives ignore this field.
	// If negative, keep-alives are disabled.
	KeepAlive time.Duration
}


func (lc *ListenConfig) Listen(network, address string) (net.Listener, error) {
	listener, err := Listen(address)
	if err != nil{
		return nil, &net.OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
	}
	return listener, nil
}

func (lc *ListenConfig) ListenPacket(address string) (net.PacketConn, error) {
	panic("Реализуй \"Listen Packet\"")
	// listener, err := Listen(address)
	// if err != nil{
	// 	return nil, &net.OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
	// }
	// return listener, nil
}

// func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error) {
// 	addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
// 	if err != nil {
// 		return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
// 	}
// 	sl := &sysListener{
// 		ListenConfig: *lc,
// 		network:      network,
// 		address:      address,
// 	}
// 	var l Listener
// 	la := addrs.first(isIPv4)
// 	switch la := la.(type) {
// 	case *TCPAddr:
// 		l, err = sl.listenTCP(ctx, la)
// 	case *UnixAddr:
// 		l, err = sl.listenUnix(ctx, la)
// 	default:
// 		return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
// 	}
// 	if err != nil {
// 		return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err} // l is non-nil interface containing nil pointer
// 	}
// 	return l, nil
// }

// // ListenPacket announces on the local network address.
// //
// // See func ListenPacket for a description of the network and address
// // parameters.
// func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error) {
// 	addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
// 	if err != nil {
// 		return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
// 	}
// 	sl := &sysListener{
// 		ListenConfig: *lc,
// 		network:      network,
// 		address:      address,
// 	}
// 	var c PacketConn
// 	la := addrs.first(isIPv4)
// 	switch la := la.(type) {
// 	case *UDPAddr:
// 		c, err = sl.listenUDP(ctx, la)
// 	case *IPAddr:
// 		c, err = sl.listenIP(ctx, la)
// 	case *UnixAddr:
// 		c, err = sl.listenUnixgram(ctx, la)
// 	default:
// 		return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: address}}
// 	}
// 	if err != nil {
// 		return nil, &OpError{Op: "listen", Net: sl.network, Source: nil, Addr: la, Err: err} // c is non-nil interface containing nil pointer
// 	}
// 	return c, nil
// }