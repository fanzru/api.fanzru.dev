package main

import (
	"api.fanzru.dev/config"
	"api.fanzru.dev/routes"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := routes.Init()

	config := config.GetDatabaseConfig()
	_ = mgm.SetDefaultConfig(nil, "fanzru", options.Client().ApplyURI(config.URL))

	e.Logger.Fatal(e.Start(":5000"))
}

