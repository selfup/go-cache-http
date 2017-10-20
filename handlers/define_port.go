package handlers

import "os"

// DefinePort either grabs the PORT ENV or defines 8080 as the port
func DefinePort() string {
	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		return ":" + portEnv
	}

	return ":8081"
}
