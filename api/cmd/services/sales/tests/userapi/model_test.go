package user_test

import (
	"time"

	"github.com/David-Kalashir/crs-front/app/domain/userapp"
	"github.com/David-Kalashir/crs-front/business/domain/userbus"
	"github.com/David-Kalashir/crs-front/business/types/role"
)

func toAppUser(bus userbus.User) userapp.User {
	return userapp.User{
		ID:           bus.ID.String(),
		Name:         bus.Name.String(),
		Email:        bus.Email.Address,
		Roles:        role.ParseToString(bus.Roles),
		PasswordHash: nil, // This field is not marshalled.
		Department:   bus.Department,
		Enabled:      bus.Enabled,
		DateCreated:  bus.DateCreated.Format(time.RFC3339),
		DateUpdated:  bus.DateUpdated.Format(time.RFC3339),
	}
}

func toAppUsers(users []userbus.User) []userapp.User {
	items := make([]userapp.User, len(users))
	for i, usr := range users {
		items[i] = toAppUser(usr)
	}

	return items
}

func toAppUserPtr(bus userbus.User) *userapp.User {
	appUsr := toAppUser(bus)
	return &appUsr
}
