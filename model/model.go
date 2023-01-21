package model

type User struct {
	Id       uint64 `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var Users = []User{
	{Id: 1, User: "user1", Password: "12345"},
	{Id: 2, User: "user2", Password: "12345"},
	{Id: 3, User: "user3", Password: "12345"},
	{Id: 4, User: "user4", Password: "12345"},
	{Id: 5, User: "user5", Password: "12345"},
}
