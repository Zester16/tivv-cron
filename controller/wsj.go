package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
)

// controller handles REST endpoint relating to WSJ
func Test(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAsiaIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func WsjAsia(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAsiaIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func WsjEurope(w http.ResponseWriter, r *http.Request) {
	sb := network.GetEuropeIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func WSJAllIndex(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAllIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

func WSJBonds(w http.ResponseWriter, r *http.Request) {
	sb := network.GetWSJBonds()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

func WSJUsaIndex(w http.ResponseWriter, r *http.Request) {
	sb := network.GetWSJUSAIndex()
	j, _ := json.Marshal(sb)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

// gets WSJ Letters stored in DB and sends back
func WSJNewsLetterIndex(w http.ResponseWriter, r *http.Request) {
	topic := model.WSJ_ALL_NS
	rdb := datasource.RedisConnect()
	resArray, err := rdb.RedisDBConnector.Get(ctx, topic).Result()

	if err != nil {
		fmt.Println("WSJNewsletterIndex error:", err.Error())
		w.WriteHeader(400)

		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(resArray))

}
