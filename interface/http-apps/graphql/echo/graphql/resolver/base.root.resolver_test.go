package resolver

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRootResolver_SetContext(t *testing.T) {
	e := echo.New()
	c := e.AcquireContext()
	r := &RootResolver{}
	r.SetContext(c)
}
