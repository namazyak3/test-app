package main

import (
	"github.com/namzyak3/test-app/server/connector"
	"github.com/namzyak3/test-app/server/libs/envloader"
)

func main() {
	env := envloader.GetEnv()

	connector.ConnectDB(env)
}
