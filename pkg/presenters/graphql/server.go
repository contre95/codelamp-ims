package graphql

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/health"
	"codelamp-ims/pkg/presenters/graphql/graph"
	"codelamp-ims/pkg/presenters/graphql/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func MapRoutes(he *health.Service, c *clients.Service) {

	// Pass here the domain usecases to the GraphQL handlers
	resolver := generated.Config{Resolvers: &graph.Resolver{}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

}
