package main

import (
	"golang-clean-architecture-ent-gqlgen/config"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/controller"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/datastore"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/graphql"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/router"
	"golang-clean-architecture-ent-gqlgen/pkg/registry"
	"log"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
