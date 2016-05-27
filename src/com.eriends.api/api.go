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

func main() {
	load_conf()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	logo, err := ioutil.ReadFile("conf/logo")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(logo))

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

func load_conf() {
	var err error
	err = gcfg.ReadFileInto(&cfg, "conf/eriends.ini")
	PanicOnError(err)
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Eriends!")
}
