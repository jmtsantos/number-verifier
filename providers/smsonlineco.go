package providers

import (
	"fmt"
	"number-verifier/util"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SMSOnlineCO main
type SMSOnlineCO struct {
}

// GetNumbers returns numbers
func (p SMSOnlineCO) GetNumbers() (numbers []Number, err error) {
	var doc *goquery.Document
	if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s", p.GetBaseURL())); err != nil {
		return nil, err
	}

	doc.Find("div.number-boxes-item").Each(func(i int, s *goquery.Selection) {

		numberunparsed := s.Find(".number-boxes-item-number").First().Text()
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
func (p SMSOnlineCO) GetMessages(number string) (messages []string, err error) {
	var doc *goquery.Document
	if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s/%s", p.GetBaseURL(), number)); err != nil {
		return nil, err
	}

	doc.Find("div.list-item").Each(func(i int, s *goquery.Selection) {
		origin := s.Find("h3.list-item-title").First().Text()
		text := s.Find("div.list-item-content").First().Text()

		messages = append(messages, origin+" - "+text)
	})

	return messages[:5], err
}

// GetName gets provider name
func (p SMSOnlineCO) GetName() string {
	return "sms-online.co"
}

// GetBaseURL gets baseURL
func (p SMSOnlineCO) GetBaseURL() string {
	return "https://sms-online.co/receive-free-sms"
}

// GetProviders provider
func (p SMSOnlineCO) GetProviders() Providers {
	return p
}
