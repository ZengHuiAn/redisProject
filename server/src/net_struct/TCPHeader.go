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


type TCPClientData struct {
	Header TCPClientHeader // 头部信息
	body []byte
}

func NewTCPClientData(header TCPClientHeader, body []byte) *TCPClientData {
	return &TCPClientData{Header: header, body: body}
}
func (SELF* TCPClientData)GetBody() []byte  {
	return SELF.body
}


type TCPServerTargetAddr struct {
	fromServer string
	toServer string
}




type TCPServerData struct {
	Target TCPServerTargetAddr
	ClientData TCPClientData
}

