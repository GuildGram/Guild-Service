package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GuildGram/Character-Service.git/handlers"

	"github.com/gorilla/mux"
)

func main() {
	handlers.StartMsgBrokerConnection()

	//old code
	l := log.New(os.Stdout, "guild-api", log.LstdFlags)

	ch := handlers.NewGuild(l)

	sm := mux.NewRouter()

	//handle get
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/guilds/getall", ch.GetGuilds)

	//should change to get by name for when user services are implemented
	getRouter.HandleFunc("/guilds/get{id:[0-9]+}", ch.GetGuild)

	//handle put
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/guilds/update{id:[0-9]+}", ch.UpdateGuild)

	//handle add
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/guilds/add", ch.AddGuild)

	//handle delete
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/guilds/delete{id:[0-9]+}", ch.DeleteGuild)

	//Server stuff for testing, will be deleted soon
	s := &http.Server{
		Addr:         ":9091",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("received kill signal, shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
