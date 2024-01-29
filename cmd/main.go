package main

import (
	"log"
	"rest_project/internal/server"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	
	server := server.NewServer()
	if err := server.Start(); err != nil{
		log.Fatal(err)
	}
}