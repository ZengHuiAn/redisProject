package pack

import (
	"testing"
)

func TestPack_bool_data(t *testing.T) {
	t.Log(Pack_bool_data(true))
	t.Log(Pack_rune_data('安'))
	t.Log(Pack_int16_data(int16(50)))
	//Pack_int16_data
	//runeData:= Pack_rune_data('安')
}