package config

import "fmt"

type Config struct {
	DbFilename string
	Port       string
}

func NewDefaultConfig() *Config {
	return &Config{DbFilename: "db.txt", Port: ":8080"}
}

func SayHello() {
	fmt.Println("SayHello()-->Hello")
}
