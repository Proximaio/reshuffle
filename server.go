package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var client Doer
var endpoints Endpoints

type ServerOpts struct {
	Endpoints Endpoints
	Client    Doer
}

func StartServer(opts ServerOpts) {
	client = opts.Client
	endpoints = opts.Endpoints

	router := gin.Default()
	router.GET("/deck/:count", deck)
	router.Run(":8080")
}

func deck(c *gin.Context) {
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
