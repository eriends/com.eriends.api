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
	$(GO) get github.com/go-ini/ini
	$(GO) get gopkg.in/gcfg.v1

celan:
	$(RM) -r ${BIN_DIR}

