package v2

import (
	"net/http"

	codegen "github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"github.com/labstack/echo/v4"
)

func (s *Firewall) GetState(ctx echo.Context) error {
	return nil
}

func (s *Firewall) OpenPort(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, codegen.ChangePortResponseOK{})
}

func (s *Firewall) GetOpenedPorts(ctx echo.Context) error {
	return nil
}

func (s *Firewall) DeletePortRule(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, codegen.BaseResponse{})
}

func (s *Firewall) UpdatePortRule(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, codegen.BaseResponse{})
}
