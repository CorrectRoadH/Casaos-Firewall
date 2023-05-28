package v2

import (
	"net/http"

	"github.com/CorrectRoadH/CasaOS-Firewall/service"

	codegen "github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"github.com/labstack/echo/v4"
)

func (s *Firewall) GetRules(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, codegen.GetRulesResponseOK{})
}

func (s *Firewall) GetVersion(ctx echo.Context) error {
	version := service.MyService.Firewall().GetVersion()
	return ctx.JSON(http.StatusOK, codegen.GetVersionResponseOK{Data: &version})
}
