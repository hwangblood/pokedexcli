package main

import "os"

func callbackExit() error {
	os.Exit(0) // ending program without any error
	return nil
}
