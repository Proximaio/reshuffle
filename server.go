package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var client Doer
var endpoints Endpoints

type Server struct {
	Endpoints Endpoints
	Client    Doer
	Port      string
	Router    *gin.Engine
}

func SetUp(opts Server) *Server {
	client = opts.Client
	endpoints = opts.Endpoints

	router := gin.Default()
	router.GET("/deck/:count", Deck)

	return &Server{
		Endpoints: endpoints,
		Client:    client,
		Port:      opts.Port,
		Router:    router,
	}
}

func (s *Server) Start() {
	s.Router.Run()
}

func Deck(c *gin.Context) {
	count := c.Param("count")
	i, err := strconv.ParseUint(count, 10, 0)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad card count")
		return
	}
	countUint := uint(i)
	fetchedDeck, err := GetDeck(DeckOpts{
		Shuffle:   true,
		Cards:     countUint,
		Client:    client,
		Endpoints: endpoints,
	})
	if err != nil {
		c.String(http.StatusNotFound, "Oops")
		return
	}
	c.String(http.StatusOK, fetchedDeck.String())
}
