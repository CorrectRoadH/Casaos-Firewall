package service_test

import (
	"testing"

	"github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"github.com/CorrectRoadH/CasaOS-Firewall/service"
	"gotest.tools/v3/assert"
)

// func TestGetALLRule(t *testing.T) {
// 	firewallService := service.NewFirewallService()
// 	ports, err := firewallService.GetOpenedPorts()
// 	assert.NilError(t, err)
// 	assert.Equal(t, len(ports), 3)
// }

func TestOpenAPort(t *testing.T) {
	firewallService := service.NewFirewallService()
	err := firewallService.RemoveAllRule()
	assert.NilError(t, err)
	ports, err := firewallService.GetOpenedPorts()
	assert.NilError(t, err)
	assert.Equal(t, len(ports), 0)

	err = firewallService.OpenPort(codegen.Port{
		Port:     "80",
		Protocol: codegen.Tcp,
	})
	assert.NilError(t, err)

	err = firewallService.OpenPort(codegen.Port{
		Port:     "80",
		Protocol: codegen.Udp,
	})
	assert.NilError(t, err)

	ports, err = firewallService.GetOpenedPorts()
	assert.NilError(t, err)
	assert.Equal(t, len(ports), 2)

	assert.Equal(t, ports[0].Port, "80")
	assert.Equal(t, ports[0].Protocol, codegen.Tcp)
	assert.Equal(t, ports[1].Port, "80")
	assert.Equal(t, ports[1].Protocol, codegen.Udp)
}

func TesCloseAPort(t *testing.T) {
	firewallService := service.NewFirewallService()
	err := firewallService.RemoveAllRule()
	assert.NilError(t, err)

	err = firewallService.OpenPort(codegen.Port{
		Port:     "80",
		Protocol: codegen.Tcp,
	})
	assert.NilError(t, err)

	err = firewallService.OpenPort(codegen.Port{
		Port:     "80",
		Protocol: codegen.Udp,
	})
	assert.NilError(t, err)

	err = firewallService.ClosePort(codegen.Port{
		Port:     "80",
		Protocol: codegen.Tcp,
	})
	assert.NilError(t, err)

	ports, err := firewallService.GetOpenedPorts()
	assert.NilError(t, err)

	assert.Equal(t, len(ports), 1)
	assert.Equal(t, ports[0].Port, "80")
	assert.Equal(t, ports[0].Protocol, codegen.Udp)
}
