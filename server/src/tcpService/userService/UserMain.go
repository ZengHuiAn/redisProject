package userService

import (
	"fmt"
	"redisProject/src/common"
	"redisProject/src/common/CustomUser"
	"redisProject/src/eventManager"
	"redisProject/src/static/res"
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
	return user.userForUUIDMap[userName] != nil
}

//将用户存到redis或者存到mysql 或其他数据库
func (user *UserMain) SaveUserForSql(uuid string) bool {
	itemUser := user.userForUUIDMap[uuid]
	fmt.Println(fmt.Sprintf("TODO 将用户%s存储到数据库 密码%s  UUID: %s 创建时间 %s  ", itemUser.Name, itemUser.Passwd, itemUser.UUID, itemUser.CreateTime.String()))
	return true
}

func (user *UserMain) UserLogin(from_client string, args interface{}) {
	user.AddUser(from_client, userName, passwd)
}

func (user *UserMain) AddUser(from_client string, userName string, passwd string) {
	if user.ContainsUser(userName) {
		//

	} else {
		//

	}
}

var userMain *UserMain

var (
	onlogin = func(ip string, msgid uint32, args interface{}) {
		userMain.UserLogin(ip)
	}
)

func init() {
	userMain = NewUserMain()
	eventManager.GetEventManagerForName(res.EVENTMGR_PROTOCOL_Name).AddProtoEventAction(res.PROTOCOL_C2S, res.LOGIN_C2S, OnLogin)
}
