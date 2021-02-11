package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"number-verifier/util"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// ReceiveSMSOnlineInfo main
type ReceiveSMSOnlineInfo struct {
}

// GetNumbers returns numbers
func (p ReceiveSMSOnlineInfo) GetNumbers() (numbers []Number, err error) {
	var (
		rNum = regexp.MustCompile(`([0-9]+)-(.*)`)
	)

	var doc *goquery.Document
	if doc, err = util.GetGoQueryDoc(fmt.Sprintf("%s", p.GetBaseURL())); err != nil {
		return nil, err
	}

	doc.Find("div.Cell").Each(func(i int, s *goquery.Selection) {

		href, _ := s.Find("a").First().Attr("href")

		var number Number
		number.Provider = p.GetName()
		number.PhoneNumber = rNum.FindStringSubmatch(href)[1]
		number.Country, _ = util.ParseCountry(rNum.FindStringSubmatch(href)[2])

		numbers = append(numbers, number)
	})

	return
}

// GetMessages gets messages
func (p ReceiveSMSOnlineInfo) GetMessages(number string) (messages []string, err error) {
	var (
		req    *http.Request
		resp   *http.Response
		result []ReceiveSMSOnlineInfoMessage
	)

	client := util.GetHTTPTransport()

	req, err = http.NewRequest("GET", fmt.Sprintf("%s/get_sms_439704.php?phone=%s", p.GetBaseURL(), number), nil)
	req.Header.Set("Referer", "https://receive-sms-online.info/")

	resp, err = client.Do(req)
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
		if v.Telefon == "xxx" || v.Mesaj == "xxx" {
			continue
		} else {
			messages = append(messages, v.Telefon+" - "+v.Mesaj)
		}
	}

	return
}

// GetName gets provider name
func (p ReceiveSMSOnlineInfo) GetName() string {
	return "receive-sms-online.info"
}

// GetBaseURL gets baseURL
func (p ReceiveSMSOnlineInfo) GetBaseURL() string {
	return "https://receive-sms-online.info/"
}

// GetProviders provider
func (p ReceiveSMSOnlineInfo) GetProviders() Providers {
	return p
}

// ReceiveSMSOnlineInfoMessage message format
type ReceiveSMSOnlineInfoMessage struct {
	Data      string `json:"data"`
	Mesaj     string `json:"mesaj"`
	MesajeID  string `json:"mesaje_id"`
	Telefon   string `json:"telefon"`
	TelefonID string `json:"telefon_id"`
}
