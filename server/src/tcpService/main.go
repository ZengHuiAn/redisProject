package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"redisProject/src/net_struct"
	"strconv"
)

var s_instance = net_struct.ServerInstance{
	Config: &net_struct.TCPStruct{
		Name:         "ConnectServer",
		Port:         10000,
		ProtocolType: "tcp",
		Host:         "127.0.0.1",
	},
}

func main() {

	fmt.Println(s_instance.Config.GetHost() + ":" + strconv.Itoa(s_instance.Config.Port))
	var listener, err = net.Listen(s_instance.Config.ProtocolType, s_instance.Config.GetHost()+":"+strconv.Itoa(s_instance.Config.Port))

	if err != nil {
		log.Fatal("lister error ", err)
		return
	}

	defer listener.Close()
	fmt.Println(fmt.Sprintf("ProtocolType %s, addr %s", listener.Addr().Network(), listener.Addr().String()))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error ", err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("message %s -> %s", conn.RemoteAddr(), conn.LocalAddr()))
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	ip := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(fmt.Sprintf("disconnect: %s", ip))
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		headerBuf := make([]byte, 4)
		fmt.Println("读取数据---------->>>>",ip)
		readLen, err := reader.Read(headerBuf)
		if err!=nil || err == io.EOF {
			log.Println("read error ", err)
			break
		}

		fmt.Println(fmt.Sprintf(" read success length : %d, msg : %s", readLen, headerBuf))
		writerLen, err := writer.Write(headerBuf)
		_ = writer.Flush()
		if err != nil {
			log.Println("writer error ", err)
			return
		}

		fmt.Println(fmt.Sprintf(" read success length : %d, msg : %s", writerLen, headerBuf))

	}

}
