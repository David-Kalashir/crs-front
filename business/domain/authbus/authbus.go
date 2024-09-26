package authbus

import (
	"fmt"

	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/google/uuid"
)

type Business struct {
	log *logger.Logger
}

func NewBusiness(log *logger.Logger) *Business {
	return &Business{
		log: log,
	}
}

func (b *Business) Login(usr string, password string) error {
	if usr == "admin" && password == "admin" {
		return nil
	}
	return fmt.Errorf("invalid input")
}

func (b *Business) IsAuth(id uuid.UUID) bool {
	//TODO request to DB
	return true
}
