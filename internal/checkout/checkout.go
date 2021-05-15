package checkout

import (
	"github.com/maximilienandile/backend-go-tuto/internal/cart"
	"github.com/maximilienandile/backend-go-tuto/internal/user"
)

// Session represents a checkout session that is initialized
// when the customer wants to pay his order
type Session struct {
	ID        string    `json:"id"`
	CreatedAt string    `json:"createdAt"`
	Provider  string    `json:"provider"`
	Cart      cart.Cart `json:"cart"`
	User      user.User `json:"user"`
}
