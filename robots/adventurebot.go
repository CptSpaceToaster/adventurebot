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
var Actions = make(map[string]Action)
var Requirements = make(map[string]Requirement)

var Verbs = make(map[string]string)
var Nouns = make(map[string]string)

var filter = []string{"around", "by", "near", "towards", "to", "the", "a", "an", "on", "in"}
var movement = []string{"go", "move", "walk", "frollic", "climb", "travel", "crawl", "roll", "skip", "stumble", "meander", "cartwheel"}
var norths = []string{"n", "north", "nort", "norht", "norh"}
var northeasts = []string{"ne", "northeast", "norhteast", "norteast", "norheast", "norhteas", "norteas", "norheas"}
var easts = []string{"e", "east", "eas"}
var southeasts = []string{"se", "southeast", "souhteast", "souteast", "souheast", "souhteas", "souteas", "souheas"}
var souths = []string{"s", "south", "sout", "souht", "souh"}
var southwests = []string{"sw", "southwest", "souhtwest", "soutwest", "souhwest", "souhtwes", "soutwes", "souhwes"}
var wests = []string{"w", "west", "wes"}
var northwests = []string{"nw", "northwest", "norhtwest", "nortwest", "norhwest", "norhtwes", "nortwes", "norhwes"}

type AdventureBot struct {
}

func init() {
	n := new(AdventureBot)
	n.Load()
	RegisterRobot("adventurebot", func() (robot Robot) { return n })
}

func (p AdventureBot) Load() {
	//RegisterRequirements("../src/github.com/cptspacetoaster/adventurebot/requirements")
	RegisterActions("../src/github.com/cptspacetoaster/adventurebot/actions")
	RegisterWidgets("../src/github.com/cptspacetoaster/adventurebot/widgets")
	RegisterItems("../src/github.com/cptspacetoaster/adventurebot/items")
	RegisterRooms("../src/github.com/cptspacetoaster/adventurebot/rooms")
}

func RegisterActions(dirloc string) {
	fmt.Println("Registering Actions")

	names := getDirNames(dirloc)

	for _, element := range names {
		input, err2 := ioutil.ReadFile(dirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var a Action
			json.Unmarshal(input, &a)
			if a.ID != "" {
				for _, s := range a.Commands {
					Verbs[s] = a.ID
				}
				Actions[a.ID] = a
				//fmt.Println("Loaded: " + a.ID)
			}
		}
	}
}

func RegisterWidgets(dirloc string) {
	fmt.Println("Registering Widgets")

	names := getDirNames(dirloc)

	for _, element := range names {
		input, err2 := ioutil.ReadFile(dirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var w Widget
			json.Unmarshal(input, &w)
			if w.ID != "" {
				for _, s := range w.Names {
					Nouns[s] = w.ID
				}
				Widgets[w.ID] = w
				//fmt.Println("Loaded: " + w.ID)
			}
		}
	}
}

func RegisterItems(dirloc string) {
	fmt.Println("Registering Items")

	names := getDirNames(dirloc)

	for _, element := range names {
		input, err2 := ioutil.ReadFile(dirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var i Item
			json.Unmarshal(input, &i)
			if i.ID != "" {
				for _, s := range i.Names {
					Nouns[s] = i.ID
				}
				Items[i.ID] = i
				//fmt.Println("Loaded: " + i.ID)
			}
		}
	}
}

func RegisterRooms(dirloc string) {
	fmt.Println("Registering Rooms")

	names := getDirNames(dirloc)

	for _, element := range names {
		input, err2 := ioutil.ReadFile(dirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var r Room
			json.Unmarshal(input, &r)
			if r.ID != "" {
				for _, s := range r.Names {
					Nouns[s] = r.ID
				}
				Rooms[r.ID] = r
				//fmt.Println("Loaded: " + r.ID)
			}
		}
	}
}

