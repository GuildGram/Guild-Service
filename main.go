package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GuildGram/Character-Service.git/handlers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	//old code
	l := log.New(os.Stdout, "guild-api", log.LstdFlags)

	ch := handlers.NewGuild(l)

	router := mux.NewRouter()

	//handle get
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/guilds/getall", ch.GetGuilds)
	getRouter.HandleFunc("/guilds/addchars", ch.AddCharToRoster)
	getRouter.HandleFunc("/guilds/getroster{id:[0-9]+}", ch.GetRoster)

	//should change to get by name for when user services are implemented
	getRouter.HandleFunc("/guilds/get{id:[0-9]+}", ch.GetGuild)

	//handle put
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/guilds/update{id:[0-9]+}", ch.UpdateGuild)

	//handle add
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/guilds/add", ch.AddGuild)

	//handle delete
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/guilds/delete{id:[0-9]+}", ch.DeleteGuild)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	//Server stuff
	s := &http.Server{
		Addr:         ":9091",
		Handler:      handler,
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
