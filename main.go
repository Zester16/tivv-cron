package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"stockpull/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	//fmt.Printf("Hello fom golang")
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal(errEnv)
	}
	test := os.Getenv("test")
	fmt.Println(test)

	redisUrl := os.Getenv("redis")
	fmt.Println(redisUrl)
	opt, redisEr := redis.ParseURL(redisUrl)

	if redisEr != nil {
		panic(redisEr)
	}
	rdb := redis.NewClient(opt)

	rdb.Set(ctx, "test", "test", 0).Err()
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	r.HandleFunc("/test", controller.Test)

	r.HandleFunc("/wsj-asia", controller.WsjAsia)
	r.HandleFunc("/bloomberg", controller.GetBloombergTechNews)
	r.HandleFunc("/bqprime", controller.GetBQPrimeTodaysAllYouNeedToKnowNews)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }

	err := http.ListenAndServe(":"+port, r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)

	}

}
