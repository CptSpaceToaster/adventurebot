package robots

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

var Robots = make(map[string]func() Robot)
var Config = new(Configuration)
var ConfigDirectory = flag.String("c", ".", "Configuration directory (default .)")

func init() {
	flag.Parse()
	configFile := filepath.Join(*ConfigDirectory, "config.json")
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error opening config: ", err)
	}

	err = json.Unmarshal(config, Config)
	if err != nil {
		log.Fatal("Error parsing config: ", err)
	}

	log.Printf("Found %d configurations", len(Config.Credentials))
}

func RegisterRobot(command string, RobotInitFunction func() Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		log.Printf("Registered: %s", command)
		Robots[command] = RobotInitFunction
	}
}

func MakeIncomingWebhookCall(payload *IncomingWebhook) error {
	n := -1
	for index, c := range Config.Credentials {
		if payload.Domain == c.Domain {
			n = index
		}
	}

	if n < 0 {
		return errors.New("Domain not found in configuration")
	}

	webhook := url.URL{
		Scheme: "https",
		Host:   Config.Credentials[n].Domain + ".slack.com",
		Path:   "/services/hooks/incoming-webhook",
	}

	json_payload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	post_data := url.Values{}
	post_data.Set("payload", string(json_payload))
	post_data.Set("token", Config.Credentials[n].Token)

	webhook.RawQuery = post_data.Encode()
	resp, err := http.PostForm(webhook.String(), post_data)
	if resp.StatusCode != 200 {
		message := fmt.Sprintf("ERROR: Non-200 Response from Slack Incoming Webhook API: %s", resp.Status)
		log.Println(message)
	}

	return err
}
