# Makefile
#	05/27 2016

ROOT_DIR =$(shell cd "$(dirname "$0")"; pwd)
NAME =com.eriends.api
BIN_DIR =${ROOT_DIR/bin
GO =go

.PHONY: clean init

build: clean init compile
	echo "============= finish build ============="
	
compile:
	$(GO) install ${NAME}

run:
	cd ${ROOT_DIR}; ./bin/${NAME}
	
init:
	export GOPATH=${ROOT_DIR}
	$(GO) get github.com/gorilla/mux
	$(GO) get gopkg.in/gcfg.v1
	$(GO) get github.com/chanxuehong/rand
	$(GO) get github.com/chanxuehong/session
	$(GO) get github.com/chanxuehong/sid
	$(GO) get github.com/chanxuehong/wechat.v2/mp/base
	$(GO) get github.com/chanxuehong/wechat.v2/mp/core
	$(GO) get github.com/chanxuehong/wechat.v2/mp/menu
	$(GO) get github.com/chanxuehong/wechat.v2/mp/message/callback/request
	$(GO) get github.com/chanxuehong/wechat.v2/mp/message/callback/response
	$(GO) get github.com/chanxuehong/wechat.v2/mp/user
	$(GO) get github.com/chanxuehong/wechat.v2/mp/oauth2

celan:
	$(RM) -r ${BIN_DIR}

