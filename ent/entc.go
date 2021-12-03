//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("Error: failed creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./template"),
	}

	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("Error: failed running ent codegen: %v", err)
	}
}
