package net_struct

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestMakeHeader(t *testing.T) {
	var buffer  = bytes.NewBuffer([]byte{})
	binary.Write(buffer,binary.LittleEndian,MakeHeader(100))
	t.Log(buffer.Bytes())
}