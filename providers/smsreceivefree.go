package providers

import (
	"fmt"
	"number-verifier/util"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	countries = []string{"usa", "canada"}
)

// SMSReceiveFree main
type SMSReceiveFree struct {
}

// GetNumbers returns numbers
func (p SMSReceiveFree) GetNumbers() ([]Number, error) {
	var (
		numbers []Number
		err     error
		r       = regexp.MustCompile(`\+([0-9]+)`)
	)

	for _, country := range countries {

		var doc *goquery.Document
		if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s/%s", p.GetBaseURL(), country)); err != nil {
			return nil, err
		}

		// Cloudflare
		html, _ := doc.Html()
		if strings.Contains(html, "Please complete the security check to access") {
			return nil, fmt.Errorf("got busted by cloudflare, check your cookies")
		}

		doc.Find("a.numbutton").Each(func(i int, s *goquery.Selection) {
			var number Number
			number.PhoneNumber = r.FindStringSubmatch(s.Text())[1]
			number.Country, _ = util.ParseCountry(country)
			number.Provider = p.GetName()
			numbers = append(numbers, number)
		})
	}

	return numbers, err
}

// GetMessages gets messages
func (p SMSReceiveFree) GetMessages(number string) ([]string, error) {
	var (
		messages []string
	)

	doc, err := util.GetGoQueryDoc(fmt.Sprintf("%s/info/%s", p.GetBaseURL(), number))
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`\r?\n`)
	doc.Find(".messagesTable tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		message := ""

		s.Find("td").Each(func(i int, s *goquery.Selection) {
			cleanMessage := re.ReplaceAllString(s.Text(), " ")
			message += strings.TrimSpace(cleanMessage) + " "

			if s.Size()+1 != i {
				message += "- "
			}
		})

		messages = append(messages, message)
		return len(messages) <= 5
	})

	return messages, nil
}

// GetName gets provider name
func (p SMSReceiveFree) GetName() string {
	return "smsreceivefree.com"
}

// GetBaseURL gets baseURL
func (p SMSReceiveFree) GetBaseURL() string {
	return "https://smsreceivefree.com/country"
}

// GetProviders provider
func (p SMSReceiveFree) GetProviders() Providers {
	return p
}
