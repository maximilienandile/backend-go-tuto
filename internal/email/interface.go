package email

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
