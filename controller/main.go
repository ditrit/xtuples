package main

import (
	"context"
	"go-http/conf"
	"go-http/pkg/app"
	"log"

	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	err := supertokens.Init(conf.SuperTokensConfig)
	if err != nil {
		panic(err.Error())
	}

	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer app.Exit(context.Background())

	NewServer(app).Start()
}
