package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Guild struct {
	OwnerID      int         `json:"userid"`
	GuildID      string      `json:"guildid"`
	GuildOwnerID int         `json:"guildowner"`
	Roster       []Character `json:"roster"`
	Bio          string      `json:"bio"`
	Progression  string      `json:"progression"`
}

type Character struct {
	UserID           int    `json:"userid"`
	Class            string `json:"class"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"region-server"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	Ilvl             int    `json:"ilvl"`
	GuildID          string `json:"guildid"`
	GuildRole        string `json:"guildRole"`
}

func (c *Guild) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

type Guilds []*Guild

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
	_, pos, err := findGuild(id)
	if err != nil {
		return err
	}

	c.OwnerID = id
	guildList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("char Not found")

func findGuild(id int) (*Guild, int, error) {
	for i, c := range guildList {
		if c.OwnerID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func AddGuild(c *Guild) {
	guildList = append(guildList, c)
}

func DeleteGuild(id int) error {
	_, pos, err := findGuild(id)
	if err != nil {
		return err
	}
	guildList[pos] = guildList[len(guildList)-1]
	guildList[len(guildList)-1] = nil
	guildList = guildList[:len(guildList)-1]
	return err
}

func GetGuild(id int) (*Guild, error) {
	_, pos, err := findGuild(id)
	if err != nil {
		return nil, err
	}
	return guildList[pos], err
}

func AddRosterInfo(c Character) {
	//in a database
	for _, g := range guildList {
		if g.GuildID == c.GuildID {
			g.Roster = append(g.Roster, c)
		}
	}
}

func AddMultipleRosterInfo(c []Character) {
	for _, char := range c {
		AddRosterInfo(char)
	}
}

var guildList = []*Guild{
	{
		OwnerID:      1,
		GuildID:      "G1",
		GuildOwnerID: 1,
		Roster:       []Character{},
		Bio:          "lorem ipsum dolores",
		Progression:  "ABC:1",
	},
}
