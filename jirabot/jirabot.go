package jirabot

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Jirabot struct {
	baseurl   string
	basicauth string
}

func (b *Jirabot) Getissue(issuenumber string) []byte {
	url1 := "https://jira-eng-gpk2.cisco.com/jira/rest/api/2/issue/" + issuenumber
	auth := "Basic " + b.basicauth
	fmt.Println(url1)
	fmt.Println(auth)
	req, _ := http.NewRequest("GET", url1, nil)

	req.Header.Add("Authorization", auth)

	//proxyUrl, err := url.Parse("http://127.0.0.1:8080")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//	Proxy:           http.ProxyURL(proxyUrl),
	}
	client := http.Client{
		Transport: tr,
		CheckRedirect: func(redirRequest *http.Request, via []*http.Request) error {
			redirRequest.Header = req.Header
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	return body

}

func (b *Jirabot) Queryissue() []byte {
	//	/rest/api/2/search?jql=project="Your Project Key"  https://community.atlassian.com/t5/Jira-questions/REST-Client-get-all-issues-of-project/qaq-p/494826
	// https://jira-eng-gpk2.cisco.com/jira/rest/api/2/search?jql=project="WEBEX"
	url1 := "https://jira-eng-gpk2.cisco.com/jira/rest/api/2/search?jql=project=WEBEX"
	auth := "Basic " + b.basicauth
	fmt.Println(url1)
	fmt.Println(auth)
	req, _ := http.NewRequest("GET", url1, nil)

	req.Header.Add("Authorization", auth)

	//proxyUrl, err := url.Parse("http://127.0.0.1:8080")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//	Proxy:           http.ProxyURL(proxyUrl),
	}
	client := http.Client{
		Transport: tr,
		CheckRedirect: func(redirRequest *http.Request, via []*http.Request) error {
			redirRequest.Header = req.Header
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	return body
}

func Newbot(url string, auth string) *Jirabot {
	return &Jirabot{baseurl: url, basicauth: auth}
}
