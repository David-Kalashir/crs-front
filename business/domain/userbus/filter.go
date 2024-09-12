package userbus

import (
	"net/mail"
	"time"

	"github.com/David-Kalashir/crs-front/business/types/name"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID               *uuid.UUID
	Name             *name.Name
	Email            *mail.Address
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
}
