package minecraft

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"sync"

	mtCore "github.com/TheMMD-X/multitunnel/core"
	"github.com/TheMMD-X/multitunnel/minecraft/internal"
	"github.com/TheMMD-X/multitunnel/versions"
	"github.com/TheMMD-X/multitunnel/versions/v121100"
	"github.com/TheMMD-X/multitunnel/versions/v121111"
	"github.com/TheMMD-X/multitunnel/versions/v121120"
	"github.com/TheMMD-X/multitunnel/versions/v121130"
	"github.com/TheMMD-X/multitunnel/versions/v1212"
	"github.com/TheMMD-X/multitunnel/versions/v12120"
	"github.com/TheMMD-X/multitunnel/versions/v12130"
	"github.com/TheMMD-X/multitunnel/versions/v12140"
	"github.com/TheMMD-X/multitunnel/versions/v12150"
	"github.com/TheMMD-X/multitunnel/versions/v12160"
	"github.com/TheMMD-X/multitunnel/versions/v12170"
	"github.com/TheMMD-X/multitunnel/versions/v12180"
	"github.com/TheMMD-X/multitunnel/versions/v12190"
	"github.com/TheMMD-X/multitunnel/versions/v12193"
	"github.com/TheMMD-X/multitunnel/versions/v12610"
	"github.com/TheMMD-X/multitunnel/versions/v12630"
	"github.com/TheMMD-X/multitunnel/versions/v12640"
    raknet "github.com/TheMMD-X/multitunnel/libs/go-raknet/v1152"
)

type ListenConfig struct {
	mtCore.ListenConfig
}

type Listener struct {
	mtCore.Listener
	config  mtCore.ListenConfig
	network mtCore.NetworkListener
	mu      sync.RWMutex
	servers map[string]server
}

type server interface {
	Handle(conn net.Conn)
	Close() error
}

type (
	ServerStatus           = mtCore.ServerStatus
	ServerStatusProvider   = mtCore.ServerStatusProvider
	ListenerStatusProvider = mtCore.ListenerStatusProvider
)

func NewStatusProvider(serverName, serverSubName string) ListenerStatusProvider {
	return mtCore.NewStatusProvider(serverName, serverSubName)
}

func (cfg ListenConfig) Listen(network string, address string) (*Listener, error) {
	if network != "raknet" {
		return nil, fmt.Errorf("listen: no network under id %v", network)
	}
	if cfg.ErrorLog == nil {
		cfg.ErrorLog = slog.New(internal.DiscardHandler{})
	}
	if cfg.StatusProvider == nil {
		cfg.StatusProvider = mtCore.NewStatusProvider("Minecraft Server", "Gophertunnel")
	}

	nl, err := raknet.ListenConfig{ErrorLog: cfg.ErrorLog.With("net origin", "raknet")}.Listen(address)
	if err != nil {
		return nil, err
	}

	listener := &Listener{
		config:  cfg.ListenConfig,
		network: nl,
		servers: make(map[string]server),
	}
	listener.Listener.Init(cfg.ListenConfig, nl)
	listener.updatePongData()
	go listener.Listener.Serve(listener.handle)
	return listener, nil
}

func Listen(network, address string) (*Listener, error) {
	var lc ListenConfig
	return lc.Listen(network, address)
}

func (listener *Listener) Accept() (net.Conn, error) {
	return listener.Listener.Accept()
}

func (listener *Listener) Disconnect(conn net.Conn, message string) error {
	return v12140.Disconnect(conn, message)
}

func (listener *Listener) Close() error {
	err := listener.Listener.Close()
	listener.mu.Lock()
	for _, s := range listener.servers {
		_ = s.Close()
	}
	listener.mu.Unlock()
	return err
}

func (listener *Listener) handle(conn net.Conn) {
	version := listener.config.Version
	if version == "" {
		detected, err := detectVersion(conn)
		if err != nil {
			listener.config.ErrorLog.Error("detect client version", "error", err)
			_ = conn.Close()
			return
		}
		version = detected
	}

	s, err := listener.serverFor(version)
	if err != nil {
		listener.config.ErrorLog.Error("create version server", "error", err)
		_ = conn.Close()
		return
	}
	s.Handle(conn)
}

func (listener *Listener) serverFor(version string) (server, error) {
	listener.mu.RLock()
	s, ok := listener.servers[version]
	listener.mu.RUnlock()
	if ok {
		return s, nil
	}

	listener.mu.Lock()
	defer listener.mu.Unlock()
	if s, ok := listener.servers[version]; ok {
		return s, nil
	}

	switch version {
	case "1.21.2":
		created, err := v1212.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.20":
		created, err := v12120.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.30":
		created, err := v12130.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.40":
		created, err := v12140.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.50":
		created, err := v12150.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.60":
		created, err := v12160.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.70":
		created, err := v12170.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.80":
		created, err := v12180.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.90":
		created, err := v12190.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.93":
		created, err := v12193.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.100":
		created, err := v121100.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.111":
		created, err := v121111.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.120":
		created, err := v121120.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.21.130":
		created, err := v121130.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.26.10":
		created, err := v12610.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.26.30":
		created, err := v12630.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil

	case "1.26.40":
		created, err := v12640.NewServer(listener.config, listener.network, listener.Deliver)
		if err != nil {
			return nil, err
		}
		listener.servers[version] = created
		return created, nil
	default:
		return nil, fmt.Errorf("minecraft: version %q is not supported yet", version)
	}
}

func detectVersion(conn net.Conn) (string, error) {
	return "", errors.New("minecraft: client version detection is not implemented")
}

func (listener *Listener) updatePongData() {
	version, ok := versions.ByMinecraft[listener.config.Version]
	if !ok {
		return
	}
	status := listener.config.StatusProvider.ServerStatus(0, listener.config.MaximumPlayers)
	if status.MaxPlayers == 0 {
		status.MaxPlayers = status.PlayerCount + 1
	}
	port := listener.Addr().(*net.UDPAddr).Port
	listener.network.PongData([]byte(fmt.Sprintf("MCPE;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;",
		status.ServerName, version.Protocol, version.Minecraft, status.PlayerCount, status.MaxPlayers,
		listener.network.ID(), status.ServerSubName, "Creative", 1, port, port, 0,
	)))
}
