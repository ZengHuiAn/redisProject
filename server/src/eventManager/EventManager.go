package eventManager

import "redisProject/src/common"

type Event struct {
	Action func(args interface{})
}

type EventManager struct {
	eventPackage map[string][]*Event
	//addr msgID
	protoEventPackage map[common.ProtoMsgHead][]*Event
}

func (manager *EventManager) initSetting() {
	manager.eventPackage = make(map[string][]*Event)
	manager.protoEventPackage = make(map[common.ProtoMsgHead][]*Event)
}

func (manager *EventManager) AddEventAction(eventName string, callback *Event) {
	if manager.eventPackage[eventName] == nil {
		manager.eventPackage [eventName] = make([]*Event, 1)
	}
	manager.eventPackage [eventName] = append(manager.eventPackage [eventName], callback)
}

//协议事件
func (manager *EventManager) AddProtoEventAction(eventName string, msgID uint32, callback *Event) {
	searchProto := common.NewProtoMsgHead(eventName, msgID)
	if manager.protoEventPackage[searchProto] == nil {
		manager.protoEventPackage [searchProto] = make([]*Event, 1)
	}
	manager.protoEventPackage [searchProto] = append(manager.protoEventPackage [searchProto], callback)
}

func (manager *EventManager) RemoveProtoEventAction(eventName string, msgID uint32, callback *Event) {
	searchProto := common.NewProtoMsgHead(eventName, msgID)
	if manager.protoEventPackage[searchProto] == nil {
		manager.protoEventPackage [searchProto] = make([]*Event, 1)
	}
	for i := 0; i < len(manager.protoEventPackage [searchProto]); i++ {
		ele := manager.protoEventPackage [searchProto][i]

		if ele == callback {
			manager.protoEventPackage [searchProto] = append(manager.protoEventPackage [searchProto][:i], manager.protoEventPackage [searchProto][i+1:]...)
			break
		}
	}
}

func (manager *EventManager) CallProto(eventName string, msgID uint32, data interface{}) {
	searchProto := common.NewProtoMsgHead(eventName, msgID)
	v := manager.protoEventPackage[searchProto]
	if len(v) > 0 {
		for i := 0; i < len(v); i++ {
			callBack := v[i]
			if callBack != nil {
				callBack.Action(data)
			}
		}
	}
}

func (manager *EventManager) RemoveEventAction(eventName string, callback *Event) {
	if manager.eventPackage[eventName] == nil {
		manager.eventPackage [eventName] = make([]*Event, 1)
	}
	for i := 0; i < len(manager.eventPackage [eventName]); i++ {
		ele := manager.eventPackage [eventName][i]

		if ele == callback {
			manager.eventPackage [eventName] = append(manager.eventPackage [eventName][:i], manager.eventPackage [eventName][i+1:]...)
			break
		}
	}
}

func (manager *EventManager) Call(eventName string, data interface{}) {
	v := manager.eventPackage[eventName]
	if len(v) > 0 {
		for i := 0; i < len(v); i++ {
			callBack := v[i]
			if callBack != nil {
				callBack.Action(data)
			}
		}
	}
}

func MakeEventManager() *EventManager {
	manager := &EventManager{}
	manager.initSetting()
	return manager
}

var eventManagerInstance map[string]*EventManager

func GetEventManagerForName(name string) *EventManager {
	if _, ok := eventManagerInstance[name]; ok == false {
		eventManagerInstance[name] = MakeEventManager()
	}
	return eventManagerInstance[name]
}

func init() {
	eventManagerInstance = make(map[string]*EventManager)
}
