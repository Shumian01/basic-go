package domain

import "time"

// User 领域对象 是DDD中的entity
type User struct {
	Id       int64
	Email    string
	Password string

	Nickname string
	// YYYY-MM-DD
	Birthday time.Time
	AboutMe  string

	Phone string

	// UTC 0 的时区
	Ctime time.Time

	//Addr Address
}

//type Address struct {
//}
