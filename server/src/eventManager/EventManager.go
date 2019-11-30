package eventManager

import "redisProject/src/net_struct"

type EventManager struct {
	readChan chan net_struct.TCPServerData
}