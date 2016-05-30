package main

import (
	"fmt"
	"net/http"
	"github.com/chanxuehong/wechat.v2/mp/core"
)

var TOKEN = "blueantelope"


func Wechat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Wechat!")
	
//	params := r.URL.Query()
//	signature := params.Get("signature")
//	timestamp := params.Get("timestamp")
//	nonce := params.Get("nonce")
//
//	var tmpArr [3]string
//	tmpArr[0] = params.Get("signature")
//	tmpArr[1] = params.Get("timestamp")
//	tmpArr[2] = params.Get("nonce")
//
//	fmt.Println(tmpArr)

	mux := core.NewServeMux()
	srv := core.NewServer("IDgh_1b42ae1dabbd", "wx0677f787bf067135", "blueantelope", "pr9P8viyjRM2QncgaCwFactM3WbgSuelSKNarJotq2v", mux, nil)

	// 在回调 URL 的 Handler 里处理消息(事件)
	http.HandleFunc("/wechat_callback", func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r, nil)
	})

}


