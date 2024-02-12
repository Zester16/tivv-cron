package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stockpull/network"
)

//////////////*********ALL ENDPOINTS WILL BE FETCHING CACHED DATA FROM REDIS

// This is for all functions related to NYTimes
func GetNYTimesArrayDealBook(w http.ResponseWriter, r *http.Request) {

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_dealbook")

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
	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_morning_aus")

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

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_morning_apac")

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

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_morning_europe")

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

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_morning_us")

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

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing("nyt_evening_us")

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

//********FUNCTION WILL GIVE Cached DATA**********************//
