package core

import (

        "crypto/ecdsa"
        "log/slog"
        "net"
        "net/http"
        //"slices"
        "sync"
        "sync/atomic"
        "time"

		//"github.com/TheMMD-X/multitunnel/minecraft/internal"
		//"github.com/TheMMD-X/multitunnel/minecraft/protocol"                                                          
		//mtCore "github.com/TheMMD-X/multitunnel/core"                                                            
		"github.com/coreos/go-oidc/v3/oidc"
		"github.com/TheMMD-X/multitunnel/minecraft/protocol/login"                                                            
		"github.com/TheMMD-X/multitunnel/minecraft/protocol/packet"                                                           
		"github.com/TheMMD-X/multitunnel/minecraft/resource"                                                                  
		//"github.com/TheMMD-X/multitunnel/minecraft/service"
)

/*type NetworkListener interface {
	net.Listener
	ID() int64
	PongData([]byte)
}*/

type Listener struct {
        cfg      ListenConfig
        listener NetworkListener

        packs   []*resource.Pack
        packsMu sync.RWMutex

        playerCount atomic.Int32

        incoming chan *net.Conn
        close    chan struct{}
        closeOnce *sync.Once

        key      *ecdsa.PrivateKey
        verifier *oidc.IDTokenVerifier
}

type ListenConfig struct {
        ErrorLog                   *slog.Logger
        HTTPClient                 *http.Client
        AuthenticationDisabled     bool
        DisablePacketEncryption    bool
        MaximumPlayers             int
        AllowUnknownPackets        bool
        AllowInvalidPackets        bool
        StatusProvider             ServerStatusProvider
        //AcceptedProtocols          []Protocol
        Compression                packet.Compression
        //CompressionSelector        func(proto Protocol) packet.Compression
        CompressionThreshold       int
        FlushRate                  time.Duration
        ResourcePacks              []*resource.Pack
        TexturePacksRequired       bool
        ForceDisableVibrantVisuals bool
        FetchResourcePacks         func(identityData login.IdentityData, clientData login.ClientData, current []*resource.Pack) []*resource.Pack
        //ResourcePackDelivery       ResourcePackDeliveryConfig
        PacketFunc                 func(header packet.Header, payload []byte, src, dst net.Addr)
        MaxDecompressedLen         int
        Allow                      func(addr net.Addr, identityData login.IdentityData, clientData login.ClientData) (string, bool)
        Version                    string
}

func (l Listener) Incoming() <-chan *net.Conn {
	return l.incoming
}

func (l Listener) Addr() net.Addr {
	return l.listener.Addr()
}

func (l *Listener) Init(cfg ListenConfig, nl NetworkListener) {
	l.cfg = cfg
	l.listener = nl
	l.incoming = make(chan *net.Conn)
	l.close = make(chan struct{})
	l.closeOnce = new(sync.Once)
}

func (l *Listener) Serve(handle func(conn net.Conn)) {
	for {
		conn, err := l.listener.Accept()
		if err != nil {
			l.closeOnce.Do(func() { close(l.close) })
			return
		}
		go handle(conn)
	}
}

func (l *Listener) Deliver(conn net.Conn) {
	select {
	case l.incoming <- &conn:
	case <-l.close:
		_ = conn.Close()
	}
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.incoming:
		return *conn, nil
	case <-l.close:
		return nil, &net.OpError{Op: "accept", Net: "minecraft", Addr: l.Addr(), Err: net.ErrClosed}
	}
}

func (l *Listener) Close() error {
	if l.listener == nil {
		return net.ErrClosed
	}
	l.closeOnce.Do(func() { close(l.close) })
	return l.listener.Close()
}

type ServerStatusProvider interface {
	ServerStatus(playerCount, maxPlayers int) ServerStatus
}

type ServerStatus struct {
	ServerName    string
	ServerSubName string
	PlayerCount   int
	MaxPlayers    int
}

type ListenerStatusProvider struct {
	name    string
	subName string
}

func NewStatusProvider(serverName, serverSubName string) ListenerStatusProvider {
	return ListenerStatusProvider{name: serverName, subName: serverSubName}
}

func (l ListenerStatusProvider) ServerStatus(playerCount, maxPlayers int) ServerStatus {
	return ServerStatus{
		ServerName:    l.name,
		ServerSubName: l.subName,
		PlayerCount:   playerCount,
		MaxPlayers:    maxPlayers,
	}
}
