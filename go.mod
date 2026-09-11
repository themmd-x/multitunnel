module github.com/themmd-x/multitunnel

go 1.26.5

replace (
	gophertunnel/v1200 => ./libs/1.20.0
	gophertunnel/v12010 => ./libs/1.20.10
	gophertunnel/v12030 => ./libs/1.20.30
	gophertunnel/v12040 => ./libs/1.20.40
	gophertunnel/v12050 => ./libs/1.20.50
	gophertunnel/v12070 => ./libs/1.20.70
	gophertunnel/v12080 => ./libs/1.20.80
	gophertunnel/v121100 => ./libs/1.21.100
	gophertunnel/v121111 => ./libs/1.21.111
	gophertunnel/v121120 => ./libs/1.21.120
	gophertunnel/v121130 => ./libs/1.21.130
	gophertunnel/v1212 => ./libs/1.21.2
	gophertunnel/v12120 => ./libs/1.21.20
	gophertunnel/v12130 => ./libs/1.21.30
	gophertunnel/v12140 => ./libs/1.21.40
	gophertunnel/v12150 => ./libs/1.21.50
	gophertunnel/v12160 => ./libs/1.21.60
	gophertunnel/v12170 => ./libs/1.21.70
	gophertunnel/v12180 => ./libs/1.21.80
	gophertunnel/v12190 => ./libs/1.21.90
	gophertunnel/v12193 => ./libs/1.21.93
	gophertunnel/v12610 => ./libs/1.26.10
	gophertunnel/v12630 => ./libs/1.26.30
	gophertunnel/v12640 => ./libs/1.26.40
)

require (
	github.com/coreos/go-oidc/v3 v3.17.0 // indirect
	github.com/df-mc/go-nethernet v1.0.17 // indirect
	github.com/df-mc/go-playfab v1.0.0 // indirect
	github.com/df-mc/go-xsapi v1.0.1 // indirect
	github.com/df-mc/jsonc v1.0.5 // indirect
	github.com/go-gl/mathgl v1.1.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.3 // indirect
	github.com/go-jose/go-jose/v4 v4.1.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.18.1 // indirect
	github.com/muhammadmuzzammil1998/jsonc v1.0.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pion/datachannel v1.6.0 // indirect
	github.com/pion/dtls/v3 v3.1.2 // indirect
	github.com/pion/ice/v4 v4.2.1 // indirect
	github.com/pion/interceptor v0.1.44 // indirect
	github.com/pion/logging v0.2.4 // indirect
	github.com/pion/mdns/v2 v2.1.0 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.16 // indirect
	github.com/pion/rtp v1.10.1 // indirect
	github.com/pion/sctp v1.9.2 // indirect
	github.com/pion/sdp/v3 v3.0.18 // indirect
	github.com/pion/srtp/v3 v3.0.10 // indirect
	github.com/pion/stun/v3 v3.1.1 // indirect
	github.com/pion/transport/v4 v4.0.1 // indirect
	github.com/pion/turn/v4 v4.1.4 // indirect
	github.com/pion/webrtc/v4 v4.2.10-0.20260224155637-aa3b95c72dd2 // indirect
	github.com/sandertv/go-raknet v1.14.3-0.20250305181847-6af3e95113d6 // indirect
	github.com/sandertv/gophertunnel v1.43.0 // indirect
	github.com/wlynxg/anet v0.0.5 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.48.0 // indirect
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0 // indirect
	golang.org/x/image v0.21.0 // indirect
	golang.org/x/net v0.50.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	golang.org/x/time v0.10.0 // indirect
	gophertunnel/v1200 v1200.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12010 v12010.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12030 v12030.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12040 v12040.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12050 v12050.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12070 v12070.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12080 v12080.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v121100 v121100.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v121111 v121111.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v121120 v121120.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v121130 v121130.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v1212 v1212.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12120 v12120.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12130 v12130.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12140 v12140.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12150 v12150.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12160 v12160.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12170 v12170.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12180 v12180.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12190 v12190.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12193 v12193.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12610 v12610.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12630 v12630.0.0-00010101000000-000000000000 // indirect
	gophertunnel/v12640 v12640.0.0-00010101000000-000000000000 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
