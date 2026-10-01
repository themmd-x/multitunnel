package core

import (
	"io"

	"github.com/TheMMD-X/multitunnel/minecraft/resource"
)

func PackReader(pack *resource.Pack) io.Reader {
	return io.NewSectionReader(pack, 0, int64(pack.Len()))
}
