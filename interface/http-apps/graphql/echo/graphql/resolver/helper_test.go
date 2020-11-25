package resolver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	ht "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/features/helper_test"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	c.IAM.Casbin.ModelPath = "../../../../../../conf/casbin/casbin_rbac_rest_model.conf"

	return c, v, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	i.Claims.Username = "test.d3tago"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}

func newEchoContext(h *handler.Handler, t *testing.T) (echo.Context, error) {
	req := httptest.NewRequest(http.MethodPost, "/graphql/v1/operations", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextHTML)
	res := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, res)

	token, claims, err := ht.GenerateUserTestToken(h, t)
	if err != nil {
		return ctx, err
	}
	ctx.Set("identity.token.jwt", token)
	ctx.Set("identity.token.jwt.claims", claims)

	return ctx, nil
}

func newRootResolver(t *testing.T) (*RootResolver, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	ctx, err := newEchoContext(h, t)
	if err != nil {
		return nil, nil, err
	}

	r, err := NewRootResolver(h, ctx)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}
