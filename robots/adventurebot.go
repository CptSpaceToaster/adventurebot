package robots

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Players = make(map[string]Player)
var Items = make(map[string]Item)
var Rooms = make(map[string]Room)
var Widgets = make(map[string]Widget) // Note, widgets are basically just items that can't be picked up.
var Actions = make(map[string]Action)
var Requirements = make(map[string]Requirement)

var Nouns = make(map[string][]string)
var Adjectives = make(map[string][]string)
var Verbs = make(map[string][]string)
var Adverbs = make(map[string][]string)

type AdventureBot struct {
}

func init() {
	n := new(AdventureBot)
	n.Load()
	RegisterRobot("adventurebot", func() (robot Robot) { return n })
}

func (p AdventureBot) Load() {
	fmt.Println("Registering Requirements")
	registerJson("../src/github.com/cptspacetoaster/adventurebot/requirements", handleRequirements)
	fmt.Println("Registering Actions")
	registerJson("../src/github.com/cptspacetoaster/adventurebot/actions", handleActions)
	fmt.Println("Registering Widgets")
	registerJson("../src/github.com/cptspacetoaster/adventurebot/widgets", handleWidgets)
	fmt.Println("Registering Items")
	registerJson("../src/github.com/cptspacetoaster/adventurebot/items", handleItems)
	fmt.Println("Registering Rooms")
	registerJson("../src/github.com/cptspacetoaster/adventurebot/rooms", handleRooms)
}

func registerJson(root_dir string, handle func(input []byte) error) {
	filepath.Walk(root_dir, func(path string, fi os.FileInfo, err error) error {
		if err == nil && fi.Mode().IsRegular() && filepath.Ext(path) == ".json" {
			input, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("Could not read " + path)
			} else {
				err = handle(input)
				if err != nil {
					fmt.Print(err)
					fmt.Println(":\n" + path)
				}
			}
		}
		return err
	})
}

func handleRequirements(input []byte) error {
	var r Requirement
	json.Unmarshal(input, &r)
	if r.Name != "" {
		Requirements[r.Name] = r
		//fmt.Println("Loaded: " + r.Name)
	} else {
		return errors.New("Input could not be Unmarshalled into type Requirement")
	}
	return nil
}

func handleActions(input []byte) error {
	var a Action
	json.Unmarshal(input, &a)
	if a.ID != "" {
		for _, s := range a.Commands {
			Verbs[strings.ToLower(s)] = append(Verbs[strings.ToLower(s)], a.ID)
		}
		for _, s := range a.Adverbs {
			Adverbs[strings.ToLower(s)] = append(Adverbs[strings.ToLower(s)], a.ID)
		}
		Actions[a.ID] = a
		//fmt.Println("Loaded: " + a.ID)
	} else {
		return errors.New("Input could not be Unmarshalled into type Action")
	}
	return nil
}

func handleWidgets(input []byte) error {
	var w Widget
	json.Unmarshal(input, &w)
	if w.ID != "" {
		for _, s := range w.Adjectives {
			Adjectives[strings.ToLower(s)] = append(Adjectives[strings.ToLower(s)], w.ID)
		}
		for _, s := range w.Names {
			Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], w.ID)
		}
		Widgets[w.ID] = w
		//fmt.Println("Loaded: " + w.ID)
	} else {
		return errors.New("Input could not be Unmarshalled into type Widget")
	}
	return nil
}

func handleItems(input []byte) error {
	var i Item
	json.Unmarshal(input, &i)
	if i.ID != "" {
		for _, s := range i.Adjectives {
			Adjectives[strings.ToLower(s)] = append(Adjectives[strings.ToLower(s)], i.ID)
		}
		for _, s := range i.Names {
			Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], i.ID)
		}
		Items[i.ID] = i
		//fmt.Println("Loaded: " + i.ID)
	} else {
		return errors.New("Input could not be Unmarshalled into type Item")
	}
	return nil
}

