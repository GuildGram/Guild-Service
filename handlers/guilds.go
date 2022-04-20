package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/gorilla/mux"
)

type Guild struct {
	l *log.Logger
}

func NewGuild(l *log.Logger) *Guild {
	return &Guild{l}
}

func (c *Guild) AddCharToRoster(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	gId := ("G" + key)

	c.l.Println("HANDLE MESSAGE BROKER REQUEST CHARS GUILD")

	//initialize message broker connection
	char, err := ReqCharactersByGID(gId)
	if err != nil {
		log.Print("unable to receive char info", err)
		return
	}
	data.AddMultipleRosterInfo(char)
}

func (c *Guild) UpdateGuild(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE PUT GUILD", id)

	char := &data.Guild{}
	err = char.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateGuild(id, char)
	if err == data.ErrCharNotFound {
		http.Error(rw, "Char not found", http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
}

func (c *Guild) GetGuilds(rw http.ResponseWriter, h *http.Request) {
	c.l.Println("HANDLE GET GUILDS")
	listChars := data.GetGuilds()
	err := listChars.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusInternalServerError)
	}
}

func (c *Guild) AddGuild(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("HANDLE POST GUILD")

	guild := &data.Guild{}
	err := guild.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddGuild(guild)
}

func (c *Guild) DeleteGuild(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE DELETE GUILD", id)

	err = data.DeleteGuild(id)
	if err != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
}

func (c *Guild) GetGuild(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE GET 1 GUILD	", id)

	char, err2 := data.GetGuild(id)
	if err2 != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
	char.ToJSON(rw)
}
