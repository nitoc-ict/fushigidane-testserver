package main

import "github.com/fushigidane-testserver/router"

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":5000"))
}
