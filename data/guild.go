package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Guild struct {
	OwnerID     int         `json:"userid"`
	GuildID     string      `json:"guildid"`
	Roster      []Character `json:"roster"`
	Bio         string      `json:"bio"`
	Progression string      `json:"progression"`
}

type Character struct {
	UserID           int    `json:"userid"`
	Class            string `json:"class"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"regionserver"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	Ilvl             int    `json:"ilvl"`
	GuildID          string `json:"guildid"`
	GuildRole        string `json:"guildrole"`
}

type Guilds []*Guild
type Characters []Character

func (c *Characters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Guild) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Guild) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Guilds) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func GetGuilds() Guilds {
	return guildList
}

func UpdateGuild(id int, c *Guild) error {
	_, pos, err := findGuilByOwner(id)
	if err != nil {
		return err
	}

	c.OwnerID = id
	guildList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("char Not found")

func GetRoster(id string) (Characters, error) {
	g, err := GetGuild(id)
	if err != nil {
		return nil, err
	}
	//print(g.Roster[0].CharaterName + "/n" + g.Roster[1].CharaterName)
	return g.Roster, nil
}

func findGuilByOwner(id int) (*Guild, int, error) {
	for i, c := range guildList {
		if c.OwnerID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func findGuildById(id string) (*Guild, int, error) {
	for i, c := range guildList {
		if c.GuildID == ("G" + id) {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func AddGuild(c *Guild) {
	guildList = append(guildList, c)
}

func DeleteGuild(id int) error {
	_, pos, err := findGuilByOwner(id)
	if err != nil {
		return err
	}
	guildList[pos] = guildList[len(guildList)-1]
	guildList[len(guildList)-1] = nil
	guildList = guildList[:len(guildList)-1]
	return err
}

func GetGuild(id string) (*Guild, error) {
	_, pos, err := findGuildById(id)
	if err != nil {
		return nil, err
	}
	return guildList[pos], err
}

func AddRosterInfo(c Character) {
	//in a database
	for _, g := range guildList {
		if g.GuildID == c.GuildID {
			if CheckRoster(c, g) {
				g.Roster = append(g.Roster, c)
			}
		}
	}
}

func ReplaceRoster(c []Character) {
	for _, g := range guildList {
		if g.GuildID == c[0].GuildID {
			g.Roster = c
		}
	}
}

func CheckRoster(check Character, g *Guild) bool {
	for _, g := range g.Roster {
		if g.UserID == check.UserID {
			return false
		}
	}
	return true
}

var guildList = []*Guild{
	{
		OwnerID:     1,
		GuildID:     "G1",
		Roster:      []Character{},
		Bio:         "lorem ipsum dolores",
		Progression: "ABC:1",
	},
	{
		OwnerID:     2,
		GuildID:     "G2",
		Roster:      []Character{},
		Bio:         "dolores ipsum lorem",
		Progression: "ABC:5",
	},
}
