out = bin/pokedexcli

start:
	go build -o $(out) && $(out)
test:
	go test ./...

.PHONY: start test