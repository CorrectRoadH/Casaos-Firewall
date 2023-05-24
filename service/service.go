package service

import (
	"github.com/CorrectRoadH/CasaOS-Firewall/pkg/config"
	"github.com/IceWhaleTech/CasaOS-Common/external"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var Cache *cache.Cache

var MyService Services

type Services interface {
	// NftablesService() NftablesService
	Gateway() external.ManagementService
}

func NewService(db *gorm.DB) Services {
	gatewayManagement, err := external.NewManagementService(config.CommonInfo.RuntimePath)
	if err != nil {
		panic(err)
	}

	// notifySystem := external.NewNotifyService(config.CommonInfo.RuntimePath)
	// sharesService := external.NewShareService(config.CommonInfo.RuntimePath)

	return &store{
		gateway: gatewayManagement,
	}
}

type store struct {
	gateway external.ManagementService
}

func (c *store) Gateway() external.ManagementService {
	return c.gateway
}

type NftablesService interface {
}
