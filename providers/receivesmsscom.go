package providers

import (
	"fmt"
	"number-verifier/util"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ReceiveSMSSCOM main
type ReceiveSMSSCOM struct {
}

// GetNumbers returns numbers
func (p ReceiveSMSSCOM) GetNumbers() (numbers []Number, err error) {
	var doc *goquery.Document
	if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s", p.GetBaseURL())); err != nil {
		return nil, err
	}

	doc.Find("div.number-boxes-item").Each(func(i int, s *goquery.Selection) {

		numberunparsed := s.Find(".number-boxes-itemm-number").First().Text()
		numberunparsed = strings.Replace(numberunparsed, "+", "", -1)
		numberunparsed = strings.Replace(numberunparsed, " ", "", -1)
		numberunparsed = strings.Replace(numberunparsed, "-", "", -1)

		var number Number
		number.Provider = p.GetName()
		number.PhoneNumber = numberunparsed
		number.Country = s.Find(".number-boxes-item-country").First().Text()

		numbers = append(numbers, number)
	})

	return
}

// GetMessages gets messages
func (p ReceiveSMSSCOM) GetMessages(number string) (messages []string, err error) {
	var doc *goquery.Document
	if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s/sms/%s/", p.GetBaseURL(), number)); err != nil {
		return nil, err
	}

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		from := s.Find("td").Eq(1).Text()
		text := s.Find("td").Eq(3).Text()
		ago := s.Find("td").Eq(4).Text()

		messages = append(messages, fmt.Sprintf("%s - %s (%s)", from, text, ago))
	})

	return messages[:5], err
}

// GetName gets provider name
func (p ReceiveSMSSCOM) GetName() string {
	return "receive-smss.com"
}

// GetBaseURL gets baseURL
func (p ReceiveSMSSCOM) GetBaseURL() string {
	return "https://receive-smss.com/"
}

// GetProviders provider
func (p ReceiveSMSSCOM) GetProviders() Providers {
	return p
}
