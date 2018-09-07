package main

import (
	"github.com/FernandoCagale/serverless-auth/api/routers"
	"github.com/urfave/negroni"
)

func main() {
	router := routers.NewRouter()

	routers.MakeHandlers(router)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run("127.0.0.1:3000")
}
