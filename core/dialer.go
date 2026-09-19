package core

import (
	"net/http"
	"log/slog"
	"time"

	"golang.org/x/oauth2"
	"github.com/df-mc/go-playfab/v2"
	"github.com/df-mc/go-xsapi/v2"
	"github.com/TheMMD-X/multitunnel/minecraft/protocol/login"
)

type Dialer struct {
        ErrorLog                   *slog.Logger
        HTTPClient                 *http.Client
        ClientData                 login.ClientData
        IdentityData               login.IdentityData
        TokenSource                oauth2.TokenSource
        XBLClient                  *xsapi.Client
        PlayFabClient              *playfab.Client
        //PacketFunc func(header packet.Header, payload []byte, src, dst net.Addr)
        //DownloadResourcePack func(id uuid.UUID, version string, current, total int) bool
        //ResourcePackCache          T
        DisconnectOnUnknownPackets bool
        DisconnectOnInvalidPackets bool
		// tf is the point of multitunnel with this existing in it?
        //Protocol                   T
        FlushRate                  time.Duration
        EnableClientCache          bool
        KeepXBLIdentityData        bool
        EnableLegacyAuth           bool
        Version                    string
}
