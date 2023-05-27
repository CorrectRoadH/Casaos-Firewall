package v2

import (
	"net/http"

	codegen "github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"github.com/labstack/echo/v4"
)

func (s *Firewall) GetRules(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, codegen.GetRulesResponseOK{})
}
