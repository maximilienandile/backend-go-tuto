package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/maximilienandile/backend-go-tuto/internal/cart"
)

type SendInput struct {
	ToAddress   string
	FromAddress string
	HtmlBody    string
	TextBody    string
	Subject     string
}

type Sender interface {
	Send(input SendInput) error
}

type OrderConfirmationEmailInput struct {
	ProductName string
	StoreLink   string
	LogoURL     string
	OrderItems  map[string]cart.Item
}

type Email struct {
	BodyAsHTML string
	BodyAsText string
}

func NewOrderConfirmationEmail(input OrderConfirmationEmailInput) (Email, error) {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: input.ProductName,
			Link: input.StoreLink,
			// Optional product logo
			Logo:      input.LogoURL,
			Copyright: "Copyright",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Customer",
			Intros: []string{
				"Your order has been processed successfully.",
			},
			Table: hermes.Table{
				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignment: map[string]string{
						"Price": "right",
					},
				},
			},
			Actions: []hermes.Action{},
		},
	}
	lines := make([][]hermes.Entry, 0)
	for _, item := range input.OrderItems {
		line := []hermes.Entry{
			{Key: "Item", Value: item.Title},
			{Key: "Description", Value: item.ShortDescription},
			{Key: "Price", Value: item.TotalPriceVATInc.Display},
		}
		lines = append(lines, line)
	}
	email.Body.Table.Data = lines

	output := Email{}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return output, fmt.Errorf("impossible to generate email HTML: %w", err)
	}
	output.BodyAsHTML = emailBody

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		return output, fmt.Errorf("impossible to generate email Text: %w", err)
	}
	output.BodyAsText = emailText

	return output, nil
}
