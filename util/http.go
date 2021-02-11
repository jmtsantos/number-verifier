package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// GetGoQueryDoc fetches an URL and parses the document to GoQuery.
func GetGoQueryDoc(link string) (*goquery.Document, error) {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)
	client := GetHTTPTransport()

	var cookies []*http.Cookie
	var cookie *http.Cookie

	cookie = &http.Cookie{
		Name:     "cf_clearance",
		Value:    "28e13d1b3ef951fb5dc18e499ae884157ca369c8-1612956376-0-250",
		Path:     "/",
		Domain:   ".smsreceivefree.com",
		Secure:   true,
		HttpOnly: true,
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:     "__cfduid",
		Value:    "d9a6abd6462e4c9883ad5ceb1f82a423f1612956365",
		Path:     "/",
		Domain:   ".smsreceivefree.com",
		HttpOnly: true,
	}
	cookies = append(cookies, cookie)

	u, _ := url.Parse(link)
	client.Jar.SetCookies(u, cookies)

	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Convert HTML into goquery document
	return goquery.NewDocumentFromReader(resp.Body)
}

// GetHTTPTransport http client setup
func GetHTTPTransport() *http.Client {
	var rtTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 2 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ExpectContinueTimeout: 1 * time.Second,
	}

	cookieJar, _ := cookiejar.New(nil)

	return &http.Client{
		Transport: rtTransport,
		Jar:       cookieJar,
	}
}
