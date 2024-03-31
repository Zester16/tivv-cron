package model

type ForexIndividualModel struct {
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Value  float64 `json:"value"`
	Change float64 `json:"change"`
	CP     float64 `json:"cp"`
}

type ForexRedisModel struct {
	Data []ForexIndividualModel `json:"data"`
	Date string                 `json:"date"`
}

const FOREX_KEY_NAME = "all_country_forex_mh"
