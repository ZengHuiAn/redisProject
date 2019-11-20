package pack

import "testing"

func TestPack_bool_data(t *testing.T) {
	t.Log(Pack_bool_data(true))
}


// pack byte to do

func TestPack_byte_data(t *testing.T)  {
	var bs = make([]byte,2)
	//byte[] numArray = new byte[2];
	t.Log(Pack_byte_data(bs))
}