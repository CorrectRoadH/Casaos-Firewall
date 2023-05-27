package service

import (
	"github.com/CorrectRoadH/CasaOS-Firewall/codegen/message_bus"
	"github.com/CorrectRoadH/CasaOS-Firewall/pkg/config"
	v2 "github.com/CorrectRoadH/CasaOS-Firewall/service/v2"
	"github.com/IceWhaleTech/CasaOS-Common/external"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var Cache *cache.Cache

var MyService Services

type Services interface {
	// NftablesService() NftablesService
	Gateway() external.ManagementService
	Firewall() *v2.FirewallService
	MessageBus() *message_bus.ClientWithResponses
}

func NewService(db *gorm.DB) Services {
	gatewayManagement, err := external.NewManagementService(config.CommonInfo.RuntimePath)
	if err != nil {
		panic(err)
	}

	// notifySystem := external.NewNotifyService(config.CommonInfo.RuntimePath)
	// sharesService := external.NewShareService(config.CommonInfo.RuntimePath)

	return &store{
		firewall: v2.NewFirewallService(db),
		gateway:  gatewayManagement,
	}
}

type store struct {
	gateway  external.ManagementService
	firewall *v2.FirewallService
}

func (c *store) Gateway() external.ManagementService {
	return c.gateway
}

func (c *store) Firewall() *v2.FirewallService {
	return c.firewall
}

type NftablesService interface {
}

func (c *store) MessageBus() *message_bus.ClientWithResponses {
	client, _ := message_bus.NewClientWithResponses("", func(c *message_bus.Client) error {
		// error will never be returned, as we always want to return a client, even with wrong address,
		// in order to avoid panic.
		//
		// If we don't avoid panic, message bus becomes a hard dependency, which is not what we want.

		messageBusAddress, err := external.GetMessageBusAddress(config.CommonInfo.RuntimePath)
		if err != nil {
			c.Server = "message bus address not found"
			return nil
		}

		c.Server = messageBusAddress
		return nil
	})

	return client
}
