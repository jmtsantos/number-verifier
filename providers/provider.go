package providers

import (
	"fmt"
)

// SupportedProviders list of supported providers
var SupportedProviders = []Providers{
	SMSReceiveFree{},
	Upmasked{},
	ReceiveSMSOnlineInfo{},
	SMSOnlineCO{},
	ReceiveSMSSCOM{},
}

// Providers is the interface where all providers need to confirm to.
type Providers interface {
	GetProviders() Providers
	GetNumbers() ([]Number, error)
	GetMessages(number string) ([]string, error)
	GetName() string
	GetBaseURL() string
}

// Provider is a struct which contains some properties about a provider.
type Provider struct {
	Name    string
	BaseURL string
}

// Number is a struct that contains the properties of a phone number
type Number struct {
	PhoneNumber string
	Country     string
	Provider    string
}

func (n Number) String() string {
	return fmt.Sprintf("(%s) [%s] %s", n.Provider, n.Country, n.PhoneNumber)
}
