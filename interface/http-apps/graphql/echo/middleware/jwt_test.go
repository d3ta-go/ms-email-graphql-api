package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/labstack/echo/v4"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetDefaultConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	return h, nil
}

func TestJWTVerifier(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
		return
	}

	if h != nil {

		type args struct {
			h *handler.Handler
		}
		tests := []struct {
			name string
			args args
			want echo.MiddlewareFunc
		}{
			// TODO: Add test cases.
			{
				name: "Test JWTVerifier(): OK",
				args: args{h: h},
				want: JWTVerifier(h),
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := JWTVerifier(tt.args.h); !reflect.DeepEqual(fmt.Sprintf("%T", got), fmt.Sprintf("%T", tt.want)) {
					t.Errorf("JWTVerifier() = %v, want %v", fmt.Sprintf("%T", got), fmt.Sprintf("%T", tt.want))
				}
			})
		}
	}
}
