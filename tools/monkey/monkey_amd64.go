package monkey

import (
	"bytes"
	"encoding/binary"
)

func getJumpCode(f uintptr) []byte {
	buf := &bytes.Buffer{}
	buf.WriteByte(0xba)
	binary.Write(buf, binary.LittleEndian, uint32(f))
	buf.WriteByte(0xff)
	buf.WriteByte(0xe2)
	return buf.Bytes()
}
