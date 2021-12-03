package main

import (
	"golang-clean-architecture-ent-gqlgen/config"
	"golang-clean-architecture-ent-gqlgen/graph"
	"golang-clean-architecture-ent-gqlgen/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})
	e := echo.New()

	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	e.POST("/query", echo.WrapHandler(srv))
	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}
