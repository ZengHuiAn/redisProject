package net_struct

type TCPClientHeader struct {
	Length    uint32
	Flag      uint32 //代表客户端发来的信息
	MessageID uint32
	ProtoType uint32
}

const ClientClientHeaderLength = 16

func MakeHeader(msgID uint32) TCPClientHeader {
	return TCPClientHeader{MessageID: msgID, Flag: 1}
}



type TCPServerData struct {
	Header TCPClientHeader // 头部信息
	body []byte
}

func (SELF* TCPServerData)GetBody() []byte  {
	return SELF.body
}