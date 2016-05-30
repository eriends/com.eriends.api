package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
	Route {
		"Api",
		"GET",
		"/api",
		Api,
	},
	Route {
		"ApiId",
		"GET",
		"/api/{apiId}",
		ApiId,
	},
	Route {
		"Wechat",
		"GET",
		"/wechat",
		Wechat,
	},
	Route {
		"WechatCallback",
		"GET",
		"/wechat_callback",
		WechatCallback,
	},
	Route {
		"WechatGetUserList",
		"GET",
		"/wechat/users",
		WechatGetUserList,
	},
}

var router mux.Router

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

