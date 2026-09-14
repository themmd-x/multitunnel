package client

import (
	"github.com/TheMMD-X/multitunnel/api/adaptor/v12140"
	"github.com/TheMMD-X/multitunnel/api/minecraft"
)

func Connect(version string, host string, dialer *minecraft.Dialer) {
	switch version {
	case "1.21.40":
		err := v12140.Connect(host, dialer)
		if err != nil {
			panic(err)
		}
	}
}
