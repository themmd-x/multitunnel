package minecraft

//import "github.com/TheMMD-X/multitunnel/minecraft/protocol/packet"

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

// do we need adapters for these? no idea
func (c *Conn[T]) DoSpawn() error {
	return c.GtConn.DoSpawn()
}

func (c *Conn[T]) Close() error {
	return c.GtConn.Close()
}
/*
func (c *Conn[T]) WritePacket(pk packet.Packet) error {
	return c.GtConn.WritePacket(pk)
}

func (c *Conn[T]) ReadPacket[P any]() (pk P, err error) {
	return c.GtConn.ReadPacket()
}*/
