package minecraft

import (
	"errors"
	"net"
	"reflect"
	"strings"
)

var errBufferTooSmall = errors.New("a message sent was larger than the buffer used to receive the message into")

type DisconnectError string

func (d DisconnectError) Error() string {
	return string(d)
}

func asDisconnect(err error) (DisconnectError, bool) {
	if err == nil {
		return "", false
	}
	if d, ok := err.(DisconnectError); ok {
		return d, true
	}
	t := reflect.TypeOf(err)
	if t.Kind() != reflect.String || t.Name() != "DisconnectError" || !strings.Contains(t.PkgPath(), "gophertunnel") {
		return "", false
	}
	return DisconnectError(reflect.ValueOf(err).String()), true
}

func wrapError(err error) error {
	if err == nil {
		return nil
	}
	if d, ok := asDisconnect(err); ok {
		return d
	}
	var opErr *net.OpError
	if !errors.As(err, &opErr) {
		return err
	}
	if _, ours := opErr.Err.(DisconnectError); ours {
		return err
	}
	d, ok := asDisconnect(opErr.Err)
	if !ok {
		return err
	}
	converted := *opErr
	converted.Err = d
	return &converted
}
