package main

import (
	"fmt"
	"golang-doc/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	viper "github.com/spf13/viper"
)

func Add(a, b int) int{
	return a + b
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("application.port"))
	port := conf.Application.Port

	log.Println(Add(5,6))

	router := mux.NewRouter()
	http.ListenAndServe(":" + port, router)
}
