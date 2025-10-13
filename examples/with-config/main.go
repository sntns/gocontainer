package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	Port    string `json:"port"`
	Message string `json:"message"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("/etc/myapp/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return &config, err
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Printf("Warning: Could not load config: %v", err)
		config = &Config{
			AppName: "DefaultApp",
			Port:    "8080",
			Message: "Hello from default config!",
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s says: %s\n", config.AppName, config.Message)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Printf("%s starting on port %s", config.AppName, config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}