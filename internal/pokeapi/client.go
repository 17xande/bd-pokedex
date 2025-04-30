package pokeapi

import (
	"net/http"
	"time"

	"github.com/17xande/bd-pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(timeout)

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
