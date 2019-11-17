package net_struct

type (
	TCPStruct struct {
		Host         string
		Port         int
		ProtocolType string
		Name         string
	}
)

func (T TCPStruct) GetPort() int {
	return T.Port
}


func (T TCPStruct) GetHost() string {
	if T.Host == "" {
		T.Host = "0.0.0.0"
	}

	return T.Host
}