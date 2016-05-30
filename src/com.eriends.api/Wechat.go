package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/base"
	"github.com/chanxuehong/wechat.v2/mp/user"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
)

const (
	ORIID 					= "gh_1b42ae1dabbd"
	APPID 					= "wx0677f787bf067135"
	APPSECRET 			= "384f19db6203cc27de60c0eaf5395119"          
 	TOKEN 					= "blueantelope"
	BASE64AESKEY 		= "pr9P8viyjRM2QncgaCwFactM3WbgSuelSKNarJotq2v"
)

func Wechat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Wechat!")
}

func WechatCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("wechat callback")
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	srv := core.NewServer(ORIID, APPID, TOKEN, BASE64AESKEY, mux, nil)
	srv.ServeHTTP(w, r, nil)
	fmt.Println(r)
}


func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(APPID, APPSECRET, nil)
	wechatClient *core.Client = core.NewClient(accessTokenServer, nil)	
)

func WechatGetUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get User List")
	fmt.Println(wechatClient)
	userList, err := user.List(wechatClient, "")
	PrintError(err)
	fmt.Println("user list: ", userList)
}

func WechatGetServerIp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Server Ip")
	ipList, err := base.GetCallbackIP(wechatClient)
	PrintError(err)
	fmt.Println("server ip list: ", ipList)
}

