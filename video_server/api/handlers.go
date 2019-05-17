package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 通过 ByName方法拿到 user_name字段的内容.
	uname := p.ByName("user_name")

	// 通过io流 将 uname发送出去.
	io.WriteString(w, uname)
}
