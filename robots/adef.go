package robots

type Player struct {
    Name            string          `json:"name"`
    Location        string          `json:"location"`
    Last_Room       string          `json:"last_room"`
    Items           []Item          `json:"inventory"`
}

type Room struct {
    Name            string          `json:"name"`
    Description     string          `json:"description"`
    Adjacent        []string        `json:"adjacent"`
    Parent          string          `json:"parent,omitempty"`
    Items           []Item          `json:"items,omitempty"`
    Widgets         []Widget        `json:"widgets,omitempty"`
}

type Item struct {
    Name            string          `json:"name"`
    Description     string          `json:"description"`
    Amount          int             `json:"amount,omitempty"`
}

type Widget struct {
    Name            string          `json:"name"`
    Description     string          `json:"description"`
    Actions         []string        `json:"actions"`
}
