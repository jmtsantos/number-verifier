package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"number-verifier/util"
	"time"
)

// Upmasked main
type Upmasked struct {
}

// GetNumbers returns numbers
func (p Upmasked) GetNumbers() ([]Number, error) {
	type results []struct {
		Number    string    `json:"number"`
		Country   string    `json:"country"`
		CreatedAt time.Time `json:"created_at"`
	}

	var (
		result  results
		numbers []Number
	)

	resp, err := http.Get(fmt.Sprintf("%s/api/sms/numbers", p.GetBaseURL()))
	if err != nil {
		return numbers, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return numbers, err
	}

	err = json.Unmarshal(contents, &result)
	if err != nil {
		return numbers, err
	}

	for _, v := range result {
		var number Number
		number.PhoneNumber = v.Number
		number.Provider = p.GetName()
		number.Country, _ = util.ParseCountry(v.Country)
		numbers = append(numbers, number)
	}

	return numbers, nil
}

// GetMessages gets messages
func (p Upmasked) GetMessages(number string) ([]string, error) {
	type results []struct {
		Body       string    `json:"body"`
		Originator string    `json:"originator"`
		CreatedAt  time.Time `json:"created_at"`
	}

	var (
		result   results
		messages []string
	)

	resp, err := http.Get(fmt.Sprintf("%s/api/sms/messages/%s", p.GetBaseURL(), number))
	if err != nil {
		return messages, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messages, err
	}

	err = json.Unmarshal(contents, &result)
	if err != nil {
		return messages, err
	}

	for _, v := range result {
		messages = append(messages, v.Originator+" - "+v.Body)

		if len(messages) > 5 {
			break
		}
	}

	return messages, nil
}

// GetName gets provider name
func (p Upmasked) GetName() string {
	return "upmasked.com"
}

// GetBaseURL gets baseURL
func (p Upmasked) GetBaseURL() string {
	return "https://upmasked.com"
}

// GetProviders provider
func (p Upmasked) GetProviders() Providers {
	return p
}
