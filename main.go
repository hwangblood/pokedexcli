package main

import "github.com/hwangblood/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)

}
