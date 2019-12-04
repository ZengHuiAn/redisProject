package client

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"redisProject/src/eventManager"
	"redisProject/src/net_struct"
	"redisProject/src/static/res"
	"sync"
)

type TAGGER struct {
	TAG string
}

// 客户端---->>>
type CustomClient struct {
	tag          TAGGER
	conn         net.Conn
	wg           sync.WaitGroup
	writerChanel chan *net_struct.TCPClientData
	isClosed     bool
	managerName  string
}

func GetRemoteAddr(c CustomClient) string {
	return c.conn.RemoteAddr().String()
}

func NewCustomClient(name string, conn net.Conn) *CustomClient {
	client := &CustomClient{tag: TAGGER{TAG: "Custom----->>>Client\n"}, conn: conn, managerName: name}
	client.initSetting()
	return client
}

func (c *CustomClient) initSetting() {
	c.writerChanel = make(chan *net_struct.TCPClientData)
	c.isClosed = false
}
func (c *CustomClient) SendMessage(data *net_struct.TCPClientData) error {
	fmt.Println("发送", c.isClosed)
	if c.isClosed == true {
		return errors.New("管道已关闭")
	}
	c.writerChanel <- data

	return nil
}
func (c *CustomClient) OnDispose() {
	defer c.conn.Close()
	fmt.Println("断开连接", c.conn.RemoteAddr().String())
}

func (c *CustomClient) OnConnect() {
	ip := c.conn.RemoteAddr().String()

	defer func() {
		fmt.Println(fmt.Sprintf("disconnect: %s", ip))
		close(c.writerChanel)
		c.conn.Close()
	}()
	c.wg.Add(2)
	go c.GOReader(c.conn)
	go c.GOWriter(c.conn)
	c.wg.Wait()
}

// 读取携程
func (c *CustomClient) GOReader(conn net.Conn) {
	defer c.wg.Done()
	ip := conn.RemoteAddr().String()
	reader := bufio.NewReader(conn)
	for {
		headerBuf := make([]byte, net_struct.ClientClientHeaderLength)
		fmt.Println("读取数据---------->>>>", ip)
		readLen, err := reader.Read(headerBuf)
		if err != nil || err == io.EOF {
			log.Println("read error ", err)
			eventManager.GetEventManagerForName(res.CONNECT_MGR_Name).
				Call(res.CONTENT_NAME_EVENT_CLIENT_READ_ERROR, ip)
			break
		}
		var buffer = bytes.NewBuffer(headerBuf)
		var header net_struct.TCPClientHeader
		err = binary.Read(buffer, binary.LittleEndian, &header)

		if err != nil {
			log.Println("read buffer error ", err)
			break
		}

		fmt.Println("read data :", header)
		dataBuffer := make([]byte, header.Length-net_struct.ClientClientHeaderLength)
		readLen, err = reader.Read(dataBuffer)

		if err != nil || err == io.EOF {
			log.Println("read error ", err)
			break
		}

		fmt.Println("read data :", dataBuffer)
		fmt.Println(fmt.Sprintf(" read success length : %d, msg : %v", readLen, dataBuffer))

		GetManagerForName(c.managerName).receiveMsg(c.conn.RemoteAddr().String(), *net_struct.NewTCPClientData(header, dataBuffer))
	}
}

func (c *CustomClient) GOWriter(conn net.Conn) {
	defer c.wg.Done()
	ip := conn.RemoteAddr().String()
	writer := bufio.NewWriter(conn)

	fmt.Println("开始接收数据", ip)
	for {
		select {

		case tcpClientData, isClose := <-c.writerChanel:
			fmt.Println("开始写入", tcpClientData)
			if !isClose {
				c.isClosed = true
				log.Println("写入管道已关闭------->>>", ip)
				return
			}

			var header net_struct.TCPClientHeader = tcpClientData.Header
			var buffer = bytes.NewBuffer([]byte{})
			err := binary.Write(buffer, binary.LittleEndian, header)
			if err != nil {
				log.Println("writer error ", err)
				return
			}

			err = binary.Write(buffer, binary.LittleEndian, tcpClientData.GetBody())
			if err != nil {
				log.Println("writer error ", err)
				return
			}
			len, err := writer.Write(buffer.Bytes())
			fmt.Println(len, err)
			err = writer.Flush()
			if err != nil {
				log.Println("writer error ", err)
				return
			}
		}
	}
}
