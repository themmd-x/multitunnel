package minecraft

import (
	"fmt"

	gtV121100 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121100/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v121100"

	gtV121111 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121111/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v121111"

	gtV121120 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121120/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v121120"

	gtV121130 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121130/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v121130"

	gtV1212 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v1212/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v1212"

	gtV12120 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12120/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12120"

	gtV12130 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12130/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12130"

	gtV12140 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12140/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12140"

	gtV12150 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12150/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12150"

	gtV12160 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12160/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12160"

	gtV12170 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12170/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12170"

	gtV12180 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12180/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12180"

	gtV12190 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12190/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12190"

	gtV12193 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12193/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12193"

	gtV12610 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12610/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12610"

	gtV12630 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12630/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12630"

	gtV12640 "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12640/minecraft"
	"github.com/TheMMD-X/multitunnel/versions/v12640"
)

// TODO: Add the gophertunnel dialer fully instead of making shit up
type Dialer struct {
	Version string
}

func (d *Dialer) Dial(network string, host string) (Handle, error) {
	if network != "raknet" {
		return nil, fmt.Errorf("minecraft: unsupported network %q", network)
	}

	switch d.Version {
	case "1.21.100":
		gt, err := v121100.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV121100.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.111":
		gt, err := v121111.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV121111.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.120":
		gt, err := v121120.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV121120.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.130":
		gt, err := v121130.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV121130.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.2":
		gt, err := v1212.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV1212.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.20":
		gt, err := v12120.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12120.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.30":
		gt, err := v12130.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12130.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.40":
		gt, err := v12140.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12140.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.50":
		gt, err := v12150.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12150.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.60":
		gt, err := v12160.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12160.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.70":
		gt, err := v12170.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12170.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.80":
		gt, err := v12180.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12180.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.90":
		gt, err := v12190.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12190.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.21.93":
		gt, err := v12193.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12193.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.26.10":
		gt, err := v12610.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12610.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.26.30":
		gt, err := v12630.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12630.Conn]{GtConn: gt, Version: d.Version}, nil

	case "1.26.40":
		gt, err := v12640.Connect(host)
		if err != nil {
			return nil, err
		}
		return &Conn[*gtV12640.Conn]{GtConn: gt, Version: d.Version}, nil
	default:
		return nil, fmt.Errorf("minecraft: version %q unknown", d.Version)
	}
}
