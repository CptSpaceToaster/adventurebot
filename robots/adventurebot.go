package robots

import (
	//"io"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Players = make(map[string]Player)
var Items = make(map[string]Item)
var Rooms = make(map[string]Room)
var Widgets = make(map[string]Widget)

type AdventureBot struct {
}

func init() {
	n := new(AdventureBot)
	n.Load()
	RegisterRobot("adventurebot", func() (robot Robot) { return n })
}

func (p AdventureBot) Load() {
	fmt.Println("Loading File Configurations")
	RegisterRooms("../src/github.com/cptspacetoaster/adventurebot/rooms")
}

func RegisterRooms(roomdirloc string) {
	fmt.Println("Registering Rooms")
	roomdir, err0 := os.Open(roomdirloc)
	if err0 != nil {
		fmt.Println("rooms directory is missing")
		os.Exit(1)
	}

	names, err1 := roomdir.Readdirnames(0)
	if err1 != nil {
		fmt.Println("directory names could not be read")
		os.Exit(1)
	}

	for _, element := range names {
		input, err2 := ioutil.ReadFile(roomdirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var r Room
			json.Unmarshal(input, &r)
			if r.ID != "" {
				Rooms[r.ID] = r
				fmt.Println("Loaded: " + r.ID)
			}
		}
	}
}

func RegisterPlayer(name string) (player Player) {
	player.Name = name
	player.Location = "beach3"
	player.Last_Location = "beach3"
	//Lock
	Players[player.Name] = player
	//Unlock
	fmt.Println("Registered " + name)
	return player
}

func (p AdventureBot) Run(command *SlashCommand) (slashCommandImmediateReturn string) {
	go p.DeferredAction(command)
	return ""
}

func (p AdventureBot) DeferredAction(command *SlashCommand) {
	//if the command isn't issued from adventure, then ignore it. This shouldn't be possible anymore
	if command.Channel_Name != "adventure" {
		return
	}

	if _, exist := Players[command.User_Name]; !exist {
		RegisterPlayer(command.User_Name)
	}
	//RLock
	player := Players[command.User_Name]
	//RUnlock
	action := strings.Split(strings.Trim(strings.Trim(strings.ToLower(command.Text), "?"), " "), " ")
	fmt.Println(action)

	if action[0] == "look" {
		Say(Rooms[player.Location].Description)
	} else if action[0] == "go" && len(action) >= 2 {
		if _, exist := Rooms[action[1]]; exist {
			fmt.Println(Rooms[player.Location])
			fmt.Println(action[1])

			if player.Location == action[1] {
				Say(fmt.Sprintf("%s is already in the %s", player.Name, player.Location))
			} else if StringInSlice(action[1], Rooms[player.Location].Adjacent) {
				player.Last_Location = player.Location
				player.Location = action[1]
				Say(fmt.Sprintf("%s is now in the %s", player.Name, player.Location))
			}
		} else {
			Say(fmt.Sprintf("%s can not %s to %s", player.Name, action[0], action[1]))
		}
	} else {
		Say(fmt.Sprintf("I do not understand what %s is trying to say", command.User_Name))
	}
	//Lock
	Players[player.Name] = player //Update the instance of the player in Players
	//Unlock
}

func Say(text string) {
	response := new(IncomingWebhook)
	response.Channel = "C02HR18H4"
	response.Username = "Adventure Bot"
	response.Text = text
	fmt.Println(text)
	response.Icon_Emoji = ":adventureboticon:"
	response.Unfurl_Links = true
	response.Parse = "full"
	MakeIncomingWebhookCall(response)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (p AdventureBot) Description() (description string) {
	//Ehhh... todo?
	return "Adventure bot!\n\tUsage: /adventure\n\tExpected Response: @user: Pong!"
}
