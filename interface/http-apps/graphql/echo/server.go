package echo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func initConfig(h *handler.Handler) (*config.Config, error) {
	//init config
	cfg, viper, err := config.NewConfig("./")
	if err != nil {
		panic(err)
	}
	h.SetDefaultConfig(cfg)

	viper.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("config file changed:", e.Name)
		c := new(config.Config)
		if err := viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}

		h.SetDefaultConfig(c)
		initializeSystems(h)
	})

	return cfg, nil
}

func initializeSystems(h *handler.Handler) error {

	// initialize database
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		panic(err)
	}

	// initialize cacher
	if err := initialize.OpenAllCacheConnection(h); err != nil {
		panic(err)
	}

	return nil
}

// StartGraphQLAPIServer start GraphQL Server
func StartGraphQLAPIServer() error {

	// init super handler
	superHandler := new(handler.Handler)

	// init configuration
	cfg, err := initConfig(superHandler)
	if err != nil {
		// fmt.Errorf("StartGraphQLAPIServer.initConfig: %s", err.Error())
		return err
	}

	// initialize Systems
	err = initializeSystems(superHandler)
	if err != nil {
		// fmt.Errorf("StartGraphQLAPIServer.initializeSystems: %s", err.Error())
		return err
	}
	defer initialize.CloseDBConnections(superHandler)

	// init echo server
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	// e.Debug = true

	// set header - banner
	e.HideBanner = cfg.Applications.Servers.GraphQLAPI.Options.ShowEngineHeader
	if e.HideBanner {
		printSvrHeader(e, cfg)
	}

	// Set routers
	SetRouters(e, superHandler)

	// Start server with Graceful shutdown
	httpPort := fmt.Sprintf(":%s", cfg.Applications.Servers.GraphQLAPI.Options.Listener.Port)
	go func() {
		if err := e.Start(httpPort); err != nil {
			e.Logger.Infof("Shutting down the server [%s]", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