func handleRooms(input []byte) error {
	var r Room
	json.Unmarshal(input, &r)
	if r.ID != "" {
		for _, s := range r.Adjectives {
			Adjectives[strings.ToLower(s)] = append(Adjectives[strings.ToLower(s)], r.ID)
		}
		for _, s := range r.Names {
			Nouns[strings.ToLower(s)] = append(Nouns[strings.ToLower(s)], r.ID)
		}
		Rooms[r.ID] = r
		//fmt.Println("Loaded: " + r.ID)
	} else {
		return errors.New("Input could not be Unmarshalled into type Room")
	}
	return nil
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

func IsSymlink(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	return fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink, err
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

func RegisterPlayer(name string, id string) (player Player) {
	//TODO: Handle same names from unique servers - I think this is done
	player.Name = name
	player.ID = id
	Nouns[strings.ToLower(name)] = append(Nouns[strings.ToLower(name)], name)
	player.Location = "beach3"
	player.Last_Location = "beach3"
	//Lock
	Players[player.ID] = player
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
	/*
		if command.Channel_Name != "adventure" {
			return
		}
	*/
	if _, exist := Players[command.User_ID]; !exist {
		RegisterPlayer(command.User_Name, command.User_ID)
	}
	//RLock
	player := Players[command.User_ID]
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
		if len(nouns) == 0 && len(input) == 1 {
			sayDesc_R(command, Rooms[player.Location])
		} else if len(nouns) > 0 {
			var candidates []Widget
			//TODO: Users can only look at widgets at the moment
			for _, w := range Rooms[player.Location].Widgets {
				//find all widgets in the room
				//fmt.Print("Looking at " + r)
				if _, exist := Widgets[w]; exist {
					//does the actual entry exist?
					//fmt.Println(" Found!")
					if StringInSlice(nouns[0], Widgets[w].Names) {
						//the user typed in a valid widget
						candidates = append(candidates, Widgets[w])
					}
				}
			}
			if len(candidates) == 0 {
				if len(input) == 1 {
					sayDesc_R(command, Rooms[player.Location])
				} else {
					say(command, "I don't see a "+nouns[0]+" nearby.")
				}
			} else if len(candidates) == 1 {
				sayDesc_W(command, candidates[0])
			} else {
				if len(adjectives) == 0 {
					say(command, fmt.Sprintf("There is more than one type of %s nearby!", nouns[0]))
				} else {
					for _, w := range candidates {
						if StringInSlice(adjectives[0], w.Adjectives) {
							//the user typed in a valid widget noun
							sayDesc_W(command, w)
							break
						}
					}
				}
			}
		} else {
			say(command, "I can not look at that")
		}
	} else if action == "move" {
		if len(nouns) > 0 {
			//check for compass directions
			if Nouns[nouns[0]][0] == "norths" {
				if Rooms[player.Location].North != "" {
					player = move(command, player, Rooms[Rooms[player.Location].North])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "northeasts" {
				if Rooms[player.Location].North_East != "" {
					player = move(command, player, Rooms[Rooms[player.Location].North_East])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "easts" {
				if Rooms[player.Location].East != "" {
					player = move(command, player, Rooms[Rooms[player.Location].East])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "southeasts" {
				if Rooms[player.Location].South_East != "" {
					player = move(command, player, Rooms[Rooms[player.Location].South_East])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "souths" {
				if Rooms[player.Location].South != "" {

					player = move(command, player, Rooms[Rooms[player.Location].South])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "southwests" {
				if Rooms[player.Location].South_West != "" {
					player = move(command, player, Rooms[Rooms[player.Location].South_West])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "wests" {
				if Rooms[player.Location].West != "" {
					player = move(command, player, Rooms[Rooms[player.Location].West])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "northwests" {
				if Rooms[player.Location].North_West != "" {
					player = move(command, player, Rooms[Rooms[player.Location].North_West])
				} else {
					say(command, "There is no path to the "+Rooms[Nouns[nouns[0]][0]].Names[0]+" from here")
				}
			} else if Nouns[nouns[0]][0] == "ups" {
				if Rooms[player.Location].Up != "" {
					player = move(command, player, Rooms[Rooms[player.Location].Up])
				} else {
					say(command, "You can not "+verbs[0]+" "+nouns[0]+" at the moment")
				}
			} else if Nouns[nouns[0]][0] == "downs" {
				if Rooms[player.Location].Down != "" {
					player = move(command, player, Rooms[Rooms[player.Location].Down])
				} else {
					say(command, "You can not "+verbs[0]+" "+nouns[0]+" at the moment")
				}
			} else if Nouns[nouns[0]][0] == "backs" {
				if player.Location != player.Last_Location {
					player = move(command, player, Rooms[player.Last_Location])
				} else {
					say(command, "You are unable to retrace your steps")
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
						say(command, fmt.Sprintf("You are already in the %s.", nouns[0]))
					} else {
						say(command, fmt.Sprintf("You can not make it to the %s from here.", nouns[0]))
					}
					//TODO: Handle other -somethings-
				} else if len(candidates) == 1 {
					player = move(command, player, candidates[0])
				} else {
					if len(adjectives) == 0 {
						say(command, fmt.Sprintf("There is more than one type of %s nearby!", nouns[0]))
					} else {
						for _, r := range candidates {
							if StringInSlice(adjectives[0], r.Adjectives) {
								//the user typed in a valid adjacent noun
								player = move(command, player, r)
								break
							}
						}
					}
				}
			}

		} else {
			//user didn't tell me where to go
			say(command, "I do not know where you are trying to "+verbs[0]+".")
		}

		//check for current room
		//fmt.Println(Rooms[player.Location])
		/*
				say(command, fmt.Sprintf("%s is already in the %s", player.Name, player.Location))
				target = ""
			}
			//check for adjacent rooms
			//Can't go there
			if target != "" {
				say(command, fmt.Sprintf("%s can not %s there", player.Name, input[0]))
			}
		*/
	} else if action != "" {
		var candidates []Widget

		//TODO: Handle more than widgets
		if len(nouns) > 0 {
			for _, w := range Rooms[player.Location].Widgets {
				//find all widgets in the room
				//fmt.Print("Looking at " + r)
				if _, exist := Widgets[w]; exist {
					//does the actual entry exist?
					//fmt.Println(" Found!")
					if StringInSlice(nouns[0], Widgets[w].Names) {
						//the user typed in an action in this widget
						candidates = append(candidates, Widgets[w])
					}
				}
			}
			if len(candidates) == 0 {
				say(command, "You can't "+Actions[action].Commands[0]+" that here")
			} else if len(candidates) == 1 {
				if StringInSlice(action, candidates[0].Actions) {
					doAction(command, player, Actions[action])
				} else {
					say(command, "You can't "+Actions[action].Commands[0]+" the "+candidates[0].Display_Name)
				}
			} else {
				//TODO: Handle noun conflicts with adjectives
				say(command, "I can't handles this yet")
			}
		} else {
			say(command, "I don't know what you are trying to "+Actions[action].Commands[0])
		}

		//can the verb be done with the noun
		/*if len(nouns) == 0 {
			say(command, "I don't know what to " + action)
		} else {
			var ac_candidates []string
			var candidates []Widget
			//TODO: Users can only frob at widgets at the moment
			for _, w := range Rooms[player.Location].Widgets {
				//find all widgets in the room
				//fmt.Print("Looking at " + r)
				if _, exist := Widgets[w]; exist {
					//does the actual entry exist?
					//fmt.Println(" Found!")
					if StringInSlice(nouns[0], Widgets[w].Names) {
						//the user typed in a valid widget
						candidates = append(candidates, Widgets[w])
						if StringInSlice(action, )
					}
				}
			}
			if len(candidates) == 0 {
				say(command, "I don't see a " + nouns[0] + " nearby.")
			} else if len(candidates) == 1 {
				if
				//sayDesc_A(candidates[0])
			} else {
				if len(adjectives) == 0 {
					say(command, fmt.Sprintf("There is more than one type of %s nearby!", nouns[0]))
				} else {
					for _, w := range candidates {
						if StringInSlice(adjectives[0], w.Adjectives) {
							//the user typed in a valid widget noun
							sayDesc_W(command, w)
							break
						}
					}
				}
			}
		}*/

	} else {
		say(command, fmt.Sprintf("I do not understand what %s is trying to say", command.User_Name))
	}

	//Lock
	Players[player.ID] = player //Update the instance of the player in Players
	//Unlock
}

func move(command *SlashCommand, p Player, r Room) Player {
	p.Last_Location = p.Location
	p.Location = r.ID
	//say(command, fmt.Sprintf("%s is now in the %s", player.Name, Rooms[player.Location].Names[0]))
	sayDesc_R(command, r)
	return p
}

func sayDesc_R(command *SlashCommand, r Room) {
	say(command, command.User_Name+" is in the "+r.Display_Name+"\n______________________________________________\n"+r.Description)
}

func sayDesc_W(command *SlashCommand, w Widget) {
	say(command, w.Description)
}

func sayDesc_I(command *SlashCommand, i Item) {
	say(command, i.Description)
}

func doAction(command *SlashCommand, p Player, a Action) {
	//TODO: requirements
	say(command, a.Result)
	if a.Teleport != "" {
		if _, exist := Rooms[a.Teleport]; exist {
			move(command, p, Rooms[a.Teleport])
		}
	}
	//TODO: Gives
}

func say(command *SlashCommand, text string) {
	response := new(IncomingWebhook)
	response.Domain = command.Team_Domain
	response.Channel = command.Channel_ID
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
	return "Adventure bot!\n\tUsage: ?<command>\n\tOutput: An Adventure"
}
