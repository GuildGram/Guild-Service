package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Guild struct {
	OwnerID     int         `json:"userid"`
	GuildID     string      `json:"guildid"`
	GuildOwner  int         `json:"guildowner"`
	Roster      []Character `json:"roster"`
	Bio         string      `json:"bio"`
	Progression string      `json:"progression"`
}

type Character struct {
	UserID           int    `json:"userid"`
	Class            string `json:"class"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"region-server"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	Ilvl             int    `json:"ilvl"`
	GuildName        string `json:"guildName"`
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
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}

	c.OwnerID = id
	guildList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("char Not found")

func findChar(id int) (*Guild, int, error) {
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
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}
	guildList[pos] = guildList[len(guildList)-1]
	guildList[len(guildList)-1] = nil
	guildList = guildList[:len(guildList)-1]
	return err
}

func GetGuild(id int) (*Guild, error) {
	_, pos, err := findChar(id)
	if err != nil {
		return nil, err
	}
	return guildList[pos], err
}

func AddRosterInfo(c *Character) {
	for i, g := range guildList {
		g.Roster = append(g.Roster, *c)
		_ = i
	}
	println(guildList)
}

var guildList = []*Guild{
	&Guild{
		OwnerID:    1,
		GuildID:    "G1",
		GuildOwner: 1,
		Roster: []Character{
			Character{
				UserID:           1,
				CharaterName:     "Nemoi",
				Class:            "Striker",
				RegionServerName: "EUC-Sceptrum",
				CharacterLevel:   53,
				RosterLevel:      68,
				Ilvl:             1355,
				GuildName:        "FontysICT",
				GuildRole:        "Owner",
			},
		},
		Bio:         "lorem ipsum dolores",
		Progression: "ABC:1",
	},
}
