package robots

type Player struct {
    Name            string          `json:"name"`
    Location        string          `json:"location"`
    Last_Room       string          `json:"last_room"`
    Inventory       []string        `json:"inventory"`
}

type Room struct {
    Names           []string        `json:"names"`
    ID              string          `json:"id"`
    Description     string          `json:"description"`
    Parent          string          `json:"parent,omitempty"`

    Has_Ceiling     bool            `json:"has_ceiling"`
    Has_Floor       bool            `json:"has_floor"`
    Is_Outside      bool            `json:"is_outside"`

    Adjacent        []string        `json:"adjacent,omitempty"`
    North           string          `json:"north,omitempty"`
    North_East      string          `json:"north_east,omitempty"`
    East            string          `json:"east,omitempty"`
    South_East      string          `json:"south_east,omitempty"`
    South           string          `json:"south,omitempty"`
    South_West      string          `json:"south_west,omitempty"`
    West            string          `json:"west,omitempty"`
    North_West      string          `json:"north_west,omitempty"`
    Up              string          `json:"up,omitempty"`
    Down            string          `json:"down,omitempty"`

    Items           []string        `json:"items,omitempty"`
    Widgets         []string        `json:"widgets,omitempty"`
}

type Item struct {
    Names           []string        `json:"name"`
    ID              string          `json:"id
    Description     string          `json:"description"`
    Amount          int             `json:"amount,omitempty"`
    Actions         []Action        `json:"actions,omitempty"`
}

type Widget struct {
    Names           []string        `json:"name"`
    ID              string          `json:"id"`
    Description     string          `json:"description"`
    Actions         []Action        `json:"actions,omitempty"`
}

type Action struct {
    Commands        []string        `json:"commands"`
    Result          string          `json:"result"`
    Teleport        string          `json:"teleport,omitempty"`
    Requires        []Requirement   `json:"requires,omitempty"`
    Gives           []string        `json:"gives,omitempty"`
}

type Requirement struct {
    Name            string          `json:"name"`
    Consumed        bool            `json:"consumed"`
}