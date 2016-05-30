package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route {
		"Index",
		[]string{"GET"},
		"/",
		Index,
	},
	Route {
		"Api",
		[]string{"GET"},
		"/api",
		Api,
	},
	Route {
		"ApiId",
		[]string{"GET"},
		"/api/{apiId}",
		ApiId,
	},
	Route {
		"Wechat",
		[]string{"GET"},
		"/wechat",
		Wechat,
	},
	Route {
		"WechatCallback",
		[]string{"GET", "POST"},
		"/wechat/callback",
		WechatCallback,
	},
	Route {
		"WechatGetUserList",
		[]string{"GET"},
		"/wechat/users",
		WechatGetUserList,
	},
	Route {
		"WechatGetServerIp",
		[]string{"GET"},
		"/wechat/server/ips",
		WechatGetServerIp,
	},
}

var router mux.Router

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

