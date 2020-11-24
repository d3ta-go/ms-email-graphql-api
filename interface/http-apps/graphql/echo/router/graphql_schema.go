package router

import (
	graphqlschema "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/features/graphql_schema"
	internalMiddleware "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/middleware"
	"github.com/labstack/echo/v4"
)

// SetGraphQLSchema set GraphQLSchema
func SetGraphQLSchema(eg *echo.Group, e *echo.Echo, f *graphqlschema.FGraphQLSchema) error {

	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}
	// OpenApi/swagger-ui
	if cfg.Applications.Servers.GraphQLAPI.Options.DisplaySWAPI {
		eg.GET("/graphiql", f.GraphiQL)
		eg.GET("/playground", echo.WrapHandler(f.Playground()))
		eg.Static("/swapi-graphql-ui/assets", "./www/public/swapi-graphql-ui/assets")

	}

	gg := eg.Group("/v1")
	gg.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	// gg.POST("/operations", echo.WrapHandler(f.Operations(e)))
	gg.POST("/operations", func(c echo.Context) error {
		h := f.Operations(c)
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return nil
}
