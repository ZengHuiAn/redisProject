package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"redisProject/src/net_struct"
	"redisProject/src/pack"
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

func testWriter(writer *bufio.Writer)  {
	var header net_struct.TCPClientHeader = net_struct.TCPClientHeader{Length:net_struct.ClientClientHeaderLength,Flag:1,MessageID:101,ProtoType:1}
	var buffer = bytes.NewBuffer([]byte{})
	err:= binary.Write(buffer,binary.LittleEndian,header)
	if err != nil {
		log.Println("writer error ", err)
		return
	}

	_, err = writer.Write(buffer.Bytes())
	_ = writer.Flush()
	if err != nil {
		log.Println("writer error ", err)
		return
	}
	fmt.Println("发送数据")
}

func handleRequest(conn net.Conn) {
	ip := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(fmt.Sprintf("disconnect: %s", ip))
		conn.Close()
	}()


	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	testWriter(writer)
	for {
		headerBuf := make([]byte, net_struct.ClientClientHeaderLength)
		fmt.Println("读取数据---------->>>>",ip)
		readLen, err := reader.Read(headerBuf)
		if err!=nil || err == io.EOF {
			log.Println("read error ", err)
			break
		}


		var buffer  = bytes.NewBuffer(headerBuf)
		var header net_struct.TCPClientHeader
		err = binary.Read(buffer, binary.LittleEndian, &header)

		if err!=nil  {
			log.Println("read buffer error ", err)
			break
		}

		fmt.Println("read data :",header)
		dataBuffer := make([]byte, header.Length - net_struct.ClientClientHeaderLength);
		readLen, err = reader.Read(dataBuffer)

		if err!=nil || err == io.EOF {
			log.Println("read error ", err)
			break
		}

		fmt.Println("read data :",dataBuffer)
		var packData = pack.Encode(dataBuffer)
		fmt.Println("read data :",packData)
		fmt.Println(fmt.Sprintf(" read success length : %d, msg : %v", readLen, dataBuffer))
	}

}
