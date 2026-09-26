package v12120

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"

	mtCore "github.com/TheMMD-X/multitunnel/core"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12120/minecraft"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12120/minecraft/protocol/login"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12120/minecraft/protocol/packet"
)

func Connect(host string, mtDialer mtCore.Dialer) (*minecraft.Conn, error) {
	var errorLog *log.Logger
	if mtDialer.ErrorLog != nil {
		errorLog = slog.NewLogLogger(mtDialer.ErrorLog.Handler(), slog.LevelError)
	}

	thirdPartyNameOnly := false
	if mtDialer.ClientData.ThirdPartyNameOnly != nil {
		thirdPartyNameOnly = *mtDialer.ClientData.ThirdPartyNameOnly
	}

	dialer := minecraft.Dialer{
		ErrorLog: errorLog,
		//HTTPClient: mtDialer.HTTPClient,
		ClientData: login.ClientData{
			CapeData: mtDialer.ClientData.CapeData,
			CapeID: mtDialer.ClientData.CapeID,
			CapeImageHeight: mtDialer.ClientData.CapeImageHeight,
			CapeImageWidth: mtDialer.ClientData.CapeImageWidth,
			CapeOnClassicSkin: mtDialer.ClientData.CapeOnClassicSkin,
			ClientRandomID: mtDialer.ClientData.ClientRandomID,
			CurrentInputMode: mtDialer.ClientData.CurrentInputMode,
			DefaultInputMode: mtDialer.ClientData.DefaultInputMode,
			GameVersion: mtDialer.ClientData.GameVersion,
			GUIScale: mtDialer.ClientData.GUIScale,
			//FilterProfanity: mtDialer.ClientData.FilterProfanity,
			//ClientEditorConnectionIntent: mtDialer.ClientData.ClientEditorConnectionIntent,
			//ClientIsEditorCapable: mtDialer.ClientData.ClientIsEditorCapable,
			LanguageCode: mtDialer.ClientData.LanguageCode,
			PersonaSkin: mtDialer.ClientData.PersonaSkin,
			PlatformOfflineID: mtDialer.ClientData.PlatformOfflineID,
			PlatformOnlineID: mtDialer.ClientData.PlatformOnlineID,
			PlatformUserID: mtDialer.ClientData.PlatformUserID,
			PremiumSkin: mtDialer.ClientData.PremiumSkin,
			SelfSignedID: mtDialer.ClientData.SelfSignedID,
			ServerAddress: mtDialer.ClientData.ServerAddress,
			SkinAnimationData: mtDialer.ClientData.SkinAnimationData,
			SkinData: mtDialer.ClientData.SkinData,
			SkinGeometry: mtDialer.ClientData.SkinGeometry,
			SkinGeometryVersion: mtDialer.ClientData.SkinGeometryVersion,
			SkinID: mtDialer.ClientData.SkinID,
			PlayFabID: mtDialer.ClientData.PlayFabID,
			SkinImageHeight: mtDialer.ClientData.SkinImageHeight,
			SkinImageWidth: mtDialer.ClientData.SkinImageWidth,
			SkinResourcePatch: mtDialer.ClientData.SkinResourcePatch,
			SkinColour: mtDialer.ClientData.SkinColour,
			ArmSize: mtDialer.ClientData.ArmSize,
			ThirdPartyName: mtDialer.ClientData.ThirdPartyName,
			ThirdPartyNameOnly: thirdPartyNameOnly,
			UIProfile: mtDialer.ClientData.UIProfile,
			TrustedSkin: mtDialer.ClientData.TrustedSkin,
			OverrideSkin: mtDialer.ClientData.OverrideSkin,
			CompatibleWithClientSideChunkGen: mtDialer.ClientData.CompatibleWithClientSideChunkGen,
			//MaxViewDistance: mtDialer.ClientData.MaxViewDistance,
			//MemoryTier: mtDialer.ClientData.MemoryTier,
			//PlatformType: mtDialer.ClientData.PlatformType,
			//GraphicsMode: mtDialer.ClientData.GraphicsMode,
			//PartyID: mtDialer.ClientData.PartyID,
			//PartyLeader: mtDialer.ClientData.PartyLeader,
			//ProfileHash: mtDialer.ClientData.ProfileHash,
			//Nonce: mtDialer.ClientData.Nonce,
		},
		IdentityData: login.IdentityData{
			XUID: mtDialer.IdentityData.XUID,
			Identity: mtDialer.IdentityData.Identity,
			DisplayName: mtDialer.IdentityData.DisplayName,
			TitleID: mtDialer.IdentityData.TitleID,
			//PlayFabTitleID: mtDialer.IdentityData.PlayFabTitleID,
			//PlayFabID: mtDialer.IdentityData.PlayFabID,
		},
		TokenSource: mtDialer.TokenSource,
		//XBLClient: mtDialer.XBLClient,
		//PlayFabClient: mtDialer.PlayFabClient,
		//PacketFunc: mtDialer.PacketFunc,
		//DownloadResourcePack: mtDialer.DownloadResourcePack,
		//ResourcePackCache: mtDialer.ResourcePackCache,
		DisconnectOnUnknownPackets: mtDialer.DisconnectOnUnknownPackets,
		DisconnectOnInvalidPackets: mtDialer.DisconnectOnInvalidPackets,
		//Protocol: mtDialer.Protocol,
		FlushRate: mtDialer.FlushRate,
		EnableClientCache: mtDialer.EnableClientCache,
		KeepXBLIdentityData: mtDialer.KeepXBLIdentityData,
		//EnableLegacyAuth: mtDialer.EnableLegacyAuth,
	}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

var (
	feeders      sync.Map
	feederSerial atomic.Uint64
)

func init() {
	minecraft.RegisterNetwork("multitunnel", network{})
}

type network struct{}

func (network) DialContext(context.Context, string) (net.Conn, error) {
	return nil, errors.New("v12120: the multitunnel network cannot dial")
}

func (network) PingContext(context.Context, string) ([]byte, error) {
	return nil, errors.New("v12120: the multitunnel network cannot ping")
}

func (network) Listen(address string) (minecraft.NetworkListener, error) {
	f, ok := feeders.LoadAndDelete(address)
	if !ok {
		return nil, fmt.Errorf("v12120: no feeder registered under %q", address)
	}
	return f.(*feeder), nil
}

type feeder struct {
	nl     mtCore.NetworkListener
	conns  chan net.Conn
	closed chan struct{}
	once   sync.Once
}

func (f *feeder) Accept() (net.Conn, error) {
	select {
	case conn := <-f.conns:
		return conn, nil
	case <-f.closed:
		return nil, net.ErrClosed
	}
}

func (f *feeder) Close() error {
	f.once.Do(func() { close(f.closed) })
	return nil
}

func (f *feeder) Addr() net.Addr {
	return f.nl.Addr()
}

func (f *feeder) ID() int64 {
	return f.nl.ID()
}

func (f *feeder) PongData(data []byte) {
	f.nl.PongData(data)
}

type Server struct {
	listener *minecraft.Listener
	feeder   *feeder
}

func NewServer(mtCfg mtCore.ListenConfig, nl mtCore.NetworkListener, deliver func(net.Conn)) (*Server, error) {
	f := &feeder{
		nl:     nl,
		conns:  make(chan net.Conn),
		closed: make(chan struct{}),
	}
	key := fmt.Sprint(feederSerial.Add(1))
	feeders.Store(key, f)

	var errorLog *log.Logger
	if mtCfg.ErrorLog != nil {
		errorLog = slog.NewLogLogger(mtCfg.ErrorLog.Handler(), slog.LevelError)
	}

	cfg := minecraft.ListenConfig{
		ErrorLog:               errorLog,
		AuthenticationDisabled: mtCfg.AuthenticationDisabled,
		MaximumPlayers:         mtCfg.MaximumPlayers,
		AllowUnknownPackets:    mtCfg.AllowUnknownPackets,
		AllowInvalidPackets:    mtCfg.AllowInvalidPackets,
		FlushRate:              mtCfg.FlushRate,
		TexturePacksRequired:   mtCfg.TexturePacksRequired,
	}
	if mtCfg.StatusProvider != nil {
		cfg.StatusProvider = statusProvider{mtCfg.StatusProvider}
	}

	listener, err := cfg.Listen("multitunnel", key)
	if err != nil {
		feeders.Delete(key)
		return nil, err
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			deliver(conn)
		}
	}()
	return &Server{listener: listener, feeder: f}, nil
}

func (s *Server) Handle(conn net.Conn) {
	select {
	case s.feeder.conns <- conn:
	case <-s.feeder.closed:
		_ = conn.Close()
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func Disconnect(conn net.Conn, message string) error {
	c, ok := conn.(*minecraft.Conn)
	if !ok {
		return fmt.Errorf("v12120: disconnect: unexpected connection type %T", conn)
	}
	_ = c.WritePacket(&packet.Disconnect{
		HideDisconnectionScreen: message == "",
		Message:                 message,
	})
	return c.Close()
}

type statusProvider struct {
	provider mtCore.ServerStatusProvider
}

func (s statusProvider) ServerStatus(playerCount, maxPlayers int) minecraft.ServerStatus {
	status := s.provider.ServerStatus(playerCount, maxPlayers)
	return minecraft.ServerStatus{
		ServerName:    status.ServerName,
		ServerSubName: status.ServerSubName,
		PlayerCount:   status.PlayerCount,
		MaxPlayers:    status.MaxPlayers,
	}
}
