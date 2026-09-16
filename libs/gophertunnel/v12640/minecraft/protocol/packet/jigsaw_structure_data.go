package packet

import (
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12640/minecraft/nbt"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12640/minecraft/protocol"
)

// JigsawStructureData is sent by the server to let the client know all the rules for jigsaw structures.
type JigsawStructureData struct {
	// StructureData is a network NBT serialised compound of all the jigsaw structure rules defined
	// on the server.
	StructureData map[string]any
}

// ID ...
func (*JigsawStructureData) ID() uint32 {
	return IDJigsawStructureData
}

func (pk *JigsawStructureData) Marshal(io protocol.IO) {
	io.NBT(&pk.StructureData, nbt.NetworkLittleEndian)
}
