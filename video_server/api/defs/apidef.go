package defs

//request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

// comments
type Comment struct {
	Id      string
	ViedoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64  //用于检查用户登录是否过期
}
