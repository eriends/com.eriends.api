package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/gcfg.v1"
)

var cfg Config
var logo string

func main() {
	LoadConf()
	LoadLogo()
	router := CreateRouter()
	Start(router)
}

func Start(router *mux.Router) {
	addr := fmt.Sprintf("%s:%d", cfg.Listen.Ip, cfg.Listen.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	fmt.Println("Listen on %s", addr)
	log.Fatal(server.ListenAndServe())
}

type ListenStruct struct {
	Ip   string
	Port int
}

type Config struct {
	Listen ListenStruct
}

func LoadConf() {
	err := gcfg.ReadFileInto(&cfg, "conf/eriends.ini")
	PanicOnError(err)
}

func LoadLogo() {
	_logo, err := ioutil.ReadFile("conf/logo")
	PanicOnError(err)
	logo = string(_logo)
	fmt.Println(logo)
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

