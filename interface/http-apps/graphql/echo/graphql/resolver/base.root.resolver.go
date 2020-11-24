package resolver

import (
	appEmail "github.com/d3ta-go/ddd-mod-email/modules/email/application"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewRootResolver create new RootResolver
func NewRootResolver(h *handler.Handler, c echo.Context) (*RootResolver, error) {
	var err error

	rsv := new(RootResolver)

	rsv.SetHandler(h)
	rsv.context = c

	if rsv.appEmail, err = appEmail.NewEmailApp(h); err != nil {
		return nil, err
	}

	return rsv, nil
}

// RootResolver represent RootResolver
type RootResolver struct {
	features.BaseFeature
	context  echo.Context
	appEmail *appEmail.EmailApp
}

// SetContext set context
func (r *RootResolver) SetContext(c echo.Context) {
	r.context = c
}
