package pack

//编码
func Decode(value interface{}) []byte {
	return Pack_common(value)
}
// 解码
func Encode(bts []byte) interface{}  {
	bs ,value:= UNPack_common(bts)
	if len(bs) !=0 {
		panic("encode error "+ string(len(bs)))
	}
	return  value
}