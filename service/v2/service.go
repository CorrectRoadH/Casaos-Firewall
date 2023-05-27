package v2

import "gorm.io/gorm"

type FirewallService struct {
	_db *gorm.DB
}

func NewFirewallService(db *gorm.DB) *FirewallService {
	return &FirewallService{
		_db: db,
	}
}
