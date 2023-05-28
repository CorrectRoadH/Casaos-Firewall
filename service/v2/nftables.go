package v2

import (
	"strings"

	"github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"github.com/CorrectRoadH/CasaOS-Firewall/pkg/config"
	command2 "github.com/CorrectRoadH/CasaOS-Firewall/pkg/utils/command"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"go.uber.org/zap"
)

func (s *FirewallService) GetRules() string {
	return "Hello, World!"
}

func (s *FirewallService) GetVersion() string {
	return s.ExecGetNftablesShell()
}

func (s *FirewallService) GetOpenedPorts() []codegen.Port {
	ports_string := strings.Split(s.ExecGetOpenedShell(), " ")
	var ports []codegen.Port
	state := "open"
	for _, port := range ports_string {
		ports_array := strings.Split(port, "/")
		if len(ports_array) == 2 {
			ports = append(ports, codegen.Port{
				Port:     &ports_array[0],
				Protocol: &ports_array[1],
				Action:   &state,
			})
		}
	}
	return ports
}

func (s *FirewallService) ExecGetNftablesShell() string {
	result, err := command2.ExecResultStr("source " + config.AppInfo.ShellPath + "/firewall-helper.sh ;GetNftablesVersion")
	if err != nil {
		logger.Error("error when executing shell script to get nftables version", zap.Error(err))
	}
	return result
}

func (s *FirewallService) OpenOrClosePort(Port *string, Action *string, Protocol *string) error {
	if strings.Compare(*Action, "open") == 0 {
		if _, err := command2.OnlyExec("source " + config.AppInfo.ShellPath + "/firewall-helper.sh ;OpenPort " + *Port + " " + *Protocol); err != nil {
			logger.Error("error when executing shell script to set nftables rules", zap.Error(err))
		}

	} else {
		if _, err := command2.OnlyExec("source " + config.AppInfo.ShellPath + "/firewall-helper.sh ;ClosePort " + *Port + " " + *Protocol); err != nil {
			logger.Error("error when executing shell script to set nftables rules", zap.Error(err))
		}
	}
	return nil
}

func (s *FirewallService) ExecGetOpenedShell() string {
	result, err := command2.ExecResultStr("source " + config.AppInfo.ShellPath + "/firewall-helper.sh ;GetFirewallOpenedPorts")
	if err != nil {
		logger.Error("error when executing shell script to get opened ports", zap.Error(err))
	}
	return result
}
