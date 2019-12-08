package userService

import (
	"fmt"
	"redisProject/src/common"
	"redisProject/src/common/CustomUser"
	"redisProject/src/eventManager"
	"redisProject/src/net_struct"
	"redisProject/src/pack"
	"redisProject/src/static/res"
	"redisProject/src/tcpService/client"
	"time"
)

type UserMain struct {
	userForUUIDMap map[string]*CustomUser.User
}

func NewUserMain() *UserMain {
	return &UserMain{userForUUIDMap: make(map[string]*CustomUser.User)}
}

func (user *UserMain) CreateUser(userName string, passwd string) *CustomUser.User {
	uuid := common.MD5EncodeString(userName)
	var u = CustomUser.User{UUID: uuid, Name: userName, Passwd: passwd, CreateTime: time.Now()}
	return &u
}

func (user *UserMain) ContainsUser(userName string) bool {
	uuid := common.MD5EncodeString(userName)
	return user.userForUUIDMap[uuid] != nil
}

//将用户存到redis或者存到mysql 或其他数据库
func (user *UserMain) SaveUserForSql(uuid string) bool {
	itemUser := user.userForUUIDMap[uuid]
	fmt.Println(fmt.Sprintf("TODO 将用户%s存储到数据库 密码%s  UUID: %s 创建时间 %s  ", itemUser.Name, itemUser.Passwd, itemUser.UUID, itemUser.CreateTime.String()))
	return true
}

func (user *UserMain) UserLogin(from_client string, data *net_struct.TCPClientData) {
	args := pack.Encode(data.GetBody())
	var bodys = args.([]interface{})

	fmt.Println("用户注册---->>>>", bodys)
	uuid, err := user.AddUser(from_client, bodys[1].(string), bodys[2].(string))

	errcode := ""
	if err != nil {
		errcode = res.LoginErrorCode_Contains.Error()
	}

	bd := []interface{}{
		bodys[0].(int32) + 1,
		errcode,
		uuid,
	}
	response := net_struct.TCPClientData{Header: net_struct.MakeHeader(data.Header.MessageID + 1)}
	response.SetBody(bd)
	response.Header.Length = uint32(net_struct.ClientClientHeaderLength + len(response.GetBody()))
	connectManager := client.GetManagerForName(res.CONNECT_MGR_Name)
	connectManager.ReceiveLocalMsgDone(from_client, response)
}

func (user *UserMain) AddUser(from_client string, userName string, passwd string) (string, error) {
	if user.ContainsUser(userName) {
		//
		return "", res.LoginErrorCode_Contains
	} else {
		u := user.CreateUser(userName, passwd)
		user.userForUUIDMap[u.UUID] = u
		return u.UUID, nil
	}
}

var userMain *UserMain

var (
	onlogin = func(ip string, msgid uint32, data *net_struct.TCPClientData) {
		userMain.UserLogin(ip, data)
	}
)

func init() {
	fmt.Println("初始化用户服务----->>>>")
	userMain = NewUserMain()
	eventManager.GetEventManagerForName(res.EVENTMGR_PROTOCOL_Name).AddProtoEventAction(res.PROTOCOL_C2S, res.LOGIN_C2S, &onlogin)
}
