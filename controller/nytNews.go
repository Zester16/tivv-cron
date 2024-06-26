package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
	"stockpull/utils"
)

//var ctx = context.Background()

//////////////*********ALL ENDPOINTS WILL BE FETCHING CACHED DATA FROM REDIS

// This is for all functions related to NYTimes
func GetNYTimesArrayDealBook(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_DEALBOOK)

	if err != nil || len(res) == 0 {
		fmt.Println("controller/GETNYTTimesArrayDealBook/Error: ", err)
		w.WriteHeader(400)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Some Error in downstream"))
		}
		return

	}
	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("controller/GETNYTTimesArrayDealBook/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// This function returns array of newsletter corresponding to AustraliaNewzealand
func GetNYTimesArrayANZ(w http.ResponseWriter, r *http.Request) {
	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_MORNING_AUS)

	if err != nil || len(res) == 0 {
		fmt.Println("controller/GetNYTimesArrayANZ/Error: ", err)
		w.WriteHeader(400)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Some Error in downstream"))
		}

	}

	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("controller/GetNYTimesArrayANZ/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// This function returns array of newletter APAC
func GetNYTimesArrayAPAC(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_MORNING_APAC)

	if err != nil || len(res) == 0 {
		fmt.Println("controller/GetNYTimesArrayAPAC/Error:", err)
		w.WriteHeader(400)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Some Error in downstream"))
		}
	}
	str, err := json.Marshal(res)

	if err != nil {

		fmt.Println("controller/GetNYTimesArrayAPAC/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// This function returns array of newsletter from europe
func GetNYTimesArrayEurope(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_MORNING_EUROPE)

	if err != nil || len(res) == 0 {

		fmt.Println("controller/GetNYTimesArrayEurope/Error: ", err)
		w.WriteHeader(400)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Some Error in downstream"))
		}

	}
	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("controller/GetNYTimesArrayEurope/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// This function returns array of newsletter from NYT for USA
func GetNYTimeArrayUsa(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_MORINIG_US)

	if err != nil || len(res) == 0 {
		fmt.Println("controller/GetNYTimesArrayUsa/Error: ", err)
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Some error happend in downstream"))
		}
	}

	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("controller/GetNYTimesArrayUsa/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// This function returns array of evening news from NYT USA
func GetNYTimeArrayEveningBriefing(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing(model.NYT_EVENING_US)

	if err != nil || len(res) == 0 {
		fmt.Println("controller/GetNYTimesArrayEveningBrief/Error:", err)

		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("some error happened in downstrea,"))
		}
	}

	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("controller/GetNYTimesArrayEveningBrief/Error: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(str))
}

// ********FUNCTION WILL GIVE Cached DATA**********************//
func GetCachedNYTLetter(w http.ResponseWriter, r *http.Request) {

	topic := r.URL.Query().Get("topic")

	if topic == "" {

		utils.SetErrorForNoQuery(w)
		return
	}

	nytUrls := model.BlmTest.GetNYTUrls()
	value := nytUrls[topic]

	if value == "" {
		utils.SetErrorForNoMatchingQuery(w)

		return
	}
	rdb := datasource.RedisConnect()

	redisRep, err := rdb.RedisDBConnector.Get(ctx, topic).Result()

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(redisRep))
}
