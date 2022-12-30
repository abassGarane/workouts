package main

import "os"

func initEnv() map[string]any {
	var envmap = map[string]any{}
	var PORT string
	if os.Getenv("PORT") == "" {
		PORT = ":4000"
	} else {
		PORT = os.Getenv("PORT")
	}

	envmap["PORT"] = PORT
	return envmap
}
