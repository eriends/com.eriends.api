package main

import (
	"fmt"
	"net/http"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/user"
)

const (
	ORIID 					= "IDgh_1b42ae1dabbd"
	APPID 					= "wx0677f787bf067135"
	APPSECRET 			= "384f19db6203cc27de60c0eaf5395119"          
 	TOKEN 					= "blueantelope"
	BASE64AESKEY 		= "pr9P8viyjRM2QncgaCwFactM3WbgSuelSKNarJotq2v"
)

func Wechat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Wechat!")
}

func WechatCallback(w http.ResponseWriter, r *http.Request) {
	mux := core.NewServeMux()
	srv := core.NewServer(ORIID, APPID, TOKEN, BASE64AESKEY, mux, nil)
	srv.ServeHTTP(w, r, nil)
}

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(APPID, APPSECRET, nil)
	wechatClient *core.Client = core.NewClient(accessTokenServer, nil)	
)

func WechatGetUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get User List")
	userList, err := user.List(wechatClient, "")
	PrintError(err)
	fmt.Println("user list: ", userList)
}


