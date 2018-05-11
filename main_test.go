package main

import (
	"encoding/json"
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
	"github.com/boyzhujian/gociscospark-jira/ciscosparkbot"
	//"github.com/boyzhujian/gociscospark-jira/jirabot"
)

var conftest = new(config)

func init() {
	if _, err := toml.DecodeFile("config.toml", &conftest); err != nil {
		log.Fatal(err)
	}
}

// func Test_botsendmessage(t *testing.T) {
// 	b := ciscosparkbot.Newbot(conftest.Accesstoken)
// 	roomid := conftest.Jirawebhookroom
// 	b.Sendmessage(roomid, "test message from  cisospark_jira")

// }

func Test_botGetmessage(t *testing.T) {
	b := ciscosparkbot.Newbot(conftest.Accesstoken)
	fmt.Println(b)
	fmt.Println("test get message")
	gotmes := b.Getmessage("Y2lzY29zcGFyazovL3VzL01FU1NBR0UvYzk2YjI3ZDAtNTQyYi0xMWU4LTgyMDItY2Q4OGE4ZDNhYjQ3")
	m := new(Onesparkmessage)
	json.Unmarshal([]byte(gotmes), &m)
	fmt.Println(m.Text)

}

// func Test_jiragetoneissue(t *testing.T) {
// 	//https://jira-eng-gpk2.cisco.com/jira/rest/api/2/issue/WEBEX-350
// 	j := jirabot.Newbot("https://jira-eng-gpk2.cisco.com/jira/rest/api/2", conf.Jiraauth)
// 	outjsonbyte := j.Getissue("WEBEX-350")
// 	log.Println(string(outjsonbyte))

// }

// func Test_JiraQueryissue(t *testing.T) {
// 	//https://jira-eng-gpk2.cisco.com/jira/rest/api/2/issue/WEBEX-350
// 	j := jirabot.Newbot("https://jira-eng-gpk2.cisco.com/jira/rest/api/2", conf.Jiraauth)
// 	outjsonbyte := j.Queryissue()
// 	log.Println(string(outjsonbyte))

// }
