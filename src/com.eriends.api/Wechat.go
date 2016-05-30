package main

import (
	"fmt"
	"net/http"
	"github.com/chanxuehong/wechat.v2/mp/core"
)

var TOKEN = "blueantelope"


func Wechat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Wechat!")
}

func WechatCallback(w http.ResponseWriter, r *http.Request) {
	mux := core.NewServeMux()
	srv := core.NewServer("IDgh_1b42ae1dabbd", "wx0677f787bf067135", "blueantelope", "pr9P8viyjRM2QncgaCwFactM3WbgSuelSKNarJotq2v", mux, nil)
	srv.ServeHTTP(w, r, nil)
}