func getDirNames(dirloc string) (names []string) {
	dir, err0 := os.Open(dirloc)
	if err0 != nil {
		fmt.Println(fmt.Sprintf("%s directory is missing", dirloc))
		os.Exit(1)
	}

	names, err1 := dir.Readdirnames(0)
	if err1 != nil {
		fmt.Println("directory names could not be read")
		os.Exit(1)
	}
	return
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
	action := StringParse(command.Text)
	fmt.Println(action)

	if action[0] == "look" {
		SayDesc(Rooms[player.Location])
	} else if StringInSlice(action[0], movement) && len(action) >= 2 {
		//smash room name together
		var target = ""
		for index, s := range action {
			if index == len(action)-1 {
				target += s
			} else if index > 0 {
				target += s + " "
			}
		}
		fmt.Println(Rooms[player.Location])
		fmt.Println("-" + target + "-")
		//check for compass directions
		if StringInSlice(target, norths) {
			if Rooms[player.Location].North != "" {
				player = move(player, Rooms[Rooms[player.Location].North])
				target = ""
			}
		}
		if StringInSlice(target, northeasts) {
			if Rooms[player.Location].North_East != "" {
				player = move(player, Rooms[Rooms[player.Location].North_East])
				target = ""
			}
		}
		if StringInSlice(target, easts) {
			if Rooms[player.Location].East != "" {
				player = move(player, Rooms[Rooms[player.Location].East])
				target = ""
			}
		}
		if StringInSlice(target, southeasts) {
			if Rooms[player.Location].South_East != "" {
				player = move(player, Rooms[Rooms[player.Location].South_East])
				target = ""
			}
		}
		if StringInSlice(target, souths) {
			if Rooms[player.Location].South != "" {
				player = move(player, Rooms[Rooms[player.Location].South])
				target = ""
			}
		}
		if StringInSlice(target, southwests) {
			if Rooms[player.Location].South_West != "" {
				player = move(player, Rooms[Rooms[player.Location].South_West])
				target = ""
			}
		}
		if StringInSlice(target, wests) {
			if Rooms[player.Location].West != "" {
				player = move(player, Rooms[Rooms[player.Location].West])
				target = ""
			}
		}
		if StringInSlice(target, northwests) {
			if Rooms[player.Location].North_West != "" {
				player = move(player, Rooms[Rooms[player.Location].North_West])
				target = ""
			}
		}
		//check for current room
		if StringInSlice(target, Rooms[player.Location].Names) {
			Say(fmt.Sprintf("%s is already in the %s", player.Name, player.Location))
			target = ""
		}
		//check for adjacent rooms
		for _, r := range Rooms[player.Location].Adjacent {
			fmt.Print("Looking at " + r)
			if _, exist := Rooms[r]; exist {
				fmt.Println(" Found!")
				if StringInSlice(target, Rooms[r].Names) {
					player = move(player, Rooms[r])
					target = ""
					break
				}
			}
		}
		//Can't go there
		if target != "" {
			Say(fmt.Sprintf("%s can not %s there", player.Name, action[0]))
		}

	} else {
		Say(fmt.Sprintf("I do not understand what %s is trying to say", command.User_Name))
	}
	//Lock
	Players[player.Name] = player //Update the instance of the player in Players
	//Unlock
}

func move(p Player, r Room) Player {
	p.Last_Location = p.Location
	p.Location = r.ID
	//Say(fmt.Sprintf("%s is now in the %s", player.Name, Rooms[player.Location].Names[0]))
	SayDesc(r)
	return p
}

func SayDesc(r Room) {
	Say(r.Names[0] + "\n______________________________________________\n" + r.Description)
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
		//fmt.Println(strings.ToLower(b) + " --- " + a)
		if strings.ToLower(b) == a {
			return true
		}
	}
	return false
}

func StringParse(in string) []string {
	//remove any question marks, split strings by spaces
	action := strings.Split(strings.Trim(strings.Trim(strings.ToLower(in), "?"), " "), " ")
	//create a new array
	out := make([]string, 0)
	//fill it in order, while stilling elements in the filter.
	for _, s := range action {
		if !StringInSlice(s, filter) {
			out = append(out, s)
		}
	}
	return out
}

func (p AdventureBot) Description() (description string) {
	//Ehhh... todo?
	return "Adventure bot!\n\tUsage: /adventure\n\tExpected Response: @user: Pong!"
}
