package main

import (
  "apiserver"
  "handlers"
)

var routes = apiserver.Routes {
	apiserver.Route {
			"Index",
			"GET",
			"/",
			handlers.Index,
	},
	apiserver.Route {
		"AuthCallback",
		"GET",
		"/auth/callback",
		handlers.AuthCallback,

	},
	apiserver.Route {
		"Auth",
		"GET",
		"/auth/login",
		handlers.Login,

	},

}