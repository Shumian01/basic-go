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

	Ctime time.Time

	//Addr Address
}

//type Address struct {
//}
