package CustomUser

import "time"

type User struct {
	UUID string
	Name string
	Passwd string
	CreateTime time.Time
}