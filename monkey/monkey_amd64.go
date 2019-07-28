package monkey

import (
	"bytes"
	"encoding/binary"
)

func getJumpCode(f uintptr) []byte {
	buf := &bytes.Buffer{}
	buf.WriteByte(0x48)
	buf.WriteByte(0xba)
	binary.Write(buf, binary.LittleEndian, uint64(f))
	buf.WriteByte(0xff)
	buf.WriteByte(0x22)
	return buf.Bytes()
}
