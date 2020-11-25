package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/utils"
	"github.com/labstack/echo/v4"
)

// JWTVerifier verify JWT token from internal Identity Provider
func JWTVerifier(h *handler.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			// get token
			authHeaderPart := strings.Split(authHeader, " ")
			token := authHeaderPart[len(authHeaderPart)-1]

			// jwt tools
			j, err := identity.NewJWT(h)
			if err != nil {
				return response.FailWithMessageWithCode(http.StatusInternalServerError, err.Error(), c)
			}

			// if token is empty, create anonymous token
			var isAnonymous bool
			if token == "" {
				isAnonymous = true
				_, t, _, err := j.GenerateAnonymousToken()
				if err != nil {
					return response.FailWithMessageWithCode(http.StatusInternalServerError, err.Error(), c)
				}
				token = t
			}

			// parseToken parses the information contained in the token
			claims, err := j.ParseToken(token)
			if err != nil {
				data := map[string]interface{}{
					"reload": true,
				}
				if err == identity.ErrTokenExpired {
					return response.FailWithDetailed(response.ERROR, data, "Authorization has expired", c)
				}
				return response.FailWithDetailed(response.ERROR, data, err.Error(), c)
			}

			// verify token to persistent storage
			if isAnonymous == false {
				exist, err := isSessionExist(token, h)
				if err != nil {
					return response.FailWithMessageWithCode(http.StatusInternalServerError, err.Error(), c)
				}
				if !exist {
					return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider (EA2M) Error [%s]", "Invalid Token"), c)
				}
			}

			// if expired
			now := time.Now().Unix()
			if claims.ExpiresAt < now {
				return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider (EA2M) Error [%s]", "Token Expired"), c)
			}

			// if not valid before
			if claims.NotBefore > now {
				return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider (EA2M) Error [%s]", "Token Not Valid Berofe"), c)
			}
			c.Set("identity.token.jwt", token)
			c.Set("identity.token.jwt.claims", claims)

			return next(c)
		}
	}
}

func isSessionExist(sessionValue string, h *handler.Handler) (bool, error) {
	cfg, err := h.GetDefaultConfig()
	if err != nil {
		return false, err
	}
	// goreportcard - ineffassign report warning: ineffectual assignment to can (ineffassign)
	// ce, err := h.GetCacher(cfg.Caches.SessionCache.ConnectionName)
	ce, err := h.GetCacher(cfg.Caches.SessionCache.ConnectionName)
	if err != nil {
		return false, err
	}
	ce.Context = "interface"
	ce.Container = "session"
	ce.Component = "jwt"

	sessionKey := utils.MD5([]byte(sessionValue))

	return ce.IsExist(sessionKey), nil
}
