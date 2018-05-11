package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/boyzhujian/gociscospark-jira/ciscosparkbot"
	"github.com/boyzhujian/gociscospark-jira/jirabot"
	"github.com/gorilla/mux"
)

type config struct {
	Sparkwebhookurl string
	Accesstoken     string
	Jirawebhookroom string
	Jiraauth        string
}

//Sparkmessage is the first level sparkwebhook api return
type Sparkmessage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TargetURL string    `json:"targetUrl"`
	Resource  string    `json:"resource"`
	Event     string    `json:"event"`
	OrgID     string    `json:"orgId"`
	CreatedBy string    `json:"createdBy"`
	AppID     string    `json:"appId"`
	OwnedBy   string    `json:"ownedBy"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	ActorID   string    `json:"actorId"`
	Data      struct {
		ID          string    `json:"id"`
		RoomID      string    `json:"roomId"`
		RoomType    string    `json:"roomType"`
		PersonID    string    `json:"personId"`
		PersonEmail string    `json:"personEmail"`
		Created     time.Time `json:"created"`
	} `json:"data"`
}

// Onesparkmessage use the DataId of the sparkmessage to retrive the real message content
type Onesparkmessage struct {
	ID          string    `json:"id"`
	RoomID      string    `json:"roomId"`
	RoomType    string    `json:"roomType"`
	Text        string    `json:"text"`
	PersonID    string    `json:"personId"`
	PersonEmail string    `json:"personEmail"`
	Created     time.Time `json:"created"`
}

var conf = new(config)

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("helloworld")
	//bot := ciscosparkbot.Newbot(conf.Accesstoken)
	//bot.Registermessagewebhook()
	runwebserver()

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use ciscospark to query jira issue")

}

func jirawebhookhandler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	m := new(Sparkmessage)
	err := json.Unmarshal(s, &m)
	if err != nil {
		return
	}
	b := ciscosparkbot.Newbot(conf.Accesstoken)
	gotmes := b.Getmessage(m.Data.ID)
	mes := new(Onesparkmessage)
	json.Unmarshal([]byte(gotmes), &mes)
	fmt.Println(mes.Text)

	mm := mes.Text
	if strings.HasPrefix(mm, "issue") {
		mmm := strings.Replace(mm, "issue", "", -1)
		fmt.Println(mmm)
		mmm = strings.Join(strings.Fields(mmm), "")
		j := jirabot.Newbot("https://jira-eng-gpk2.cisco.com/jira/rest/api/2", conf.Jiraauth)
		outjsonbyte := j.Getissue(mmm)
		fmt.Println(outjsonbyte)
		jsonJirames := &Jiramessage{}
		err := json.Unmarshal(outjsonbyte, &jsonJirames)
		if err != nil {
			return
		}
		b.Sendmessage(conf.Jirawebhookroom, jsonJirames.Fields.Description)

	}

}

func runwebserver() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/jirawebhook", jirawebhookhandler)

	s := &http.Server{
		Addr:           ":6080",
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type Jiramessage struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		FixVersions      []interface{} `json:"fixVersions"`
		Customfield10473 struct {
			Self  string `json:"self"`
			Value string `json:"value"`
			ID    string `json:"id"`
		} `json:"customfield_10473"`
		Customfield11443 []struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"customfield_11443"`
		Resolution struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"resolution"`
		Customfield11442 interface{} `json:"customfield_11442"`
		Customfield11201 interface{} `json:"customfield_11201"`
		Customfield11720 interface{} `json:"customfield_11720"`
		Customfield10500 struct {
			Self  string `json:"self"`
			Value string `json:"value"`
			ID    string `json:"id"`
		} `json:"customfield_10500"`
		Customfield10469 interface{} `json:"customfield_10469"`
		Customfield10505 interface{} `json:"customfield_10505"`
		Customfield11717 interface{} `json:"customfield_11717"`
		Customfield10506 interface{} `json:"customfield_10506"`
		Customfield10507 interface{} `json:"customfield_10507"`
		Customfield10508 interface{} `json:"customfield_10508"`
		Customfield11718 interface{} `json:"customfield_11718"`
		LastViewed       string      `json:"lastViewed"`
		Customfield10100 interface{} `json:"customfield_10100"`
		Priority         struct {
			Self    string `json:"self"`
			IconURL string `json:"iconUrl"`
			Name    string `json:"name"`
			ID      string `json:"id"`
		} `json:"priority"`
		Customfield10102 string        `json:"customfield_10102"`
		Labels           []interface{} `json:"labels"`
		Customfield10466 []struct {
			Self  string `json:"self"`
			Value string `json:"value"`
			ID    string `json:"id"`
		} `json:"customfield_10466"`
		Customfield11700              interface{}   `json:"customfield_11700"`
		Customfield11425              interface{}   `json:"customfield_11425"`
		Customfield11424              interface{}   `json:"customfield_11424"`
		Customfield11702              interface{}   `json:"customfield_11702"`
		Customfield11701              interface{}   `json:"customfield_11701"`
		Customfield11426              interface{}   `json:"customfield_11426"`
		Customfield11704              interface{}   `json:"customfield_11704"`
		Customfield11307              interface{}   `json:"customfield_11307"`
		Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
		Timeestimate                  interface{}   `json:"timeestimate"`
		Versions                      []interface{} `json:"versions"`
		Customfield11428              interface{}   `json:"customfield_11428"`
		Customfield11703              interface{}   `json:"customfield_11703"`
		Customfield11706              interface{}   `json:"customfield_11706"`
		Customfield11705              interface{}   `json:"customfield_11705"`
		Customfield11708              interface{}   `json:"customfield_11708"`
		Customfield11707              interface{}   `json:"customfield_11707"`
		Issuelinks                    []interface{} `json:"issuelinks"`
		Customfield11709              interface{}   `json:"customfield_11709"`
		Assignee                      struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"assignee"`
		Status struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			Name           string `json:"name"`
			ID             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Components []struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"components"`
		Customfield11421 interface{} `json:"customfield_11421"`
		Customfield11420 struct {
			Self  string `json:"self"`
			Value string `json:"value"`
			ID    string `json:"id"`
		} `json:"customfield_11420"`
		Customfield11300      interface{} `json:"customfield_11300"`
		Customfield11423      interface{} `json:"customfield_11423"`
		Customfield11422      interface{} `json:"customfield_11422"`
		Customfield11418      interface{} `json:"customfield_11418"`
		Aggregatetimeestimate interface{} `json:"aggregatetimeestimate"`
		Customfield11419      interface{} `json:"customfield_11419"`
		Creator               struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"creator"`
		Subtasks []interface{} `json:"subtasks"`
		Reporter struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"reporter"`
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Customfield10200 interface{} `json:"customfield_10200"`
		Customfield11405 interface{} `json:"customfield_11405"`
		Customfield10712 interface{} `json:"customfield_10712"`
		Customfield11800 interface{} `json:"customfield_11800"`
		Customfield11404 interface{} `json:"customfield_11404"`
		Customfield10713 interface{} `json:"customfield_10713"`
		Progress         struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"progress"`
		Votes struct {
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
			HasVoted bool   `json:"hasVoted"`
		} `json:"votes"`
		Worklog struct {
			StartAt    int           `json:"startAt"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Issuetype struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarID    int    `json:"avatarId"`
		} `json:"issuetype"`
		Timespent interface{} `json:"timespent"`
		Project   struct {
			Self       string `json:"self"`
			ID         string `json:"id"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			ProjectCategory struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"projectCategory"`
		} `json:"project"`
		Aggregatetimespent interface{} `json:"aggregatetimespent"`
		Customfield11401   interface{} `json:"customfield_11401"`
		Customfield11400   interface{} `json:"customfield_11400"`
		Customfield10302   interface{} `json:"customfield_10302"`
		Customfield10303   interface{} `json:"customfield_10303"`
		Customfield10304   interface{} `json:"customfield_10304"`
		Resolutiondate     string      `json:"resolutiondate"`
		Workratio          int         `json:"workratio"`
		Watches            struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
			IsWatching bool   `json:"isWatching"`
		} `json:"watches"`
		Created              string      `json:"created"`
		Customfield10300     interface{} `json:"customfield_10300"`
		Customfield10301     interface{} `json:"customfield_10301"`
		Customfield11744     interface{} `json:"customfield_11744"`
		Customfield11743     interface{} `json:"customfield_11743"`
		Customfield11746     string      `json:"customfield_11746"`
		Customfield11745     interface{} `json:"customfield_11745"`
		Customfield11747     string      `json:"customfield_11747"`
		Updated              string      `json:"updated"`
		Timeoriginalestimate interface{} `json:"timeoriginalestimate"`
		Description          string      `json:"description"`
		Customfield11100     interface{} `json:"customfield_11100"`
		Customfield11740     interface{} `json:"customfield_11740"`
		Customfield11101     interface{} `json:"customfield_11101"`
		Customfield11102     interface{} `json:"customfield_11102"`
		Customfield11103     interface{} `json:"customfield_11103"`
		Customfield11742     interface{} `json:"customfield_11742"`
		Customfield11104     interface{} `json:"customfield_11104"`
		Customfield11741     interface{} `json:"customfield_11741"`
		Timetracking         struct {
		} `json:"timetracking"`
		Customfield11458 interface{}   `json:"customfield_11458"`
		Customfield11733 interface{}   `json:"customfield_11733"`
		Customfield11457 interface{}   `json:"customfield_11457"`
		Customfield10006 interface{}   `json:"customfield_10006"`
		Customfield11732 interface{}   `json:"customfield_11732"`
		Customfield10007 interface{}   `json:"customfield_10007"`
		Customfield11735 interface{}   `json:"customfield_11735"`
		Customfield11734 interface{}   `json:"customfield_11734"`
		Customfield11737 string        `json:"customfield_11737"`
		Attachment       []interface{} `json:"attachment"`
		Customfield11736 interface{}   `json:"customfield_11736"`
		Customfield11739 interface{}   `json:"customfield_11739"`
		Customfield10803 interface{}   `json:"customfield_10803"`
		Customfield11738 interface{}   `json:"customfield_11738"`
		Summary          string        `json:"summary"`
		Customfield11450 interface{}   `json:"customfield_11450"`
		Customfield11452 interface{}   `json:"customfield_11452"`
		Customfield11451 interface{}   `json:"customfield_11451"`
		Customfield10000 interface{}   `json:"customfield_10000"`
		Customfield11454 interface{}   `json:"customfield_11454"`
		Customfield10001 string        `json:"customfield_10001"`
		Customfield11453 interface{}   `json:"customfield_11453"`
		Customfield10002 string        `json:"customfield_10002"`
		Customfield11731 interface{}   `json:"customfield_11731"`
		Customfield11730 interface{}   `json:"customfield_11730"`
		Customfield11722 interface{}   `json:"customfield_11722"`
		Customfield10511 interface{}   `json:"customfield_10511"`
		Customfield11446 interface{}   `json:"customfield_11446"`
		Customfield11721 string        `json:"customfield_11721"`
		Customfield10512 interface{}   `json:"customfield_10512"`
		Customfield11205 interface{}   `json:"customfield_11205"`
		Customfield11724 interface{}   `json:"customfield_11724"`
		Customfield11603 interface{}   `json:"customfield_11603"`
		Customfield11206 interface{}   `json:"customfield_11206"`
		Environment      interface{}   `json:"environment"`
		Customfield11602 interface{}   `json:"customfield_11602"`
		Customfield11207 interface{}   `json:"customfield_11207"`
		Customfield11726 interface{}   `json:"customfield_11726"`
		Customfield11605 interface{}   `json:"customfield_11605"`
		Customfield11725 interface{}   `json:"customfield_11725"`
		Customfield11604 interface{}   `json:"customfield_11604"`
		Customfield11728 interface{}   `json:"customfield_11728"`
		Duedate          interface{}   `json:"duedate"`
		Customfield11606 interface{}   `json:"customfield_11606"`
		Customfield11729 interface{}   `json:"customfield_11729"`
		Comment          struct {
			Comments   []interface{} `json:"comments"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			StartAt    int           `json:"startAt"`
		} `json:"comment"`
	} `json:"fields"`
}
