package main

import (
	"fmt"
	"net/http"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/base"
	"github.com/chanxuehong/wechat.v2/mp/user"
)

const (
	ORIID 					= "IDgh_1b42ae1dabbd"
	APPID 					= "wx0677f787bf067135"
	APPSECRET 			= "384f19db6203cc27de60c0eaf5395119"          
 	TOKEN 					= "blueantelope"
	BASE64AESKEY 		= "pr9P8viyjRM2QncgaCwFactM3WbgSuelSKNarJotq2v"
 	HANDLER 				= mux
 	ERRORHANDLER 		= nil
)

func Wechat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Wechat!")
}

func WechatCallback(w http.ResponseWriter, r *http.Request) {
	mux := core.NewServeMux()
	srv := core.NewServer(ORIID, APPID, TOKEN, BASE64AESKEY, HANDLER, ERRORHANDLER)
	srv.ServeHTTP(w, r, nil)
}

const (
	wxAppId					=  APPID
	wxAppSecret			= "appsecret"
	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(APPDI, APPSECRET, nil)
	wechatClient *core.Client = core.NewClient(accessTokenServer, nil)	
	userList := user.list(wechatClient, "")
	fmt.Println("user total: ", len(userList))
}


