package features

import (
	graphqlschema "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/features/graphql_schema"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features/system"
	"github.com/d3ta-go/system/system/handler"
)

// NewFeatures create new Features
func NewFeatures(h *handler.Handler) (*Features, error) {
	var err error

	f := new(Features)
	f.handler = h

	if f.System, err = system.NewSystem(h); err != nil {
		return nil, err
	}

	if f.GraphQLSchema, err = graphqlschema.NewFGraphQLSchema(h, ""); err != nil {
		return nil, err
	}

	return f, nil
}

// Features represent Features
type Features struct {
	handler *handler.Handler

	System        *system.FSystem
	GraphQLSchema *graphqlschema.FGraphQLSchema
}
