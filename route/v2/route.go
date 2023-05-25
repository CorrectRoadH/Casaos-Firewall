package v2

import (
	codegen "github.com/CorrectRoadH/CasaOS-Firewall/codegen"
)

type Firewall struct{}

func NewFirewall() codegen.ServerInterface {
	return &Firewall{}
}
