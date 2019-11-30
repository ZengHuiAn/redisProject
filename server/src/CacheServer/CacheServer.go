package CacheServer

import "errors"

type MicroserviceAddr struct {
	Host string
	Port string
	Name string
}

type CacheServer struct {
	Microservices map[string]*MicroserviceAddr
}

func (server * CacheServer) RegedisterServices(name string, addr *MicroserviceAddr) bool  {
	if server.Microservices[name] != nil  {
		return false
	}
	server.Microservices[name] = addr

	return  true
}

func (server *CacheServer) RemoveServices (name string) error {
	if server.Microservices[name] == nil  {
		return errors.New("未注册服务!")
	}
	server.Microservices[name] = nil
	return nil
}