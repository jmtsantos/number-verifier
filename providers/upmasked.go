package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Upmasked struct {
}


func (p Upmasked) GetNumbers() ([]string, error) {
	type results []struct {
		Number    string    `json:"number"`
		Country   string    `json:"country"`
		CreatedAt time.Time `json:"created_at"`
	}

	var (
		result results
		numbers []string
	)

	resp, err := http.Get(fmt.Sprintf("%s/api/sms/numbers", p.GetProvider().BaseURL))
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
		numbers = append(numbers, v.Number)
	}

	return numbers, nil
}

func (p Upmasked) GetMessages(number string) ([]string, error) {
	type results[] struct {
		Body       string    `json:"body"`
		Originator string    `json:"originator"`
		CreatedAt  time.Time `json:"created_at"`
	}

	var (
		result results
		messages []string
	)

	resp, err := http.Get(fmt.Sprintf("%s/api/sms/messages/%s", p.GetProvider().BaseURL, number))
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
		messages = append(messages, v.Originator + " - " + v.Body)

		if len(messages) > 5 {
			break
		}
	}

	return messages, nil
}

func (p Upmasked) GetProvider() Provider {
	return Provider{
		Name: "Upmasked",
		BaseURL: "https://upmasked.com",
	}
}
