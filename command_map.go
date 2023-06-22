package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}

func callbackMapBack(cfg *config) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("don't have any previous page of location areas")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}
