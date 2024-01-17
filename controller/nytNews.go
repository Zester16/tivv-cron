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

	res, err := network.PostCrawlGetNYTimeArrayEveningBriefing()

	if err != nil {
		fmt.Println("GETNYTTimesArrayDealBook", err)

	}
	str, err := json.Marshal(res)

	if err != nil {
		fmt.Println("GETNYTTimesArrayDealBook", err)

	}
	w.Write([]byte(str))
}

// This function returns array of newsletter corresponding to AustraliaNewzealand
func GetNYTimesArrayANZ(w http.ResponseWriter, r *http.Request) {

}

// This function returns array of newletter APAC
func GetNYTimesArrayAPAC(w http.ResponseWriter, r *http.Request) {

}

// This function returns array of newsletter from europe
func GetNYTimesArrayEurope(w http.ResponseWriter, r *http.Request) {

}

// This function returns array of newsletter from NYT for USA
func GetNYTimeArrayUsa(w http.ResponseWriter, r *http.Request) {

}

// This function returns array of evening news from NYT USA
func GetNYTimeArrayEveningBriefing(w http.ResponseWriter, r *http.Request) {

}

//********FUNCTION WILL GIVE LIVE DATA**********************//
