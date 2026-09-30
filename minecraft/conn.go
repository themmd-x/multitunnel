package minecraft

import (
	"fmt"
	"path"
	"reflect"

	"github.com/TheMMD-X/multitunnel/minecraft/protocol/packet"
	"github.com/TheMMD-X/multitunnel/versions"
)

type Spawner interface {
	DoSpawn() error
	Close() error
}

type Handle interface {
	DoSpawn() error
	Close() error
}

type Conn[T Spawner] struct {
	GtConn  T
	Version string
}

// TODO: instead of this unversioned mess of wrappers move it to the version adapters
func (c *Conn[T]) DoSpawn() error {
	return c.GtConn.DoSpawn()
}

func (c *Conn[T]) Close() error {
	return c.GtConn.Close()
}

func (c *Conn[T]) WritePacket(pk packet.Packet) error {
	return writePacket(c.GtConn, c.Version, pk)
}

func (c *Conn[T]) ReadPacket() (packet.Packet, error) {
	return ReadPacket(c.GtConn)
}

func ReadPacket(conn any) (packet.Packet, error) {
	out := reflect.ValueOf(conn).MethodByName("ReadPacket").Call(nil)
	if e := out[1].Interface(); e != nil {
		return nil, e.(error)
	}
	return fromVendor(out[0].Interface())
}

func WritePacket(conn any, pk packet.Packet) error {
	version, err := versionOf(conn)
	if err != nil {
		return err
	}
	return writePacket(conn, version, pk)
}

func writePacket(conn any, version string, pk packet.Packet) error {
	vendorPk, err := toVendor(pk, version)
	if err != nil {
		return err
	}
	out := reflect.ValueOf(conn).MethodByName("WritePacket").Call([]reflect.Value{reflect.ValueOf(vendorPk)})
	if e := out[0].Interface(); e != nil {
		return e.(error)
	}
	return nil
}

func StartGame(conn any, data GameData) error {
	m := reflect.ValueOf(conn).MethodByName("StartGame")
	if !m.IsValid() {
		return fmt.Errorf("minecraft: %T has no StartGame method", conn)
	}
	vendorData := reflect.New(m.Type().In(0)).Elem()
	copyValue(reflect.ValueOf(data), vendorData)
	out := m.Call([]reflect.Value{vendorData})
	if e := out[0].Interface(); e != nil {
		return e.(error)
	}
	return nil
}

func versionOf(conn any) (string, error) {
	t := reflect.TypeOf(conn)
	for t != nil && t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t == nil {
		return "", fmt.Errorf("minecraft: cannot detect the version of a nil connection")
	}
	id := path.Base(path.Dir(t.PkgPath()))
	v, ok := versions.ByID[id]
	if !ok {
		return "", fmt.Errorf("minecraft: %T does not come from a known gophertunnel version", conn)
	}
	return v.Minecraft, nil
}
