package helper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/dgrijalva/jwt-go"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// NewConfig new config - helper test
func NewConfig(h *handler.Handler) (*config.Config, *viper.Viper, error) {
	configPath := "../../../../../../conf"

	//init config
	c, v, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	c.IAM.Casbin.ModelPath = "../../../../../../conf/casbin/casbin_rbac_rest_model.conf"

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

	viper.OnConfigChange(func(e fsnotify.Event) {
		c := new(config.Config)
		if err := viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}
		h.SetDefaultConfig(c)
	})

	return c, v, nil
}

// NewHandler new handler - helper test
func NewHandler() *handler.Handler {
	h, _ := handler.NewHandler()
	// init configuration
	_, _, err := NewConfig(h)
	if err != nil {
		panic(err)
	}
	return h
}

// GenerateUserTestToken generate user test token - helper test
func GenerateUserTestToken(h *handler.Handler, t *testing.T) (string, *identity.JWTCustomClaims, error) {
	j, err := identity.NewJWT(h)
	if err != nil {
		return "", nil, err
	}

	claims := identity.JWTCustomClaims{
		ID:          0,
		UUID:        "test-test-test-test-test",
		Username:    "test.d3tago",
		NickName:    "Test User",
		AuthorityID: "group:admin",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,           // signature effective time
			ExpiresAt: time.Now().Unix() + 60*60*24*30*12, // expiration time 12 month
			Issuer:    j.Issuer,
		},
	}

	token, _, err := j.GenerateToken(claims)
	if err != nil {
		return "", nil, err
	}

	return token, &claims, nil
}
