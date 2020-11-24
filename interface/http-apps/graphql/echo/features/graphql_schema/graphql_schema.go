package graphqlschema

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	appEmail "github.com/d3ta-go/ddd-mod-email/modules/email/application"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/schema"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/system/handler"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
)

// NewFGraphQLSchema new FGraphQLSchema
func NewFGraphQLSchema(h *handler.Handler, schemaDir string) (*FGraphQLSchema, error) {
	var err error

	f := new(FGraphQLSchema)
	f.SetHandler(h)

	if f.appEmail, err = appEmail.NewEmailApp(h); err != nil {
		return nil, err
	}

	f.graphqlSchemaStr, err = schema.LoadGraphQLSchema(schemaDir)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// FGraphQLSchema represent GraphQLSchema Feature
type FGraphQLSchema struct {
	features.BaseFeature
	appEmail *appEmail.EmailApp

	graphqlSchemaStr string
	rootResolver     *resolver.RootResolver
}

// GraphiQL generate GraphiQL html page
func (f *FGraphQLSchema) GraphiQL(c echo.Context) error {

	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"htmlTitle":       cfg.Applications.Servers.GraphQLAPI.Name,
		"graphqlEndpoint": "/graphql/v1/operations",
	}

	return c.Render(http.StatusOK, "graphql/graphiql", data)
}

// Playground handler GraphQL Playground
func (f *FGraphQLSchema) Playground() http.HandlerFunc {
	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return nil
	}

	return playground.Handler("GraphQL Playground: "+cfg.Applications.Servers.GraphQLAPI.Name, "/graphql/v1/operations")
}

// Operations handle GraphQL Operations
func (f *FGraphQLSchema) Operations(c echo.Context) *relay.Handler {
	if f.rootResolver == nil {
		rootResolver, err := resolver.NewRootResolver(f.GetHandler(), c)
		f.rootResolver = rootResolver
		if err != nil {
			panic(err)
		}
	} else {
		f.rootResolver.SetContext(c)
	}

	graphqlSchema, err := graphql.ParseSchema(f.graphqlSchemaStr, f.rootResolver)
	if err != nil {
		panic(err)
	}
	graphqlHandler := &relay.Handler{Schema: graphqlSchema}

	return graphqlHandler
}
