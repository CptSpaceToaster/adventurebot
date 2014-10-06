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

var Nouns = make(map[string][]string)
var Adjectives = make(map[string][]string)
var Verbs = make(map[string][]string)
var Adverbs = make(map[string][]string)

//var filter = []string{"around", "by", "near", "towards", "to", "the", "a", "an", "on", "in"}

//var movement = []string{"go", "move", "walk", "frollic", "climb", "travel", "crawl", "roll", "skip", "stumble", "meander", "cartwheel"}
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
	RegisterRequirements("../src/github.com/cptspacetoaster/adventurebot/requirements")
	RegisterActions("../src/github.com/cptspacetoaster/adventurebot/actions")
	RegisterWidgets("../src/github.com/cptspacetoaster/adventurebot/widgets")
	RegisterItems("../src/github.com/cptspacetoaster/adventurebot/items")
	RegisterRooms("../src/github.com/cptspacetoaster/adventurebot/rooms")
}

func RegisterRequirements(dirloc string) {
	fmt.Println("Registering Requirements")

	names := getDirNames(dirloc)

	for _, element := range names {
		input, err2 := ioutil.ReadFile(dirloc + "/" + element)
		if err2 != nil {
			fmt.Println("Could not read " + element)
		} else {
			var r Requirement
			json.Unmarshal(input, &r)
			if r.Name != "" {
				Requirements[r.Name] = r
				//fmt.Println("Loaded: " + r.Name)
			}
		}
	}
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
					Verbs[strings.ToLower(s)] = append(Verbs[strings.ToLower(s)], a.ID)
				}
				for _, s := range a.Adverbs {
					Verbs[strings.ToLower(s)] = append(Adverbs[strings.ToLower(s)], a.ID)
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
					Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], w.ID)
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
					Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], i.ID)
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
					Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], r.ID)
				}
				for _, s := range r.Adjectives {
					Adjectives[strings.ToLower(s)] = append(Adjectives[strings.ToLower(s)], r.ID)
				}
				Rooms[r.ID] = r
				//fmt.Println("Loaded: " + r.ID)
			} else {
				fmt.Println("Warning: " + element + " could not be read.")
			}
		}
	}
	/*
		fmt.Print("\n")
		for k, _ := range Nouns {
			fmt.Print(k + ", ")
		}
		fmt.Print("\n\n")
		for k, _ := range Adjectives {
			fmt.Print(k + ", ")
		}
		fmt.Print("\n\n")
		for k, _ := range Verbs {
			fmt.Print(k + ", ")
		}
		fmt.Print("\n\n")
		for k, _ := range Adverbs {
			fmt.Print(k + ", ")
		}
		fmt.Print("\n\n")
	*/
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

	//remove any question marks, split strings by spaces
	input := strings.Split(strings.Trim(strings.Trim(strings.ToLower(command.Text), "?"), " "), " ")
	//create new arrays to filter user input
	nouns := make([]string, 0)
	adjectives := make([]string, 0)
	verbs := make([]string, 0)
	adverbs := make([]string, 0)

	for _, s := range input {
		//fmt.Print(s + ", ")
		if _, exist := Nouns[s]; exist {
			nouns = append(nouns, s)
		}
		if _, exist := Adjectives[s]; exist {
			adjectives = append(adjectives, s)
		}
		if _, exist := Verbs[s]; exist {
			verbs = append(verbs, s)
		}
		if _, exist := Adverbs[s]; exist {
			adverbs = append(adverbs, s)
		}
	}
	//fmt.Print("\n")

	for _, s := range adverbs {
		fmt.Print(" a: " + s)
	}
	for _, s := range verbs {
		fmt.Print(" v: " + s)
	}
	for _, s := range adjectives {
		fmt.Print(" a: " + s)
	}
	for _, s := range nouns {
		fmt.Print(" n: " + s)
	}
	fmt.Print("\n")

	action := ""
	//parse through verbs and adverbs to determine which ones the user was talking about
	for _, v := range verbs {
		//For each local verb the user types
		for _, s := range Verbs[v] {
			//get all of the Verbs in the Verb Map that match
			if len(adverbs) == 0 {
				//the user did not type an adverb... use the the matching verb
				action = s
				break
			} else {
				//user DID type an adverb... what did they want?
				for _, a := range adverbs {
					fmt.Println(">")
					//look at each adverb the user typed, and use entries in Actions[] to match
					if len(Actions[s].Adverbs) == 0 {
						fmt.Println("2")
						//there are no adverbs in the Verb Map to help decide,
						//we'll take the given verb
						action = s
						break
					} else {
						fmt.Println(len(Actions[s].Adverbs))
						//If there is an adverb though, we'll use it instead
						if StringInSlice(a, Actions[s].Adverbs) {
							//The last match will be used
							action = s
							break
						}
					}
				}
			}
		}
	}
	//fmt.Println(action)

	if action == "look" {
		SayDesc(Rooms[player.Location])
	} else if action == "move" {
		//smash room name together TODO
		//fmt.Println(nouns)

		if len(nouns) > 0 {
			//check for compass directions
			if Nouns[nouns[0]][0] == "norths" {
				if Rooms[player.Location].North != "" {
					player = move(player, Rooms[Rooms[player.Location].North])
				}
			} else if Nouns[nouns[0]][0] == "northeasts" {
				if Rooms[player.Location].North_East != "" {
					player = move(player, Rooms[Rooms[player.Location].North_East])
				}
			} else if Nouns[nouns[0]][0] == "easts" {
				if Rooms[player.Location].East != "" {
					player = move(player, Rooms[Rooms[player.Location].East])
				}
			} else if Nouns[nouns[0]][0] == "southeasts" {
				if Rooms[player.Location].South_East != "" {
					player = move(player, Rooms[Rooms[player.Location].South_East])
				}
			} else if Nouns[nouns[0]][0] == "souths" {
				if Rooms[player.Location].South != "" {

					player = move(player, Rooms[Rooms[player.Location].South])
				}
			} else if Nouns[nouns[0]][0] == "southwests" {
				if Rooms[player.Location].South_West != "" {
					player = move(player, Rooms[Rooms[player.Location].South_West])
				}
			} else if Nouns[nouns[0]][0] == "wests" {
				if Rooms[player.Location].West != "" {
					player = move(player, Rooms[Rooms[player.Location].West])
				}
			} else if Nouns[nouns[0]][0] == "northwests" {
				if Rooms[player.Location].North_West != "" {
					player = move(player, Rooms[Rooms[player.Location].North_West])
				}
			} else {
				//Not a compass direction, user may have typed a room name

				//make a list of nearby rooms that use the noun the user typed
				var candidates []Room

				for _, r := range Rooms[player.Location].Adjacent {
					//find all adjacent rooms
					//fmt.Print("Looking at " + r)
					if _, exist := Rooms[r]; exist {
						//does the actual entry exist?
						//fmt.Println(" Found!")
						if StringInSlice(nouns[0], Rooms[r].Names) {
							//the user typed in a valid adjacent noun
							candidates = append(candidates, Rooms[r])
						}
					}
				}

				if len(candidates) == 0 {
					if StringInSlice(nouns[0], Rooms[player.Location].Names) {
						Say(fmt.Sprintf("You are already in the %s.", nouns[0]))
					} else {
						Say(fmt.Sprintf("You can not make it to the %s from here.", nouns[0]))
					}
				} else if len(candidates) == 1 {
					player = move(player, candidates[0])
				} else {
					if len(adjectives) == 0 {
						Say(fmt.Sprintf("There is more than one type of %s nearby!", nouns[0]))
					} else {
						for _, r := range candidates {
							if StringInSlice(adjectives[0], r.Adjectives) {
								//the user typed in a valid adjacent noun
								player = move(player, r)
								break
							}
						}
					}
				}
			}

		} else {
			//user didn't tell me where to go
			Say("I do not know where you are trying to go.")
		}

		//check for current room
		//fmt.Println(Rooms[player.Location])
		/*
			if StringInSlice(target, Rooms[player.Location].Names) {
				Say(fmt.Sprintf("%s is already in the %s", player.Name, player.Location))
				target = ""
			}
			//check for adjacent rooms
			//Can't go there
			if target != "" {
				Say(fmt.Sprintf("%s can not %s there", player.Name, input[0]))
			}
		*/
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
	Say(r.Display_Name + "\n______________________________________________\n" + r.Description)
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

func (p AdventureBot) Description() (description string) {
	//Ehhh... todo?
	return "Adventure bot!\n\tUsage: /adventure\n\tExpected Response: @user: Pong!"
}
