package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locations.Previous
	cfg.nextLocationsURL = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locations.Previous
	cfg.nextLocationsURL = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
