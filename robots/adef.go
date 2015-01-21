package robots

type Player struct {
	Name          string   `json:"name"`
	ID            string   `json:"id"`
	Location      string   `json:"location"`
	Last_Location string   `json:"last_location"`
	Inventory     []string `json:"inventory"`
}

type Room struct {
	Adjectives   []string `json:"adjectives,omitempty"`
	Names        []string `json:"names"`
	ID           string   `json:"id"`
	Display_Name string   `json:"display_name"`
	Description  string   `json:"description"`
	Parent       string   `json:"parent,omitempty"`
	Has_Ceiling  bool     `json:"has_ceiling"`
	Has_Floor    bool     `json:"has_floor"`
	Is_Outside   bool     `json:"is_outside"`
	Adjacent     []string `json:"adjacent,omitempty"`
	North        string   `json:"north,omitempty"`
	North_East   string   `json:"north_east,omitempty"`
	East         string   `json:"east,omitempty"`
	South_East   string   `json:"south_east,omitempty"`
	South        string   `json:"south,omitempty"`
	South_West   string   `json:"south_west,omitempty"`
	West         string   `json:"west,omitempty"`
	North_West   string   `json:"north_west,omitempty"`
	Up           string   `json:"up,omitempty"`
	Down         string   `json:"down,omitempty"`
	Items        []string `json:"items,omitempty"`
	Widgets      []string `json:"widgets,omitempty"`
}

type Item struct {
	Adjectives   []string `json:"adjectives,omitempty"`
	Names        []string `json:"names"`
	ID           string   `json:"id"`
	Display_Name string   `json:"display_name"`
	Description  string   `json:"description"`
	Quantity     int      `json:"quantity,omitempty"` //zero quantity implies infinity
	Actions      []string `json:"actions,omitempty"`
}

type Widget struct {
	Adjectives   []string `json:"adjectives,omitempty"`
	Names        []string `json:"names"`
	ID           string   `json:"id"`
	Display_Name string   `json:"display_name"`
	Description  string   `json:"description"`
	Actions      []string `json:"actions,omitempty"`
}

type Action struct {
	Commands []string      `json:"commands"`
	Adverbs  []string      `json:"adverbs,omitempty"`
	ID       string        `json:"id"`
	Result   string        `json:"result,omitempty"`
	Teleport string        `json:"teleport,omitempty"`
	Requires []Requirement `json:"requires,omitempty"`
	Gives    []string      `json:"gives,omitempty"`
}

type Requirement struct {
	Name     string `json:"name"`
	Consumed bool   `json:"consumed"`
}
