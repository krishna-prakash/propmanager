package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/krishna/rogerapp/auth"
	prisma "github.com/krishna/rogerapp/generated/prisma-client"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: client,
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	router.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	router.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})))

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
