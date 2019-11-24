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
