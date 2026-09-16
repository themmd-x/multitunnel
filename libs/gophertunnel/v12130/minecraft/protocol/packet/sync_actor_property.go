package packet

import (
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12130/minecraft/nbt"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12130/minecraft/protocol"
)

// SyncActorProperty is an alternative to synced actor data.
type SyncActorProperty struct {
	// PropertyData ...
	PropertyData map[string]any
}

// ID ...
func (*SyncActorProperty) ID() uint32 {
	return IDSyncActorProperty
}

func (pk *SyncActorProperty) Marshal(io protocol.IO) {
	io.NBT(&pk.PropertyData, nbt.NetworkLittleEndian)
}
