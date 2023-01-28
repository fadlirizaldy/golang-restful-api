package main

import (
	"log"
	"net/http"
	"os"
	"project_alterra/config"
	"project_alterra/graph"
	"project_alterra/graph/generated"
	"project_alterra/router"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func GraphQLServer(){
	config.InitMigrate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func APIServer(){

	// Initansiasi DB
	config.InitMigrate()

	//Instansiasi server
	e := router.New()

	e.Logger.Fatal(e.Start(":1323"))
}

func main(){
	// os.Args[1] is the first dynamic argument
    arg1 := os.Args[1]

    // use arg1 to decide which packages to call

	// Run GQL Server
    if arg1 == "gqlserver" {
        GraphQLServer()
    }

	// Run API Server
    if arg1 == "apiserver" {
        APIServer()
    }
}
