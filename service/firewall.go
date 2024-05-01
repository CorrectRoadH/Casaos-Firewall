package service

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/CorrectRoadH/CasaOS-Firewall/codegen"
	"routerd.net/go-firewalld"
)

type FirewallService struct {
	client *firewalld.Client
}

func NewFirewallService() *FirewallService {
	// firewallClient, err := firewalld.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// if firewallClient == nil {
	// 	panic("firewallConn is nil")
	// }

	return &FirewallService{}
}
func (s *FirewallService) ReloadRule() error {
	cmd := exec.Command("firewall-cmd", "--reload")
	ouput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("执行命令时发生错误:", string(ouput))
	}
	return nil
}

func (s *FirewallService) OpenPort(port codegen.Port) error {
	cmd := exec.Command("firewall-cmd", "--permanent", "--add-port="+string(port.Port)+"/"+string(port.Protocol))
	fmt.Println("firewall-cmd", "--permanent", "--add-port="+string(port.Port)+"/"+string(port.Protocol))
	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("执行命令时发生错误:", string(output))
		return err
	}
	return s.ReloadRule()
}

func (s *FirewallService) ClosePort(port codegen.Port) error {
	cmd := exec.Command("firewall-cmd", "--permanent", "--remove-port="+string(port.Port)+"/"+string(port.Protocol))
	fmt.Println("firewall-cmd", "--permanent", "--remove-port="+string(port.Port)+"/"+string(port.Protocol))
	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("执行命令时发生错误:", string(output))
		return err
	}
	return s.ReloadRule()
}

func (s *FirewallService) RemoveAllRule() error {
	ports, err := s.GetOpenedPorts()
	if err != nil {
		return err
	}
	for _, port := range ports {
		err := s.ClosePort(port)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *FirewallService) InitRule() error {
	err := s.RemoveAllRule()
	if err != nil {
		return err
	}
	ports := []codegen.Port{
		{Port: "80", Protocol: codegen.Tcp},
		{Port: "22", Protocol: codegen.Tcp},
	}
	for _, port := range ports {
		err := s.OpenPort(port)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FirewallService) GetOpenedPorts() ([]codegen.Port, error) {
	// run shell firewall-cmd --list-ports
	cmd := exec.Command("firewall-cmd", "--list-ports")

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("执行命令时发生错误:", err)
		return []codegen.Port{}, err
	}

	// spilt 80/tcp 443/tcp
	ports := []codegen.Port{}

	// remove /n
	ouput := strings.Replace(string(output), "\n", "", -1)
	fields := strings.FieldsFunc(ouput, func(r rune) bool {
		return r == ' '
	})

	for _, field := range fields {
		subFields := strings.Split(field, "/")
		if len(subFields) != 2 {
			continue
		}
		port := codegen.Port{
			Port:     subFields[0],
			Protocol: codegen.PortProtocol(subFields[1]),
		}
		ports = append(ports, port)
	}

	return ports, nil
}

func (s *FirewallService) Zones() ([]string, error) {
	ctx := context.Background()

	zones, err := s.client.Config().GetZoneNames(ctx)

	if err != nil {
		return []string{}, err
	}
	return zones, nil
}
