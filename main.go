package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"stockpull/controller"
	"stockpull/cronjobs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	//fmt.Printf("Hello fom golang")
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal(errEnv)
	}
	r := mux.NewRouter()
	r.HandleFunc("/test", controller.Test)

	r.HandleFunc("/wsj-asia", controller.WsjAsia)
	r.HandleFunc("/bloomberg", controller.GetBloombergTechNews)
	r.HandleFunc("/bqprime", controller.GetBQPrimeTodaysAllYouNeedToKnowNews)
	r.HandleFunc("/bqprime-array", controller.GetBQPrimeAllYouNeedToKnowArray)
	r.HandleFunc("/version", controller.Version)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	err := http.ListenAndServe(":"+port, r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)

	}

	cn := cron.New()
	cn.AddFunc("5 2 * * *", cronjobs.SetBqPrimeNEwsLetterArray)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
	cn.Start()

}
