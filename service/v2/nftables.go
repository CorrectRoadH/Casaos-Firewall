package v2

import (
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

func (s *FirewallService) ExecGetNftablesShell() string {
	result, err := command2.ExecResultStr("source " + config.AppInfo.ShellPath + "/firewall-helper.sh ;GetNftablesVersion")
	if err != nil {
		logger.Error("error when executing shell script to get nftables version", zap.Error(err))
	}
	return result
}
