package router

import (
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features/system"
	"github.com/labstack/echo/v4"
)

// SetSystem set FSystem Router
func SetSystem(eg *echo.Group, f *system.FSystem) {
	eg.GET("/health", f.HealthCheck)
}
