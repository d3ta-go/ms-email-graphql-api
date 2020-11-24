package graphqlschema

import (
	"io"
	"net/http"
	"net/http/httptest"
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
