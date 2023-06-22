out = bin/pokedexcli

start:
	go build -o $(out) && $(out)

.PHONY: start