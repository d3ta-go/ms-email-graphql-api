package graphqlschema

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	ht "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/features/helper_test"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newFGraphQLSchema(t *testing.T) (*handler.Handler, *FGraphQLSchema, error) {
	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return nil, nil, err
	}

	// test feature
	graphql, err := NewFGraphQLSchema(handler, "../../graphql/schema/schema.graphql")
	if err != nil {
		t.Errorf("NewFGraphQLSchema: %s", err.Error())
		return nil, nil, err
	}

	return handler, graphql, err
}

func TestGraphiQL(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	// html template
	tpl := &Template{
		templates: template.Must(template.ParseGlob("../../../../../../www/templates/**/*.*ml")),
	}
	e.Renderer = tpl

	req := httptest.NewRequest(http.MethodGet, "/graphql/graphiql", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	_, graphql, err := newFGraphQLSchema(t)
	if err != nil {
		t.Errorf("newFGraphQLSchema: %s", err.Error())
		return
	}

	// Assertions
	if assert.NoError(t, graphql.GraphiQL(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.graphql.GraphiQL: %s", res.Body.String())
	}
}

func TestPlayground(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	// html template
	tpl := &Template{
		templates: template.Must(template.ParseGlob("../../../../../../www/templates/**/*.*ml")),
	}
	e.Renderer = tpl

	req := httptest.NewRequest(http.MethodGet, "/graphql/playground", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	e.NewContext(req, res)

	_, graphql, err := newFGraphQLSchema(t)
	if err != nil {
		t.Errorf("newFGraphQLSchema: %s", err.Error())
		return
	}

	// Assertions
	if assert.NotNil(t, graphql.Playground()) {
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.graphql.Playground: %s", res.Body.String())
	}
}

func TestOperations(t *testing.T) {
	// variables

	// Setup
	e := echo.New()

	// html template
	tpl := &Template{
		templates: template.Must(template.ParseGlob("../../../../../../www/templates/**/*.*ml")),
	}
	e.Renderer = tpl

	reqQraphQL := `{
    "operationName": "IntrospectionQuery",
    "variables": {},
    "query": "\nquery IntrospectionQuery {\n  __schema {\n    queryType {\n      name\n    }\n    mutationType {\n      name\n    }\n    subscriptionType {\n      name\n    }\n    types {\n      ...FullType\n    }\n    directives {\n      name\n      description\n      locations\n      args {\n        ...InputValue\n      }\n    }\n  }\n}\nfragment FullType on __Type {\n  kind\n  name\n  description\n  fields(includeDeprecated: true) {\n    name\n    description\n    args {\n      ...InputValue\n    }\n    type {\n      ...TypeRef\n    }\n    isDeprecated\n    deprecationReason\n  }\n  inputFields {\n    ...InputValue\n  }\n  interfaces {\n    ...TypeRef\n  }\n  enumValues(includeDeprecated: true) {\n    name\n    description\n    isDeprecated\n    deprecationReason\n  }\n  possibleTypes {\n    ...TypeRef\n  }\n}\nfragment InputValue on __InputValue {\n  name\n  description\n  type {\n    ...TypeRef\n  }\n  defaultValue\n}\nfragment TypeRef on __Type {\n  kind\n  name\n  ofType {\n    kind\n    name\n    ofType {\n      kind\n      name\n      ofType {\n        kind\n        name\n        ofType {\n          kind\n          name\n          ofType {\n            kind\n            name\n            ofType {\n              kind\n              name\n              ofType {\n                kind\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"
}`

	req := httptest.NewRequest(http.MethodPost, "/graphql/v1/operations", strings.NewReader(reqQraphQL))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	_, graphql, err := newFGraphQLSchema(t)
	if err != nil {
		t.Errorf("newFGraphQLSchema: %s", err.Error())
		return
	}

	// Assertions
	rH := graphql.Operations(c)
	if assert.NotNil(t, rH) {
		rH.ServeHTTP(c.Response(), c.Request())
		assert.Equal(t, http.StatusOK, res.Code)
		// t.Logf("RESPONSE.graphql.Operations: %s", res.Body.String())
	}
}
