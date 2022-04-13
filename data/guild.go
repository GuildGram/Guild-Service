package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Guild struct {
	OwnerID int `json:"userid"`
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

var ErrCharNotFound = fmt.Errorf("Char Not found")

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

var guildList = []*Guild{
	&Guild{
		OwnerID: 1,
	},
}
