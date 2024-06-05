package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"stockpull/controller"
	"stockpull/cronjobs"
	"stockpull/model"

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
	model.BlmTest.SetUrls()
	r.HandleFunc("/test", controller.Test)

	r.HandleFunc("/live/wsj-all-index", controller.WSJAllIndex)
	r.HandleFunc("/live/wsj-bonds", controller.WSJBonds)
	r.HandleFunc("/live/wsj-usa", controller.WSJUsaIndex)
	r.HandleFunc("/live/wsj-asia", controller.WsjAsia)
	r.HandleFunc("/live/wsj-europe", controller.WsjEurope)
	r.HandleFunc("/live/forex", controller.GetCurrencyValue)
	r.HandleFunc("/live/mint-india", controller.GetMintLiveNewsArray)
	r.HandleFunc("/live/bqprime", controller.GetBQPrimeTodaysAllYouNeedToKnowNews)
	r.HandleFunc("/bqprime-array", controller.GetBQPrimeAllYouNeedToKnowArray)
	r.HandleFunc("/forex", controller.GetCachedCurrencyValue)
	r.HandleFunc("/mint-news", controller.GetLiveMintNewsletterArray)
	r.HandleFunc("/live/nyt-dealbook-array", controller.GetNYTimesArrayDealBook)
	r.HandleFunc("/live/nyt-anz", controller.GetNYTimesArrayANZ)
	r.HandleFunc("/live/nyt-apac", controller.GetNYTimesArrayAPAC)
	r.HandleFunc("/live/nyt-europe", controller.GetNYTimesArrayEurope)
	r.HandleFunc("/live/nyt-us", controller.GetNYTimeArrayUsa)
	r.HandleFunc("/nyt-newsletter", controller.GetCachedNYTLetter)
	r.HandleFunc("/blm-news", controller.GetBloombergNewsLetter)
	r.HandleFunc("/version", controller.Version)

	cn := cron.New()
	//cron job to run every 2 min "*/2 * * * *"
	//cron job to run at 2H 5Min "5 2 * * *"
	//
	cn.AddFunc("5 2 * * *", cronjobs.SetBqPrimeNEwsLetterArray)
	cn.AddFunc("5 2 * * *", cronjobs.SetNYTNewsLetterToRedis)
	cn.AddFunc("5 2 * * *", cronjobs.BmlCronJob)
	cn.AddFunc("5 2 * * *", cronjobs.SetForexCronJob)

	//for testing
	//cn.AddFunc("* * * * *", cronjobs.SetBqPrimeNEwsLetterArray)
	//cn.AddFunc("5 2 * * *", cronjobs.SetMintTopOfMorningNewsletter)
	//cn.AddFunc("* 5 * * *", cronjobs.SetBqPrimeNEwsLetterArray)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
	cn.Start()
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
	//stubbing
	//cn.AddFunc("*/2 * * * *", cronjobs.SetBqPrimeNEwsLetterArray)
	// cn.AddFunc("*/2 * * * *", func() {
	// 	cronjobs.SetBqPrimeNEwsLetterArray()
	// 	fmt.Println("Cron running")
	// })

}
