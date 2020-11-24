package echo

import (
	"github.com/labstack/echo/v4"
	color "github.com/labstack/gommon/color"

	"github.com/d3ta-go/system/system/config"
)

const (
	// https://www.ascii-art-generator.org/ ( lean )
	myBanner = `
Welcome to:

    _/_/_/    _/_/_/    _/_/_/_/_/    _/_/                       
   _/    _/        _/      _/      _/    _/    _/_/_/    _/_/    
  _/    _/    _/_/        _/      _/_/_/_/  _/    _/  _/    _/   
 _/    _/        _/      _/      _/    _/  _/    _/  _/    _/    
_/_/_/    _/_/_/        _/      _/    _/    _/_/_/    _/_/       
                                               _/                
                                          _/_/       %s

High performance, minimalist d3ta-go web framework
Based on Echo %s (https://echo.labstack.com/)

%s
%s
___________________________________________O/____________
                                    	   O\
`
)

// printSvrHeader print server header - banner
func printSvrHeader(e *echo.Echo, cfg *config.Config) {
	colorer := color.New()
	colorer.SetOutput(e.Logger.Output())
	colorer.Printf(myBanner,
		colorer.Cyan("v"+cfg.Applications.Servers.GraphQLAPI.Version),
		colorer.Red("v"+echo.Version),
		colorer.Cyan(cfg.Applications.Servers.GraphQLAPI.Name),
		colorer.Yellow(cfg.Applications.Servers.GraphQLAPI.Description),
	)
}
