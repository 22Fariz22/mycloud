package utils

import "log"

// GetConfigPath Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		log.Println("parse config-docker")
		return "./config/config-docker"
	}
	log.Println("parse config-local")
	return "./config/config-local"
}
