package braintree

type PaymentMethod interface {
	GetCustomerId() string
	GetToken() string
	GetImageURL() string
}
